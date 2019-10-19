package main

import (
	"bufio"
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

func main() {
	if len(os.Args) == 1 {
		printPathsFromStdin()
	}

	if len(os.Args) >= 2 {
		if os.Args[1] == "-r" {
			for _, path := range os.Args[2:] {
				printAbspathsRecursively(path)
			}
		} else {
			for _, path := range os.Args[1:] {
				printAbspath(path)
			}
		}
	}
}
