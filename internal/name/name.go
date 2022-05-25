package name

import (
	"bufio"
	"os"
)

func GetNames(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	names := []string{}
	for fileScanner.Scan() {
		names = append(names, fileScanner.Text())
	}
	return names, nil
}
