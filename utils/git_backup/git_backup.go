package main

import (
	"fmt"
	"os"
	// "log"
	"path/filepath"
	"io/ioutil"
	"github.com/go-git/go-git/v5"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func update_remote(git_directory string) {
	fmt.Println("DIRECTORIO: ", git_directory)

	// r, err := git.PlainOpen(git_directory)
	// if err != nil {
	// 	fmt.Println("THIS IS NOT A GIT DIRECTORY")
	// }
	// fmt.Println(r)

	// remote, err := r.Remote("gitlab")
	// if err != nil {
	// 	fmt.Println("THIS REPO HAS NOT GITLAB REMOTE")
	// }
	// fmt.Println(remote)

	// We instance a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(git_directory)
	CheckIfError(err)

	// Get the working directory for the repository
	fmt.Println("worktree")
	w, err := r.Worktree()
	CheckIfError(err)

	// Pull the latest changes from the origin remote and merge into the current branch
	fmt.Println("pull")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	CheckIfError(err)

	// Print the latest commit that was just pulled
	ref, err := r.Head()
	CheckIfError(err)
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)

}

func main() {
	git_directory := filepath.Join(os.Getenv("HOME"), "/git")
	files, err := ioutil.ReadDir(git_directory)
	CheckIfError(err)

    for _, file := range files {
		if file.IsDir(){
			update_remote(filepath.Join(git_directory, file.Name()))
		}
    }
}
