package scanner

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

var (
	regexpCompile       = regexp.Compile
	osOpen              = os.Open
	scannerMaxTokenSize = bufio.MaxScanTokenSize

	errInvalidLine     = errors.New("Invalid line error occurred")
	errDuplicatedLines = errors.New("Duplicated lines are detected")
	errFileNotExist    = errors.New("File not exist")
	errCannotOpenFile  = errors.New("Cannnot open file")
	errScanFailed      = errors.New("Scan failed")
)

type Scanner interface {
	Get() []string
	Scan() error
	lineValidation(string) bool
}

type ScannerImpl struct {
	validLines []string
}

func (s *ScannerImpl) lineValidation(line string) error {
	r, err := regexpCompile(`^[A-Z\d]{8}-[A-Z]{2}[A-Z\d]{2}C\d{2}$`)
	if err != nil {
		return err
	}
	if !r.MatchString(line) {
		return errInvalidLine
	}
	return nil
}

func (s *ScannerImpl) detectDuplicate(lines []string) error {
	encounterd := map[string]int{}
	for _, line := range lines {
		encounterd[line]++
	}
	for k, v := range encounterd {
		if v > 1 {
			fmt.Println("Duplicated Line:", k)
			return errDuplicatedLines
		}
	}
	return nil
}

func (s *ScannerImpl) Get() []string {
	return s.validLines
}

func (s *ScannerImpl) Scan(scanFilePath string) error {
	// is file exist
	if _, err := os.Stat(scanFilePath); err != nil {
		fmt.Println(err)
		return errFileNotExist
	}
	// open file
	fp, err := osOpen(scanFilePath)
	if err != nil {
		fmt.Println(err)
		return errCannotOpenFile
	}
	defer fp.Close()
	// scan file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if err = s.lineValidation(line); err != nil {
			return err
		}
		s.validLines = append(s.validLines, line)
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return errScanFailed
	}
	if err = s.detectDuplicate(s.validLines); err != nil {
		return err
	}
	return nil
}

func NewScannerImpl() *ScannerImpl {
	return &ScannerImpl{}
}
