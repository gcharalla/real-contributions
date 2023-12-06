package filesystem

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
)

// getDotFilePath returns the dot file for the repos list.
// Creates it and the enclosing folder if it does not exist.
func GetDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFolderPath := path.Join(usr.HomeDir, ".real-contributions")
	if err := os.MkdirAll(dotFolderPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	dotFile := path.Join(dotFolderPath, ".gitlocalstats")

	return dotFile
}

// openFile opens the file located at `filePath`. Creates it if not existing.
func OpenFile(filePath string) *os.File {
	//f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0755)
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

// parseFileLinesToSlice given a file path string, gets the content
// of each line and parses it to a slice of strings.
func ParseFileLinesToSlice(filePath string) []string {
	f := OpenFile(filePath)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

// sliceContains returns true if `slice` contains `value`
func SliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// joinSlices adds the element of the `new` slice
// into the `existing` slice, only if not already there
func JoinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !SliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

// dumpStringsSliceToFile writes content to the file in path `filePath` (overwriting existing content)
func DumpStringsSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	os.WriteFile(filePath, []byte(content), 0755)
}

// addNewSliceElementsToFile given a slice of strings representing paths, stores them
// to the filesystem
// addNewSliceElementsToFile given a slice of strings representing paths, stores them
// to the filesystem
func AddNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := ParseFileLinesToSlice(filePath)
	repos := JoinSlices(newRepos, existingRepos)
	DumpStringsSliceToFile(repos, filePath)
}
