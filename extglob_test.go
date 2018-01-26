package app

import (
	"testing"
)

type TableEntry struct {
	Glob       string
	Matches    []string
	NonMatches []string
}

var (
	table = []TableEntry{
		{
			Glob:       "asdf/*.jpg",
			Matches:    []string{"asdf/asdf.jpg", "asdf/asdf_asdf.jpg", "asdf/.jpg"},
			NonMatches: []string{"asdf/asdf/asdf.jpg", "xxxasdf/asdf.jpgxxx"},
		},
		{
			Glob:       "asdf/**.jpg",
			Matches:    []string{"asdf/asdf.jpg", "asdf/asdf_asdf.jpg", "asdf/asdf/asdf.jpg", "asdf/asdf/asdf/asdf/asdf.jpg"},
			NonMatches: []string{"/asdf/asdf.jpg", "asdff/asdf.jpg", "xxxasdf/asdf.jpgxxx"},
		},
		{
			Glob:       "asdf/*.@(jpg|jpeg)",
			Matches:    []string{"asdf/asdf.jpg", "asdf/asdf_asdf.jpeg"},
			NonMatches: []string{"/asdf/asdf.jpg", "asdff/asdf.jpg"},
		},
		{
			Glob:       "**/*.js",
			Matches:    []string{"asdf/asdf.js", "asdf/asdf/asdfasdf_asdf.js", "/asdf/asdf.js", "/asdf/aasdf-asdf.2.1.4.js"},
			NonMatches: []string{"/asdf/asdf.jpg", "asdf.js"},
		},
	}
)

func Test_CompileExtGlob(t *testing.T) {
	for _, entry := range table {
		r, err := CompileExtGlob(entry.Glob)
		if err != nil {
			t.Fatalf("Couldn’t compile glob %s: %s", entry.Glob, err)
		}
		t.Logf("Compiled glob %s: %s", entry.Glob, r)
		for _, match := range entry.Matches {
			if !r.MatchString(match) {
				t.Fatalf("%s didn’t match %s", entry.Glob, match)
			}
		}
		for _, nonmatch := range entry.NonMatches {
			if r.MatchString(nonmatch) {
				t.Fatalf("%s matched %s", entry.Glob, nonmatch)
			}
		}
	}
}
