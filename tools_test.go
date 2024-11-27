package nttoolkit

import (
	"testing"
)

func Test_ParseNTLine(t *testing.T) {
	var theTests = []struct {
		name    string
		line    string
		isError bool
	}{
		{"success", `<http://lod.kipo.kr/kb/patent/resource/KoreanTerm_45966> <http://lod.kipo.kr/kb/patent/ontology/name> "가이드 펀넬" .`, false},
	}

	for _, e := range theTests {
		t.Run(e.name, func(t *testing.T) {
			triple, err := testTools.ParseNTLine(e.line)
			if (err != nil) != e.isError {
				if e.isError {
					t.Errorf("%s: returned wrong error; expected to return an error but got no error", e.name)
				} else {
					t.Errorf("%s: returned wrong error; expected to return no error but got error: %v", e.name, err)
				}
			}
			if len(triple.Object) == 0 {
				t.Errorf("%s: expected not to be empty; but got empty", e.name)
			}
		})
	}
}

func Test_ListNTFiles(t *testing.T) {
	var theTests = []struct {
		name     string
		dir      string
		isError  bool
		isExists bool
	}{
		{"success", "testdata", false, true},
		{"success but no data", "testdata/tmp", false, false},
		{"fail", "testdata/tmp/tmp", true, false},
	}

	for _, e := range theTests {
		t.Run(e.name, func(t *testing.T) {
			testTools.RootDir = e.dir
			ntFiles, err := testTools.ListNtFiles()
			if (err != nil) != e.isError {
				if e.isError {
					t.Errorf("%s: returned wrong error; expected to return an error but got no error", e.name)
				} else {
					t.Errorf("%s: returned wrong error; expected to return no error but got error: %v", e.name, err)
				}
			}

			if (len(ntFiles) != 0) != e.isExists {
				if e.isExists {
					t.Errorf("%s: returned wrong count; expected .nt files exists but got 0", e.name)
				} else {
					t.Errorf("%s: returned wrong count; expected .nt files not exists but got %d", e.name, len(ntFiles))
				}
			}
		})
	}
}

func Test_ReadNTFiles(t *testing.T) {
	testTools.RootDir = "testdata"

	var theTests = []struct {
		name     string
		filePath string
		isError  bool
		isExists bool
	}{
		{"success", "Term_Basic_20180112135752-0000632.nt", false, true},
	}
	for _, e := range theTests {
		t.Run(e.name, func(t *testing.T) {
			triples, err := testTools.ReadNtFile(e.filePath)
			if (err != nil) != e.isError {
				if e.isError {
					t.Errorf("%s: returned wrong error; expected to return an error but got no error", e.name)
				} else {
					t.Errorf("%s: returned wrong error; expected to return no error but got error: %v", e.name, err)
				}
			}

			if (len(triples) != 0) != e.isExists {
				if e.isExists {
					t.Errorf("%s: returned wrong count; expected .nt files exists but got 0", e.name)
				} else {
					t.Errorf("%s: returned wrong count; expected .nt files not exists but got %d", e.name, len(triples))
				}
			}
		})
	}
}
