package cbz

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/filesystem"
	"github.com/sergiolaverde0/mangal/key"
	"github.com/sergiolaverde0/mangal/source"
	"github.com/sergiolaverde0/mangal/util"
	"github.com/spf13/viper"
	"io"
	"os"
	"syscall"
	"time"
)

type CBZ struct{}

func (c *CBZ) SupportedConversion() (formats []constant.ConversionFormat) {
	return []constant.ConversionFormat{constant.ImageFormatWebP}
}

func New() *CBZ {
	return &CBZ{}
}

func (*CBZ) Save(chapter *source.Chapter) (string, error) {
	return save(chapter, false)
}

func (*CBZ) SaveTemp(chapter *source.Chapter) (string, error) {
	return save(chapter, true)
}

func save(chapter *source.Chapter, temp bool) (path string, err error) {
	path, err = chapter.Path(temp)
	if err != nil {
		return
	}

	err = SaveTo(chapter, path)
	if err != nil {
		return "", err
	}

	return path, nil
}

func SaveTo(chapter *source.Chapter, to string) error {

	fs := filesystem.Api()
	cbzFile, err := fs.OpenFile(to, syscall.O_WRONLY|syscall.O_CREAT|syscall.O_TRUNC, os.ModeExclusive)
	if err != nil {
		return err
	}

	defer util.Ignore(cbzFile.Close)

	zipWriter := zip.NewWriter(cbzFile)
	defer util.Ignore(zipWriter.Close)

	defer func(name string, mode os.FileMode) {
		e := fs.Chmod(name, mode)
		if e != nil {
			err = e
		}
	}(to, 0644)

	for _, page := range chapter.Pages {
		if err = addToZip(zipWriter, page.Contents, page.Filename()); err != nil {
			return err
		}
	}

	if viper.GetBool(key.MetadataComicInfoXML) {
		comicInfo := chapter.ComicInfo()
		marshalled, err := xml.MarshalIndent(comicInfo, "", "  ")
		if err == nil {
			buf := bytes.NewBuffer(marshalled)
			err = addToZip(zipWriter, buf, "ComicInfo.xml")
		}
	}
	if err != nil {
		return err
	}

	return err
}

func addToZip(writer *zip.Writer, file io.Reader, name string) error {
	header := &zip.FileHeader{
		Name:     name,
		Method:   zip.Store,
		Modified: time.Now(),
	}

	headerWriter, err := writer.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(headerWriter, file)
	return err
}
