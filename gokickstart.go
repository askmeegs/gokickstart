package main

import (
	"bufio"
	"fmt"
	"os"
)

// base
type Kickstart struct {
	pre       []string //all these are scripts
	post      []string
	packages  []string
	traceback []string
	onerror   []string
	keywords  []string //not a script!
}

//********** KICKSTART PARSER *************************************
//based on: https://github.com/rhinstaller/pykickstart/blob/master/pykickstart/parser.py
func ValidateKickstart(filename string) (bool, []error) {
	//lines := readFile(filename) //read in local text file

	return false, nil
}

//************ HELPERS *********************************************
func readFile(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	fmt.Println("Welcome to the Golang Kickstart Validator.")
	lines := readFile("kickstart1")
	fmt.Println(lines)
}
