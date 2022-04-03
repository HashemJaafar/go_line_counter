package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

const _test = "_test.go"

var (
	golang_files      []string
	golang_test_files []string
)

func read_folders(folder_name string) {
	files, err := ioutil.ReadDir(folder_name)
	err_panic(err)
	for _, f := range files {
		dir := folder_name + "/" + f.Name()
		switch {
		case f.IsDir():
			read_folders(dir)
		case filepath.Ext(f.Name()) == ".go":
			if strings.Contains(f.Name(), _test) {
				golang_test_files = append(golang_test_files, dir)
			} else {
				golang_files = append(golang_files, dir)
			}
		}
	}
}

func file_lines(file_name string) []string {
	content, err := ioutil.ReadFile(file_name)
	err_panic(err)
	return strings.Split(string(content), "\n")
}

func err_panic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	var number_line_golang int
	var number_line_golang_test int
	var number_file_golang int
	var number_file_golang_test int
	read_folders(".")
	for _, f := range golang_files {
		file_lines := file_lines(f)
		number_line_golang += len(file_lines)
		number_file_golang++
	}
	for _, f := range golang_test_files {
		file_lines := file_lines(f)
		number_line_golang_test += len(file_lines)
		number_file_golang_test++
	}

	p := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(p, "number_line_golang\t:", number_line_golang)
	fmt.Fprintln(p, "number_line_golang_test\t:", number_line_golang_test)
	fmt.Fprintln(p, "number_line_golang_all\t:", number_line_golang+number_line_golang_test)
	fmt.Fprintln(p, "number_file_golang\t:", number_file_golang)
	fmt.Fprintln(p, "number_file_golang_test\t:", number_file_golang_test)
	fmt.Fprintln(p, "number_file_golang_all\t:", number_file_golang+number_file_golang_test)
	p.Flush()
}
