package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	file := os.Args[len(os.Args)-1]
	sdir := path.Dir(file)

	isTest := false
	var args []string

	if strings.HasSuffix(file, "_test.go") {
		isTest = true
		args = append(args, "test", "-c")
	} else {
		args = append(args, "build", "-o", "/dev/null")
	}

	pkg, err := build.ImportDir(sdir, build.AllowBinary)

	if err != nil {
		args = append(args, file)
	} else {
		var files []string
		files = append(files, pkg.GoFiles...)
		files = append(files, pkg.CgoFiles...)
		if isTest {
			files = append(files, pkg.TestGoFiles...)
		}

		for _, f := range files {
			args = append(args, f)
		}
	}

	cmd := exec.Command("go", args...)
	out, err := cmd.CombinedOutput()

	fmt.Print(string(out))

	if err != nil {
		os.Exit(1)
	}
}
