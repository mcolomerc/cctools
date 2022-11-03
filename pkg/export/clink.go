package export

import (
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/model"
	"os"
)

type ClinkExporter struct {
	LinkName             string
	SourceClusterId      string
	BootstrapServer      string
	SourceApiKey         string
	SourceApiSecret      string
	DestinationClusterId string
	AclSync              bool
	OffsetSync           bool
	AutoCreate           bool
}

const (
	SH_HEADER = "#!/bin/bash"
)

func (e ClinkExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	done := make(chan bool, 4)
	go e.buildCreationScript(topics, outputPath, done)
	go e.buildProperties(topics, outputPath, done)
	go e.buildCleanup(outputPath, done)
	go e.buildPromote(topics, outputPath, done)

	for i := 0; i < 4; i++ {
		<-done
	}
	close(done)
	return nil
}

func (e ClinkExporter) buildCreationScript(topics []model.Topic, outputPath string, done chan bool) error {
	f, err := os.Create(outputPath + "_clink_create.sh")
	if err != nil {
		log.Println("Error creating file ...")
		return err
	}
	var lines []string
	lines = append(lines, SH_HEADER)
	lines = append(lines, `echo "Create the Cluster Link"`)
	lines = append(lines, e.buildLink(outputPath))
	lines = append(lines, "")

	if !e.AutoCreate {
		for _, v := range topics {
			lines = append(lines, `echo "Create Mirror Topic: `+v.Name+`"`)
			lines = append(lines, e.buildTopicMirror(v.Name))
			lines = append(lines, "")
		}
	}
	for _, v := range lines {
		_, errw := f.WriteString(v + "\n")
		if errw != nil {
			log.Println("Error writing file ...")
			return errw
		}
	}
	f.Sync()

	defer f.Close()
	done <- true
	return nil
}

func (e ClinkExporter) buildPromote(topics []model.Topic, outputPath string, done chan bool) error {
	f, err := os.Create(outputPath + "_clink_promote.sh")
	if err != nil {
		log.Println("Error creating file ...")
		return err
	}
	var lines []string
	lines = append(lines, SH_HEADER)
	lines = append(lines, `echo "Promote mirrors."`)
	lines = append(lines, "")

	for _, v := range topics {
		lines = append(lines, `echo "Promote mirror Topic: `+v.Name+`"`)
		lines = append(lines, e.buildPromoteMirror(v.Name))
		lines = append(lines, "")
	}
	for _, v := range lines {
		_, errw := f.WriteString(v + "\n")
		if errw != nil {
			log.Println("Error writing file ...")
			return errw
		}
	}
	f.Sync()
	defer f.Close()
	done <- true
	return nil
}

func (e ClinkExporter) buildLink(outputPath string) string {
	return fmt.Sprintf(`confluent kafka link create %s --cluster %s \
    --source-cluster-id %s  \
    --source-bootstrap-server %s  \
    --source-api-key %s \ 
	--source-api-secret %s \
    --config-file %s`,
		e.LinkName,
		e.DestinationClusterId,
		e.SourceClusterId,
		e.BootstrapServer,
		e.SourceApiKey,
		e.SourceApiSecret,
		outputPath+"link.properties",
	)
}

func (e ClinkExporter) buildTopicMirror(topic string) string {
	return fmt.Sprintf(`confluent kafka mirror create %s --cluster %s --link %s `, topic, e.DestinationClusterId, e.LinkName)
}
func (e ClinkExporter) buildPromoteMirror(topic string) string {
	return fmt.Sprintf(`confluent kafka mirror promote %s --cluster %s --link %s `, topic, e.DestinationClusterId, e.LinkName)
}

func (e ClinkExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return nil
}

func (e ClinkExporter) buildCleanup(outputPath string, done chan bool) error {
	f, err := os.Create(outputPath + "_clink_cleanup.sh")
	if err != nil {
		log.Println("Error creating file ...")
		return err
	}
	var lines []string
	lines = append(lines, SH_HEADER)
	lines = append(lines, `echo "Clean up..."`)

	clean := fmt.Sprintf("confluent kafka link delete %s --cluster %s", e.LinkName, e.DestinationClusterId)
	lines = append(lines, clean)
	lines = append(lines, "")

	for _, v := range lines {
		_, errw := f.WriteString(v + "\n")
		if errw != nil {
			log.Println("Error writing file ...")
			return errw
		}
	}

	f.Sync()
	defer f.Close()
	done <- true
	return nil

}

func (e ClinkExporter) buildProperties(topics []model.Topic, outputPath string, done chan bool) error {
	f, err := os.Create(outputPath + "link.properties")
	if err != nil {
		log.Println("Error creating file ...")
		return err
	}
	var lines []string
	autoCreate := fmt.Sprintf("auto.create.mirror.topics.enable=%t ", e.AutoCreate)
	lines = append(lines, autoCreate)

	if e.AutoCreate {
		filters := `auto.create.mirror.topics.filters={ "topicFilters": [ `
		for _, v := range topics {
			filters = filters + fmt.Sprintf(`{"name": "%s",  "patternType": "LITERAL",  "filterType": "INCLUDE"}, `, v.Name)
		}
		filters = filters + `] } \n`
		lines = append(lines, filters)
	}

	offset := fmt.Sprintf("consumer.offset.sync.enable=%t", e.OffsetSync)
	lines = append(lines, offset)
	acl := fmt.Sprintf("acl.sync.enable=%t ", e.AclSync)
	lines = append(lines, acl)
	if e.AclSync {
		aclsync := "acl.sync.ms=1000"
		lines = append(lines, aclsync)
		aclFilters := `acl.filters={ "aclFilters": [ { "resourceFilter": { "resourceType": "any", "patternType": "any" }, "accessFilter": { "operation": "any", "permissionType": "any" } } ] }`
		lines = append(lines, aclFilters)
	}

	for _, v := range lines {
		_, errw := f.WriteString(v + "\n")
		if errw != nil {
			log.Println("Error writing file ...")
			return errw
		}
	}
	f.Sync()
	defer f.Close()
	done <- true
	return nil
}
