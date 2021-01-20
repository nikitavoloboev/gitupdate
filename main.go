package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	// Option to only consider top level folders
	topLevelOnly := flag.Bool("top", false, "")

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide a path as argument.")
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if args[1:][0] == "." {
		update(path, *topLevelOnly)
	} else {
		update(path+"/"+args[1:][0], *topLevelOnly)
	}
}

// Track & update files in passed in path.
// If it's folder, commit entire folder. If one file, commit the file.
// Commit with file names as commit msg & push to remote.
func update(path string, topLevelOnly bool) {
	cmd := exec.Command("git")
	file, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	// if it's not dir, then strip the extension
	if !file.IsDir() {
		path = filepath.Dir(path)
	}
	cmd.Dir = path
	cmd.Args = []string{"git", "add", path}
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		cmd = exec.Command("git")
		cmd.Dir = path
		cmd.Args = []string{"git", "diff", "--cached", "HEAD", "--name-only"}
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		} else {
			outS := strings.Fields(string(out))
			filesChanged := make([]string, 0)
			// Get all files changed without extension
			for _, v := range outS {
				split := strings.Split(v, "/")
				if topLevelOnly {
					first := split[0]
					filesChanged = append(filesChanged, first)
				} else {
					last := split[len(split)-1]
					filesChanged = append(filesChanged, strings.Split(last, ".")[0])
				}
			}
			filesChanged = removeDuplicates(filesChanged)

			// Commit with a message
			commitMsg := strings.Join(filesChanged, " ")
			cmd = exec.Command("git")
			cmd.Dir = path
			println("about to commit with message:", commitMsg)
			println("buuuuut.... we're not ready for that yet")
			// return
			cmd.Args = []string{"git", "commit", "-m", commitMsg}
			out, err = cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
			println(out)

			// Push changes
			return;	
			cmd = exec.Command("git")
			cmd.Dir = path
			cmd.Args = []string{"git", "push"}
			_, err = cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
