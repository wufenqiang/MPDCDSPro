package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ProjectLocation(ProjectName string) string {
	dir, _ := os.Getwd()
	for {
		if strings.HasSuffix(dir, ProjectName) {
			pdir := filepath.Dir(dir)
			if pdir != "thrift" {
				break
			}
		}
		dir = filepath.Dir(dir)
	}
	return dir
}
