package files

import (
	"feynman/config"
	"log"
	"os"
	"regexp"
)

// function to verify if the given path exists and is a file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// function to edit a given files
// the function will search and replace the variables defined in the config file
func EditFile(file config.File) (bool, error) {
	// check if the file exists
	if FileExists(file.Path) {
		// open the file & use bufio to read the file line by line
		f, err := os.ReadFile(file.Path)
		if err != nil {
			log.Fatal(err)
		}

		// convert the file to a string
		initialFile := string(f)
		log.Println(initialFile)

		// search for the variables in the file
		// if the variable is found, replace it
		// if the variable is not found, add it to the end of the file
		// if the variable is found more than once, replace all occurrences
		for _, variable := range file.Variables {
			// log.Printf("Searching for variable: %s in file %s", variable, file.Path)

			// using regexp, find the line that starts with the variable name and select the line starting from = to the end of the line
			m := regexp.MustCompile(`(?m)^` + variable + `=(.*)$`)
			// matches := m.SubexpNames()
			log.Println(m.ReplaceAllString(initialFile, variable + "=test"))
		}

	} else {
		log.Println("File does not exist:", file.Path)
		return false, nil
	}

	return true, nil
}
