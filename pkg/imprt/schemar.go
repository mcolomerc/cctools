package imprt

import (
	"encoding/json"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"

	"github.com/mitchellh/mapstructure"
)

type SchemaImport struct {
	DestClient client.RestClient
	Conf       config.Config
	Paths      SchemaPaths
	Importer   Importer
}

type SchemaPaths struct {
	Schemas  string
	Subjects string
}

type NewSubject struct {
	Schema string `json:"schema"`
	//Subject string `json:"subject"`
	//SchemaType string `json:"schemaType,omitempty"`
	Version int `json:"version,omitempty"`
	Id      int `json:"id,omitempty"`
}

const (
	SCHEMAS_PATH  = "/schemas/"
	SUBJECTS_PATH = "/subjects/"
)

func NewSchemaImport(conf config.Config) (*SchemaImport, error) {
	paths := &SchemaPaths{
		Schemas:  conf.Import.Source + SCHEMAS_PATH,
		Subjects: conf.Import.Source + SUBJECTS_PATH,
	}
	destClient := client.NewRestClient(conf.Source.SchemaRegistry.EndpointUrl, conf.Source.SchemaRegistry.Credentials)

	imp, err := NewImporter(conf)
	if err != nil {
		log.Error("Error creating importer : " + err.Error())
		return nil, err
	}
	kafkaService := &SchemaImport{
		DestClient: *destClient,
		Conf:       conf,
		Paths:      *paths,
		Importer:   imp,
	}
	return kafkaService, nil
}

func (s *SchemaImport) Import() {
	err := s.ImportSubjects()
	if err != nil {
		log.Error("Error importing subjects : " + err.Error())
	}
}

func (s *SchemaImport) ImportSubjects() error {
	// Import Schemas
	objs, err := s.Importer.Import(s.Paths.Subjects, model.SubjectVersion{})
	if err != nil {
		log.Error("Error importing schema registry " + err.Error())
		return err
	}

	log.Info("Mode Schema Registry : " + string(model.Import))
	for _, o := range objs {
		typed := model.SubjectVersion{}
		cfg := &mapstructure.DecoderConfig{
			Metadata: nil,
			Result:   &typed,
			TagName:  "json",
		}
		decoder, _ := mapstructure.NewDecoder(cfg)
		decoder.Decode(o)

		subj := NewSubject{
			Schema:  typed.Schema,
			Version: typed.Version,
			Id:      typed.Id,
		}
		jsonBytes, err := json.Marshal(subj)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debug(string(jsonBytes))
		errMode := s.SetMode(model.Import, typed.Subject)
		if errMode != nil {
			log.Error("Error setting mode Schema Registry :" + errMode.Error())
			log.Error("Check if subject exists in destination Schema Registry")
			return errMode
		}
		// Register subject
		_, err = s.DestClient.Post(s.Conf.Destination.SchemaRegistry.EndpointUrl+"/subjects/"+typed.Subject+"/versions", jsonBytes)
		if err != nil {
			log.Error("Error post Schema Registry : " + err.Error())
			return err
		}
		//curl -X PUT -H "Content-Type: application/json" "http://localhost:8081/mode/my-cool-subject" --data '{"mode": "READWRITE"}'
		errM := s.SetMode(model.ReadWrite, typed.Subject)
		if errM != nil {
			log.Error(errM)
			return errM
		}
	}

	return nil
}

type SRMode struct {
	Mode string `json:"mode"`
}

func (s *SchemaImport) SetMode(mode model.SchemaRegistryMode, subject string) error {
	m := SRMode{Mode: string(mode)}
	modeBytes, err := json.Marshal(m)
	log.Info("Setting mode Schema Registry : " + s.Conf.Destination.SchemaRegistry.EndpointUrl + "/mode/" + subject)
	log.Debug(string(modeBytes))
	if err != nil {
		log.Error(err)
		return err
	}
	_, errResp := s.DestClient.Put(s.Conf.Destination.SchemaRegistry.EndpointUrl+"/mode/"+subject, modeBytes)
	if errResp != nil {
		log.Error("Error setting mode Schema Registry : " + errResp.Error())
		log.Error("Check if subject exists in destination Schema Registry")
		return err
	}
	return nil
}
