package git

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

type Git interface {
	DownloadCode()
	CleanCode()
}

type GitExecutor struct {
	Repo      string
	Directory string
}

func (ge *GitExecutor) DownloadCode() {
	_, err := os.Stat(ge.Directory)
	if !os.IsNotExist(err) {
		err = os.RemoveAll(ge.Directory)
		if err != nil {
			fmt.Println("Error")
		}
	}

	_, err = git.PlainClone(ge.Directory, false, &git.CloneOptions{
		URL: ge.Repo,
	})
	fmt.Println("Downloaded code in " + ge.Directory)
	if err != nil {
		fmt.Println(ge.Directory)
	}
}

func (ge *GitExecutor) CleanCode() {
	err := os.RemoveAll(ge.Directory)
	if err != nil {
		fmt.Println(err)
	}
}
