package main

import (
	"fmt"
	"github.com/mitchellh/packer/packer"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func FindOvfTool() (ovftool string, err error) {
	// Accumulate any errors
	errs := new(packer.MultiError)

	// use ovftool in PATH, so use can decide which one to use
	ovftool = "ovftool"
	if _, err := exec.LookPath(ovftool); err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("ovftool not found in path: %s", err))

		files := make([]string, 0, 6)

		// search ovftool at some specific places
		files = append(files, "/Applications/VMware Fusion.app/Contents/Library/VMware OVF Tool/ovftool")

		if os.Getenv("ProgramFiles(x86)") != "" {
			files = append(files,
				filepath.Join(os.Getenv("ProgramFiles(x86)"), "/VMware/Client Integration Plug-in 5.5/ovftool/ovftool.exe"))
		}

		if os.Getenv("ProgramFiles") != "" {
			files = append(files,
				filepath.Join(os.Getenv("ProgramFiles"), "/VMware/Client Integration Plug-in 5.5/ovftool/ovftool.exe"))
		}

		if os.Getenv("ProgramFiles(x86)") != "" {
			files = append(files,
				filepath.Join(os.Getenv("ProgramFiles(x86)"), "/VMware/VMware Workstation/ovftool/ovftool.exe"))
		}

		if os.Getenv("ProgramFiles") != "" {
			files = append(files,
				filepath.Join(os.Getenv("ProgramFiles"), "/VMware/VMware Workstation/ovftool/ovftool.exe"))
		}

		file := findFile(files)
		if file == "" {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("ovftool not found: %s", err))
		} else {
			ovftool = file
		}
	}
	return
}

func findFile(files []string) string {
	for _, file := range files {
		file = normalizePath(file)
		log.Printf("Searching for file '%s'", file)

		if _, err := os.Stat(file); err == nil {
			log.Printf("Found file '%s'", file)
			return file
		}
	}

	log.Printf("File not found")
	return ""
}

func normalizePath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	path = strings.Replace(path, "//", "/", -1)
	path = strings.TrimRight(path, "/")
	return path
}
