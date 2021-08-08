package main

import (
	"fmt"

	"github.com/MiteshSharma/TerraformExec/cli"
	"github.com/MiteshSharma/TerraformExec/git"
	"github.com/MiteshSharma/TerraformExec/terraform"
)

var directory = "/tmp/test"
var repo = ""
var postgresConnString = "conn_str='postgresql://mitesh:mitesh@localhost/terraform?sslmode=disable'"

func main() {
	cliExecutor := &cli.CliExecutor{
		Directory: directory,
	}
	git := &git.GitExecutor{
		Repo:      repo,
		Directory: directory,
	}
	git.DownloadCode()
	te := &terraform.TerraformExecutor{
		Executor:  cliExecutor,
		Directory: directory,
		Vars: map[string]interface{}{
			"str": "foooz",
			"num": 2,
		},
		BackendConfig: postgresConnString,
	}
	te.Init()
	te.Plan()
	te.Apply()
	out, _ := te.Output()
	fmt.Println(out)
	te.Destroy()
}
