package terraform

import (
	"encoding/json"
	"fmt"

	"github.com/MiteshSharma/TerraformExec/cli"
)

type Terraform interface {
	Init()
	Apply()
	Plan()
	Output()
	Destroy()
}

type TerraformExecutor struct {
	Directory     string
	Vars          map[string]interface{}
	BackendConfig string
	Executor      cli.Executor
}

func (te *TerraformExecutor) getTfArgument(command string, isAutoApply bool) []string {
	var args []string
	args = append(args, command)

	var flags []string
	if isAutoApply {
		flags = append(flags, "-auto-approve")
	}

	var vars []string
	for key, value := range te.Vars {
		vars = append(vars, "-var", fmt.Sprintf("%s=%x", key, value))
	}

	var backend []string
	if te.BackendConfig != "" && command == "init" {
		backend = append(backend, "-backend-config")
		backend = append(backend, te.BackendConfig)
	}

	if len(backend) > 0 {
		args = append(args, backend...)
	}

	if len(flags) > 0 {
		args = append(args, flags...)
	}

	if len(vars) > 0 {
		args = append(args, vars...)
	}
	return args
}

func (te *TerraformExecutor) getTfOutputArgument() []string {
	args := []string{"output", "-json"}
	return args
}

func (te *TerraformExecutor) runTfCommand(args []string) (*[]byte, error) {
	var output *[]byte
	var err error

	if output, err = te.Executor.Execute("terraform", args); err != nil {
		return nil, err
	}

	return output, nil
}

func (te *TerraformExecutor) Init() {
	fmt.Println("Initing")
	args := te.getTfArgument("init", false)
	te.runTfCommand(args)
	fmt.Println("Init complete")
}

func (te *TerraformExecutor) Plan() {
	fmt.Println("Planning")
	args := te.getTfArgument("plan", false)
	te.runTfCommand(args)
	fmt.Println("Plan complete")
}

func (te *TerraformExecutor) Apply() {
	fmt.Println("Applying")
	args := te.getTfArgument("apply", true)
	te.runTfCommand(args)
	fmt.Println("Apply complete")
}

func (te *TerraformExecutor) Destroy() {
	fmt.Println("destroying")
	args := te.getTfArgument("destroy", true)
	te.runTfCommand(args)
	fmt.Println("destroy complete")
}

func (te *TerraformExecutor) Output() (map[string]interface{}, error) {
	var outputMap map[string]interface{}

	args := te.getTfOutputArgument()
	output, err := te.runTfCommand(args)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(*output, &outputMap); err != nil {
		return nil, err
	}

	return outputMap, nil
}
