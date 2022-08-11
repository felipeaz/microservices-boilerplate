package pkg

import (
	"os"

	"path/filepath"
)

func GetProjectRootDirectory() string {
	currDir, _ := os.Getwd()
	return filepath.Dir(filepath.Dir(currDir))
}
