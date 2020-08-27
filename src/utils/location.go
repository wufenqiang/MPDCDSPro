package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ProjectLocation(ProjectName string) string {
	dir, _ := os.Getwd()
	tmpdir := dir
	for {
		if strings.HasSuffix(dir, ProjectName) {
			pdir := filepath.Dir(dir)
			if pdir != "thrift" {
				break
			}
		}
		dir = filepath.Dir(dir)
		if tmpdir == dir {
			panic("无法找到本工程")
			break
		} else {
			tmpdir = dir
		}
	}
	return dir
}
