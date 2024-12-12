package reader

import nttoolkit "github.com/freeKrpark/nt-toolkit"

type NtReader struct {
}

func (n *NtReader) GetData(dir string) ([]string, error) {
	var tools nttoolkit.Tools
	tools.RootDir = dir
	ntFiles, err := tools.ListNtFiles()
	if err != nil {
		return nil, err
	}
	var triples []nttoolkit.Triple
	for _, ntFile := range ntFiles {
		triple, err := tools.ReadNtFile(ntFile)
		if err != nil {
			return nil, err
		}
		triples = append(triples, triple...)
	}

	var words []string
	for _, triple := range triples {
		if triple.Object[2] != "" {
			words = append(words, triple.Object[3])
		}
	}
	return words, nil
}
