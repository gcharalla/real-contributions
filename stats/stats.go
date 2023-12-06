package stats

import (
	"github.com/gcharalla/real-contributions/constants"
	"github.com/gcharalla/real-contributions/filesystem"
	"github.com/gcharalla/real-contributions/table"
	"github.com/gcharalla/real-contributions/timeMachine"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// stats calculates and prints the stats.
func Stats(email string) {
	commits := processRepositories(email)
	table.PrintCommitsStats(commits)
}

// fillCommits given a repository found in `path`, gets the commits and
// puts them in the `commits` map, returning it when completed
func fillCommits(email string, path string, commits map[int]int) map[int]int {
	// instantiate a git repo object from path
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}
	// get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	// get the commits history starting from HEAD
	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}
	// iterate the commits
	offset := timeMachine.CalcOffset()
	err = iterator.ForEach(func(c *object.Commit) error {
		daysAgo := timeMachine.CountDaysSinceDate(c.Author.When) + offset

		if c.Author.Email != email {
			return nil
		}

		if daysAgo != constants.OutOfRange {
			commits[daysAgo]++
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return commits
}

// processRepositories given a user email, returns the
// commits made in the last 6 months
func processRepositories(email string) map[int]int {
	filePath := filesystem.GetDotFilePath()
	repos := filesystem.ParseFileLinesToSlice(filePath)
	daysInMap := constants.DaysInLastSixMonths

	commits := make(map[int]int, daysInMap)
	for i := daysInMap; i > 0; i-- {
		commits[i] = 0
	}

	for _, path := range repos {
		commits = fillCommits(email, path, commits)
	}

	return commits
}

// calcOffset determines and returns the amount of days missing to fill
// the last row of the stats graph
