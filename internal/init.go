package internal

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func Init() error {
	if err := unix.Sethostname([]byte("fukuro")); err != nil {
		return fmt.Errorf("Setting hostname failed: %w", err)
	}
	path, err := exec.LookPath("/bin/sh")
	if err != nil {
		return fmt.Errorf("Looking for /bin/sh failed: %w", err)
	}
	return unix.Exec(path, []string{"/bin/sh"}, os.Environ())
}
