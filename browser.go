package main

import (
	"os/exec"
)

func OpenBrowser(path string) error {
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
	if err != nil {
		return err
	}
	return nil
}
