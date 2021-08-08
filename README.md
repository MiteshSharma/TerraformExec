# TerraformExec

This project is used to run terraform files using go code. We have git module which clone repository in a given directory. Using cli exec command, go code execute terraform files cloned using git.

We need a git repository to download tf code and directory where it needs to be downloaded. We are using postgres as terraform backend, so we need to provide a connection staring for this.

Go to main.go, replace repository, directory, any variables to be added and postgres connection string.

Actions performed:
1. Init
2. Plan
3. Apply
4. Output: Returns map of type map[string]interface{} which can be used to fetch output parameters.
5. Destroy

For more details around terraform: https://www.terraform.io/ 

Requirement:
terraform cli must be downloaded on system and must be executable. 
