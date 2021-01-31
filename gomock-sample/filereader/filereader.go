package filereader

import (
	"bufio"
	"fmt"
	"os"
)

type FileReader interface {
	Read(filepath string) error
}

type fileReaderInstance struct {
	FileReader
}

type fileReaderImpl struct {
}

func NewFileReaderInstance() *fileReaderInstance {
	return &fileReaderInstance{
		&fileReaderImpl{},
	}
}

// pointer receiver can handle own variables
func (s *fileReaderImpl) Read(filepath string) error {
	// open file
	fp, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fp.Close()
	// scan file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		license := scanner.Text()
		fmt.Println(license)
	}

	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}
