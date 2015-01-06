package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"syscall"
)

func ConvertOfficeDocToPdf2(file string) {
	args := []string{"-f", "pdf",
		"-eSelectPdfVersion=1",
		"-eReduceImageResolution=true",
		"-eMaxImageResolution=300",
		"-p",
		"8200",
		"-o",
		"~/foo1.pdf",
		"~/foo.pptx",
	}
	cmd := exec.Command("unoconv", args...)
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		//fmt.Printf("Error:", err)
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			fmt.Printf("Failed: %d", waitStatus.ExitStatus())
		}
	} else {
		// Command was successful
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("Success: %d", waitStatus.ExitStatus())
	}
}

func ConvertOfficeDocToPdf(file string) {
	args := []string{"-f", "pdf",
		"-eSelectPdfVersion=1",
		"-eReduceImageResolution=true",
		"-eMaxImageResolution=300",
		"-p",
		"8100",
		"-o",
		"~/foo1.pdf",
		"~/foo.pptx",
	}
	path, err := exec.LookPath("unoconv")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("unoconv is available at %s\n", path)
	out, err := exec.Command("unoconv", args...).Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Success: %s\n", out)
	}

}

func GetFilenameExt(filename string) {
	extension := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(extension)]
	fmt.Println(name)
}

func main() {
	GetFilenameExt("foo.tex")
	ConvertOfficeDocToPdf("foo")
}
