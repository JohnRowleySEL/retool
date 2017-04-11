package main

import (
	"path/filepath"
	"testing"

	"github.com/Masterminds/semver"
)

func TestSpecVersioning(t *testing.T) {
	readTest := func(file string, wantVersion *semver.Version) func(*testing.T) {
		return func(t *testing.T) {
			t.Parallel()
			path := filepath.Join("testdata", "specs", file)
			s, err := readPath(path)
			if err != nil {
				t.Fatalf("unable to read spec file: %s", err)
			}

			if wantVersion == nil {
				if s.RetoolVersion != nil {
					t.Errorf("unexpected spec retool version, have=%q want=nil", s.RetoolVersion)
				}
			} else if !s.RetoolVersion.Equal(wantVersion) {
				t.Errorf("unexpected spec retool version, have=%q want=%q", s.RetoolVersion, wantVersion)
			}
		}
	}

	t.Run("read", func(t *testing.T) {
		t.Parallel()
		t.Run("unversioned", readTest("unversioned.json", nil))
		t.Run("v1.2.0", readTest("v1.2.0.json", semver.MustParse("v1.2.0")))
	})
}
