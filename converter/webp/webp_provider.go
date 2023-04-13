package webp

import (
	"github.com/nickalie/go-binwrapper"
	"github.com/nickalie/go-webpbin"
	"image"
	"io"
)

// NewCWebP creates new CWebP instance.
func NewCWebP(folder string) *webpbin.CWebP {
	bin := &webpbin.CWebP{
		BinWrapper: createBinWrapper(folder),
	}
	bin.ExecPath("cwebp")

	return bin
}

func Encode(w io.Writer, m image.Image, quality uint) error {
	return NewCWebP(".bin/webp").
		Quality(quality).
		InputImage(m).
		Output(w).
		Run()
}

func createBinWrapper(dest string) *binwrapper.BinWrapper {
	base := "https://storage.googleapis.com/downloads.webmproject.org/releases/webp/"

	b := binwrapper.NewBinWrapper().AutoExe()
	libwebpVersion := "1.3.0"

	b.Src(
		binwrapper.NewSrc().
			URL(base + "libwebp-" + libwebpVersion + "-mac-arm64.tar.gz").
			Os("darwin").
			Arch("arm64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-mac-x86-64.tar.gz").
				Os("linux").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-x86-32.tar.gz").
				Os("linux").
				Arch("x86")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-x86-64.tar.gz").
				Os("linux").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-windows-x64.zip").
				Os("win32").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-windows-x86.zip").
				Os("win32").
				Arch("x86"))

	return b.Strip(2).Dest(dest)
}
