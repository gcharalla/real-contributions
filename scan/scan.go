package scan

import (
	"fmt"

	"log"
	"os"

	"strings"

	"github.com/gcharalla/real-contributions/filesystem"
)

// scanGitFolders returns a list of subfolders of `folder` ending with `.git`.
// Returns the base folder of the repo, the .git folder parent.
// Recursively searches in the subfolders by passing an existing `folders` slice.
func ScanGitFolders(folders []string, folder string) []string {
	// trim the last `/`
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			folders = ScanGitFolders(folders, path)
		}
	}

	return folders
}

// recursiveScanFolder starts the recursive search of git repositories
// living in the `folder` subtree
func RecursiveScanFolder(folder string) []string {
	return ScanGitFolders(make([]string, 0), folder)
}

// scan scans a new folder for Git repositories
func Scan(folder string) {
	fmt.Printf("Found folders:\n\n")
	repositories := RecursiveScanFolder(folder)
	filePath := filesystem.GetDotFilePath()
	filesystem.AddNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}
