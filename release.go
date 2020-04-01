package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
)

func main() {
	out, err := executeTagCommand()
	if err != nil {
		log.Fatal(err)
	}

	nextPatch, err := findNextPatch(out)
	if err != nil {
		log.Fatal(err)
	}

	newVersion, err := promptForNewVersion(nextPatch)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Creating new release: %s\n", newVersion)
	if err = tagNewVersion(newVersion); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

// executeTagCommand executes 'git tag' and returns the output
func executeTagCommand() ([]byte, error) {
	out, err := exec.Command("git", "tag").Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

// findNextPatch expects the output from an invocation of 'git tag' as input.
// findNextPatch will find the next semver patch by looking through all tags and
// incrementing the highest semver found.
// Returns 0.0.1 if no semver tags where found.
func findNextPatch(bs []byte) (string, error) {
	rs := strings.Split(strings.TrimSpace(string(bs)), "\n")

	var vs []*semver.Version

	for _, r := range rs {
		v, err := semver.NewVersion(r)
		if err != nil {
			continue
		}
		vs = append(vs, v)
	}

	if vs == nil {
		return "0.0.1", nil
	}

	sort.Sort(semver.Collection(vs))
	nextPatch := vs[len(vs)-1].IncPatch()

	return nextPatch.Original(), nil
}

// promptForNewVersion expects as input a semver string.
// promptForNewVersion creates a command line prompt to allow a
// user to either accept a proposed version or specify another one.
func promptForNewVersion(s string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is the release version? %s: ", s)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	trimmed := strings.TrimSpace(input)
	newVersion := s
	if len(trimmed) > 0 {
		newVersion = trimmed
	}

	version, err := semver.NewVersion(newVersion)
	if err != nil {
		return "", err
	}
	return version.Original(), nil
}

// tagNewVersion creates and pushes an annotated tag of a given semver.
func tagNewVersion(s string) error {
	_, err := exec.Command("git", "tag", "-a", s, "-m", "New release").Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("git", "push", "origin", s).Output()
	if err != nil {
		return err
	}
	return nil
}
