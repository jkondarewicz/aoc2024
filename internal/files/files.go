package files

import (
	"os"
	"strings"
)

func Open(filename string) (FileContent, error) {
	data, error := os.ReadFile(filename)
	if error != nil {
		return FileContent{}, error
	}
	fileContent := string(data)
	fileContentLineByLine := strings.Split(fileContent, "\n")
	return FileContent {
		fileContent: fileContent,
		fileContentLineByLine: fileContentLineByLine,
		Lines: len(fileContentLineByLine),
	}, nil
}

type FileContent struct {
	fileContent string
	fileContentLineByLine []string
	Lines int
}
func (fileContent FileContent) ProcessLineByLine(lineByLineProcessor LineByLineProcessor) {
	for index, line := range fileContent.fileContentLineByLine {
		lineByLineProcessor.ReadLine(index, line)
	}	
}
func (fileContent FileContent) ProcessFullContent(fullContentProcessor FullContentProcessor) {
	fullContentProcessor.ReadAll(fileContent.fileContent)
}

type LineByLineProcessor interface {
	ReadLine(index int, line string)
}
type FullContentProcessor interface {
	ReadAll(content string)
}


