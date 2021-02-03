package scanner

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
)

type Scanner interface {
	Scan()
	isMatchFormat()
}

type ScannerImpl struct {
	scanFilePath string
	lines        []string
}

func (s *ScannerImpl) Scan() (string, error) {
	// open file
	fp, err := os.Open(s.scanFilePath)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	// scan file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if s.isMatchFormat(line) {
			return line, errors.New("ERROR")
		}
		s.lines = append(s.lines, line)
	}
	if err = scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}

func (s *ScannerImpl) Get() []string {
	return s.lines
}

func (s *ScannerImpl) isMatchFormat(line string) bool {
	r, err := regexp.Compile(`[a-z]{4}[A-Z]{4}\d{3}`)
	if err != nil {
		log.Fatal(nil)
	}
	return r.MatchString(line)
}

func NewScannerImpl(scanFilePath string) *ScannerImpl {
	return &ScannerImpl{
		scanFilePath: scanFilePath,
	}
}
