/*
Package whodis provides a function to return the URL of the current file and line number in a git(hub) repo.

This is useful for printing a link to the source code of the calling function, for example in error messages or templates.

Usage:

import (

	"fmt"
	"github.com/gwillem/go-whodis"

)

	func MyFunc() {
		fmt.Println("This mail was sent from", whodis.URL())
		// Prints: https://github.com/gwillem/someapp/blob/HEAD/pkg/pkg.go#L11
	}
*/
package whodis

import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
)

const (
	slash          = string(filepath.Separator)
	maxStackFrames = 100
)

func findCommonRoot(path1, path2 string) string {
	c1 := strings.Split(filepath.Clean(path1), slash)
	c2 := strings.Split(filepath.Clean(path2), slash)

	commonRoot := ""
	minComponents := len(c1)
	if len(c2) < minComponents {
		minComponents = len(c2)
	}

	for i := 0; i < minComponents; i++ {
		if c1[i] != c2[i] {
			break
		}
		commonRoot = filepath.Join(commonRoot, c1[i])
	}

	return slash + commonRoot
}

func findRootDir() string {
	root := ""
	for i := 0; i < maxStackFrames; i++ {
		_, file, _, _ := runtime.Caller(i)
		if file == "" {
			break
		}

		dir := filepath.Dir(file)

		if root == "" {
			root = dir
			continue
		}

		common := findCommonRoot(root, dir)
		if common == slash {
			break
		}
		if len(common) < len(root) {
			root = dir
		}
	}
	return root
}

func URL() string {
	_, fname, line, _ := runtime.Caller(1)
	bi, _ := debug.ReadBuildInfo()
	root := findRootDir()
	relPath := fname[len(root)+1:]

	// github syntax, may also work for other repo hosts?
	return fmt.Sprintf("https://%s/blob/HEAD/%s#L%d",
		bi.Path, relPath, line)
}
