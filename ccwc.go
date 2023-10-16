package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Flags struct {
	bytes bool
	lines bool
	words bool
	chars bool
}

func parseArguments(flags *Flags, fileName *string) {
	args := os.Args[1:]
	fileCount := 0

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			for _, runeValue := range arg[1:] {
				switch runeValue {
				case 'c':
					flags.bytes = true
				case 'w':
					flags.words = true
				case 'l':
					flags.lines = true
				case 'm':
					flags.chars = true
				default:
					fmt.Printf("Unknown flag: %c\n", runeValue)
				}
			}
		} else {
			fileCount++
			if fileCount > 1 {
				os.Stderr.WriteString("More than one filename provided.")
				os.Exit(2)
			}
			*fileName = arg
		}
	}
}

func readData(fileName *string) []byte {
	var dat []byte
	var err error
	if *fileName == "" {
		// Read input from stdin until EOF or ctrl+D is hit.
		var buf bytes.Buffer
		_, err := io.Copy(&buf, os.Stdin)
		if err != nil {
			os.Stderr.WriteString("Error reading from stdin")
			os.Exit(1)
		}
		dat = buf.Bytes()
	} else {
		dat, err = os.ReadFile(*fileName)
		if err != nil {
			os.Stderr.WriteString("Failed to read file.")
			os.Exit(1)
		}
	}
	return dat
}

func main() {
	var flags Flags
	var counts []int
	var fileName string

	parseArguments(&flags, &fileName)
	dat := readData(&fileName)

	// If no flags are set, default to bytes, words, and lines.
	if !flags.bytes && !flags.words && !flags.lines && !flags.chars {
		flags.bytes, flags.words, flags.lines = true, true, true
	}

	if flags.lines {
		lineCount := 0
		for _, b := range dat {
			if b == '\n' {
				lineCount++
			}
		}
		counts = append(counts, lineCount)
	}

	if flags.words {
		counts = append(counts, len(strings.Fields(string(dat))))
	}

	if flags.chars {
		charCount := 0
		for range string(dat) {
			charCount++
		}
		counts = append(counts, charCount)
	}

	if flags.bytes {
		counts = append(counts, len(dat))
	}

	// Output results to stdout.
	for _, value := range counts {
		fmt.Printf("  %d", value)
	}

	// File name is "" when used with pipe.
	fmt.Printf("  %s\n", fileName)
}
