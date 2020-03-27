package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitTag(t *testing.T) {
	bytes, err := executeTagCommand()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes)
}

func TestGetNextPatch(t *testing.T) {

	tests := map[string]struct {
		input     []byte
		wantPatch string
		wantErr   error
	}{
		"simple": {input: []byte("0.0.1"), wantPatch: "0.0.2", wantErr: nil},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p, err := nextPatch(test.input)
			assert.Equal(t, test.wantPatch, p)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
