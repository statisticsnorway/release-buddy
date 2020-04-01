package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNextPatch(t *testing.T) {
	tests := map[string]struct {
		input     []byte
		wantPatch string
		wantErr   error
	}{
		"no tags":                      {input: []byte{}, wantPatch: "0.0.1", wantErr: nil},
		"empty string":                 {input: []byte(""), wantPatch: "0.0.1", wantErr: nil},
		"no semver tags":               {input: []byte("foo\n"), wantPatch: "0.0.1", wantErr: nil},
		"single tag":                   {input: []byte("9.9.8\n"), wantPatch: "9.9.9", wantErr: nil},
		"two tags":                     {input: []byte("4.5.1\n1.10.2\n"), wantPatch: "4.5.2", wantErr: nil},
		"semver mixed with other tags": {input: []byte("foo\n1.2.2\nbar\n"), wantPatch: "1.2.3", wantErr: nil},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p, err := findNextPatch(test.input)
			assert.Equal(t, test.wantPatch, p)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
