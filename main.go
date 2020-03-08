package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Don't do anything
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}

	return result
}

func main() {
	topPtr := flag.Bool("top", false, "")

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide a path as argument.")
	}

	if args[1:][0] == "." {
		// Use current path
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		update(path, *topPtr)
	} else {
		// Use passed in path
		update(args[1:][0], *topPtr)
	}
}

// Track all files in given folder path, commit with file names as commit msg & push to remote.
func update(path string, top bool) {
	cmd := exec.Command("git")
	cmd.Dir = path
	cmd.Args = []string{"git", "diff", "HEAD", "--name-only"}
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		outS := strings.Fields(string(out))
		filesChanged := make([]string, 0)
		// Get all files changed without extension
		for _, v := range outS {
			split := strings.Split(v, "/")
			if top {
				first := split[0]
				filesChanged = append(filesChanged, first)
			} else {
				last := split[len(split)-1]
				filesChanged = append(filesChanged, strings.Split(last, ".")[0])
			}
		}
		if top {
			filesChanged = removeDuplicates(filesChanged)
		}
		// Track files changed by Git
		cmd = exec.Command("git")
		cmd.Dir = path
		cmd.Args = []string{"git", "add", "."}
		_, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		commitMsg := strings.Join(filesChanged, " ")

		// Commit with a message
		cmd = exec.Command("git")
		cmd.Dir = path
		cmd.Args = []string{"git", "commit", "-m", commitMsg}
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}

		// Push changes
		cmd = exec.Command("git")
		cmd.Dir = path
		cmd.Args = []string{"git", "push"}
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
	}
}
