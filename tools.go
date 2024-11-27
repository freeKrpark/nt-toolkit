package nttoolkit

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

type Tools struct {
	RootDir string
}

type Triple struct {
	Subject   []string
	Predicate []string
	Object    []string
}

func (t *Tools) ParseNTLine(line string) (Triple, error) {
	re := regexp.MustCompile(`<([^>]+)>|\"([^\"]+)\"`)
	matches := re.FindAllStringSubmatch(line, -1)

	if len(matches) < 3 {
		return Triple{}, fmt.Errorf("invalid NT format: %s", line)
	}

	return Triple{
		Subject:   matches[0],
		Predicate: matches[1],
		Object:    matches[2],
	}, nil
}

func (t *Tools) ListNtFiles() ([]string, error) {
	rootDir := "./"
	if t.RootDir != "" {
		rootDir = t.RootDir
	}

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	var ntFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".nt") {
			ntFiles = append(ntFiles, file.Name())
		}
	}

	return ntFiles, nil
}

func (t *Tools) ReadNtFile(fileName string) ([]Triple, error) {
	rootDir := "./"
	if t.RootDir != "" {
		rootDir = t.RootDir
	}

	if strings.HasPrefix(fileName, ".nt") {
		return nil, fmt.Errorf("only .nt files can open")
	}

	file, err := os.Open(path.Join(rootDir, fileName))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var triples []Triple
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		triple, err := t.ParseNTLine(line)
		if err != nil {
			return nil, err
		}

		triples = append(triples, triple)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return triples, nil
}
