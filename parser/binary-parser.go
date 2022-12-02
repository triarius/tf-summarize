package parser

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"github.com/triarius/tf-summarize/terraform_state"
)

type BinaryParser struct {
	fileName string
}

func (j BinaryParser) Parse() (terraform_state.TerraformState, error) {
	cmd := exec.Command("terraform", "show", "-json", j.fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return terraform_state.TerraformState{}, fmt.Errorf(
			"error when running 'terraform show -json %s': \n%s\n\n%s",
			j.fileName, output, "Make sure you are running in terraform directory and terraform init is done")
	}
	ts := terraform_state.TerraformState{}
	err = json.Unmarshal(output, &ts)
	if err != nil {
		return terraform_state.TerraformState{}, fmt.Errorf("error when parsing input: %s", err.Error())
	}
	return ts, nil
}

func NewBinaryParser(fileName string) Parser {
	return BinaryParser{
		fileName: fileName,
	}
}
