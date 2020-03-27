package main

import (
	"bufio"
	"fmt"
	"github.com/Masterminds/semver"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	out, err := executeTagCommand()
	if err != nil {
		log.Fatal(err)
	}

	nextPatch, err := nextPatch(out)
	if err != nil {
		log.Fatal(err)
	}

	newVersion, err := newVersion(nextPatch)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tagging new version: %s\n", newVersion)
	if err = tagNewVersion(newVersion); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func executeTagCommand() ([]byte, error) {
	out, err := exec.Command("git", "tag").Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func nextPatch(bs []byte) (string, error) {
	s := string(bs)

	s = strings.Replace(s, "\n", "", -1)

	raw := strings.Split(strings.TrimSpace(string(bs)), "\n")

	vs := make([]*semver.Version, len(raw))
	for i, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			continue
		}
		vs[i] = v
	}

	if len(vs) == 0 {
		return "0.0.1", nil
	}

	sort.Sort(semver.Collection(vs))

	nextPatch := vs[len(vs)-1].IncPatch()

	return nextPatch.Original(), nil
}

func newVersion(s string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is the release version? %s: ", s)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	trimmed := strings.TrimSpace(text)
	vs := s
	if len(trimmed) > 0 {
		vs = trimmed
	}

	version, err := semver.NewVersion(vs)
	if err != nil {
		return "", err
	}
	return version.Original(), nil
}

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
