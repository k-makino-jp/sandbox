package filereader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type fileReader interface {
	Read(filepath string)
}

type fileReaderImpl struct {
}

// pointer receiver can handle own variables
func (s *fileReaderImpl) Read(filepath string) {
	// open file
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatal("[ERROR] os.Open")
	}
	defer fp.Close()
	// scan file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		license := scanner.Text()
		fmt.Println(license)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal("[ERROR] scanner.Scan()")
	}
}

func NewFileReaderImpl() *fileReaderImpl {
	return &fileReaderImpl{}
}
