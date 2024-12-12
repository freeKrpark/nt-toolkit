package reader

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type TxtReader struct {
}

func (t *TxtReader) GetWords(dir string) ([]string, error) {
	txtFiles, err := t.listTxtFiles(dir)
	if err != nil {
		return nil, err
	}
	var words []string
	for _, txtFile := range txtFiles {
		word, err := func(fileName string) ([]string, error) {
			file, err := os.Open(path.Join(dir, txtFile))
			if err != nil {
				return nil, err
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			var txt []string
			for scanner.Scan() {
				txt = append(txt, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				return nil, err
			}
			return txt, nil
		}(txtFile)
		if err != nil {
			return nil, err
		}

		words = append(words, word...)
	}

	return words, nil
}

func (t *TxtReader) listTxtFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var txtFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			txtFiles = append(txtFiles, file.Name())
		}
	}

	return txtFiles, nil
}
