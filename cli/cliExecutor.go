package cli

import (
	"io/ioutil"
	"os"
	"os/exec"
)

type Executor interface {
	Execute(command string, args []string) (*[]byte, error)
}

type CliExecutor struct {
	Directory string
}

func (c *CliExecutor) Execute(command string, args []string) (*[]byte, error) {
	execCmd := exec.Command(command, args...)
	execCmd.Dir = c.Directory
	execCmd.Stderr = os.Stderr
	// Added to test
	// execCmd.Stdout = os.Stdout

	stdout, err := execCmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = execCmd.Start()
	if err != nil {
		return nil, err
	}

	var output []byte
	output, err = ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	err = execCmd.Wait()
	if err != nil {
		return nil, err
	}

	return &output, nil
}
