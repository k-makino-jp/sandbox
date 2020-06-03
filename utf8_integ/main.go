package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func installSoftware(filePath string) {
	fileContent := fetchFile(filePath)
	contentList := strings.Split(replaceCRLFtoLF(fileContent), "\n")
	for _, content := range contentList {
		command := []string{
			"/C",
			content,
		}
		invokeCommand("cmd", command)
	}
}

// UTF-8からShift-JISに変換する関数
func UTF8ToSjis(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewEncoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

// ShiftJISからUTF-8に変換する関数
func sjisToUTF8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

func invokeCommand(command string, commandArgs []string) {
	cmd := exec.Command(command, commandArgs...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	if utf8.Valid(out) {
		fmt.Println(string(out))
	} else {
		fmt.Println("=== ShiftJIS to UTF8")
		decodedData, _ := sjisToUTF8(string(out))
		fmt.Println(decodedData)
	}
}

func replaceCRLFtoLF(contents string) string {
	return strings.Replace(contents, "\r\n", "\n", -1)
}

func fetchFile(filePath string) string {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if utf8.Valid(bytes) {
		// UTF-8
		return string(bytes)
	} else {
		// Shift-JIS to UTF-8
		utf8Str, _ := sjisToUTF8(string(bytes))
		return utf8Str
	}
}

func main() {
	fmt.Println("=== [utf8] ===============================================")
	installSoftware("utf8.bat")
	fmt.Println("=== [sjis] ===============================================")
	installSoftware("sjis.bat")
	fmt.Println("=== [sjisCallUtf8] =======================================")
	installSoftware("sjisCallUtf8.bat")
}
