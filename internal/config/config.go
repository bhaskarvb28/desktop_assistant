package config

import (
	"os"
)

func JarvisRoot() string {

	if root := os.Getenv("JARVIS_ROOT"); root != "" {
		return root
	}

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return wd
}