package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// printAbspath prints the absolute path of the given (relative) path.
// It prints an error message if the path can't be parsed.
func printAbspath(path string) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse file or directory path '%s': %v\n", path, err)
		return
	}
	fmt.Println(abspath)
}

// printPathsFromStdin reads all paths from STDIN (one per line)
// and prints their absolute paths.
func printPathsFromStdin() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		printAbspath(input.Text())
	}
}

// visit takes a relative path and prints its absolute path.
// It is a helper function used for printing absolute paths recursively
// with filepath.Walk().
func visit(path string, f os.FileInfo, err error) error {
	printAbspath(path)
	return nil
}

// printAbspathsRecursively prints the absolute path for the given relative
// path. If the given path is a directory, recursively print the absolute
// paths of all files and directories in it.
func printAbspathsRecursively(path string) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Path '%s' is not a file/folder: %v\n", path, err)
	} else if fileInfo.IsDir() {
		filepath.Walk(path, visit)
	} else {
		printAbspath(path)
	}
}

func printUsage() {
	fmt.Printf("Usage of %s:\n\n", os.Args[0])

	fmt.Printf("\tabspath file1.txt path/to/file2.pdf\n")
	fmt.Printf("\tabspath Desktop/*\n")
	fmt.Printf("\tabspath -r Desktop\n")
	fmt.Printf("\tfind . -name *.pdf | abspath\n\n")

	flag.PrintDefaults()
}

func main() {
	flag.Usage = printUsage

	var recursive bool
	flag.BoolVar(&recursive, "r", false,
		"recursive: print absolute paths of all files in the directory structure")
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 { // no paths given as command-line arguments
		printPathsFromStdin()
		os.Exit(0)
	}

	if recursive {
		for _, path := range paths {
			printAbspathsRecursively(path)
		}
	} else {
		for _, path := range paths {
			printAbspath(path)
		}
	}
}
