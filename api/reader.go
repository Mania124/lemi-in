package api

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadFile(fname string) (string, error) {
	filePath, err := filepath.Abs(fname)
	if err != nil {
		return "", err
	}
	file, errr := os.Open(filePath)
	if errr != nil {
		return "", errr
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var fcontent string
	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			fcontent += scanner.Text() + "\n"
		}
	}
	return fcontent, nil
}
