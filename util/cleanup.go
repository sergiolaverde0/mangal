package util

import (
	"github.com/belphemur/mangal/provider/generic/headless"
	"os"
)

// Exit do any cleanup needed and exit the program with the given code
func Exit(code int) {
	if headless.IsLoaded() {
		transport := headless.GetTransportSingleton()
		_ = transport.Close()
	}
	os.Exit(code)
}
