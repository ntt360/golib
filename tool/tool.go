/**
* @file tool.go
* @brief toolbox
* @author ligang
* @version 1.0
* @date 2014-12-30
 */

package tool

import (
	"os"
)

func IntSliceUnique(s []int) []int {
	m := make(map[int]bool)
	r := make([]int, 0, cap(s))

	for _, k := range s {
		_, ok := m[k]
		if !ok {
			r = append(r, k)
			m[k] = true
		}
	}

	return r
}

func StringSliceUnique(s []string) []string {
	m := make(map[string]bool)
	r := make([]string, 0, cap(s))

	for _, k := range s {
		_, ok := m[k]
		if !ok {
			r = append(r, k)
			m[k] = true
		}
	}

	return r
}

func FileExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func DirExist(path string) bool {
	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	if fi.IsDir() {
		return true
	}
	return false
}
