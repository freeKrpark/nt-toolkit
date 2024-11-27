package nttoolkit

import (
	"os"
	"testing"

	"github.com/freeKrpark/toolkit/v2"
)

var testTools Tools

func TestMain(m *testing.M) {
	var dirTool toolkit.Tools
	testTools.RootDir = "testdata/tmp"
	dirTool.CreateDirIfNotExist(testTools.RootDir)
	exitCode := m.Run()
	os.Exit(exitCode)
}
