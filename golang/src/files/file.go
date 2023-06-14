package files

import (
	"bufio"
	"feynman/config"
	"fmt"
	"log"
	"os"
	"strings"
)

// function to verify if the given path exists and is a file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func EditFile(file config.File, proxyUrl string) (bool, error) {
	// check if the file exists
	if FileExists(file.Path) {
		// open the file
		f, err := os.OpenFile(file.Path, os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		fileScanner := bufio.NewScanner(f)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		for _, variable := range file.Variables {
			varExists := false

			for i, line := range fileLines {
				// check if line is empty
				if len(line) == 0 {
					continue
				}
				
				if strings.HasPrefix(line, variable) {
					varExists = true
					fileLines[i] = fmt.Sprintf("%s=%s", variable, proxyUrl)
					break
				}
			}

			if !varExists {
				fileLines = append(fileLines, fmt.Sprintf("%s=%s", variable, proxyUrl))
			}
		}

		f.Truncate(0)
		f.Seek(0, 0)

		// write the lines to the file
		datawriter := bufio.NewWriter(f)

		for _, line := range fileLines {
			_, _ = datawriter.WriteString(line + "\n")
		}

		datawriter.Flush()
	} else {
		log.Println("File does not exist:", file.Path)
		return false, nil
	}

	return true, nil
}
