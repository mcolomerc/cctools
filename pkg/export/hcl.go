package export

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/log"
	"os"

	hclprinter "github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)

type HclExporter struct {
}

func (e HclExporter) GetPath() string {
	return "tfvars"
}

func (e HclExporter) Export(res interface{}, outputPath string) error {
	jsonString, _ := json.Marshal(res)
	fmt.Printf("s %v", string(jsonString))

	return ToHcl(jsonString, outputPath)
}

// ToHcl byte array convert **hcl** out to io.Writer
func ToHcl(input []byte, outputPath string) error {
	ast, err := jsonParser.Parse(input)
	if err != nil {
		return fmt.Errorf("unable to decodeobject ast.node: %s", err)
	}
	var buf bytes.Buffer
	if err := hclprinter.Fprint(&buf, ast); err != nil {
		return fmt.Errorf("unable to decodeobject ast.node: %s", err)
	}
	err = hclprinter.Fprint(os.Stdout, ast)
	if err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}
	log.Debug("Writing file ")
	err2 := ioutil.WriteFile(outputPath+"_hcl.tfvars", buf.Bytes(), 0644)

	if err2 != nil {
		log.Fatal(err)
	}

	return nil
}
