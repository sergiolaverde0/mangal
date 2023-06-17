package util

import (
	"github.com/belphemur/mangal/provider/generic/headless"
	"os"
)

func Exit(code int) {
	if headless.IsLoaded() {
		transport := headless.GetTransportSingleton()
		_ = transport.Close()
	}
	os.Exit(code)
}
