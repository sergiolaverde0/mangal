package plain

import (
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/filesystem"
	"github.com/sergiolaverde0/mangal/source"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type Plain struct {
}

func (p *Plain) SupportedConversion() []constant.ConversionFormat {
	return []constant.ConversionFormat{constant.ImageFormatWebP}
}

func New() *Plain {
	return &Plain{}
}

func (p *Plain) Save(chapter *source.Chapter) (string, error) {
	return p.save(chapter, false)
}

func (p *Plain) SaveTemp(chapter *source.Chapter) (string, error) {
	return p.save(chapter, true)
}

func (p *Plain) save(chapter *source.Chapter, temp bool) (path string, err error) {
	path, err = chapter.Path(temp)
	if err != nil {
		return
	}

	err = filesystem.Api().Mkdir(path, os.ModePerm)
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(chapter.Pages))
	for _, page := range chapter.Pages {
		func(page *source.Page) {
			defer wg.Done()

			if err != nil {
				return
			}

			err = p.savePage(page, path)
		}(page)
	}

	wg.Wait()
	return
}

func (p *Plain) savePage(page *source.Page, to string) error {
	file, err := filesystem.Api().Create(filepath.Join(to, page.Filename()))
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	_, err = io.Copy(file, page.Contents)
	if err != nil {
		return err
	}

	_ = file.Close()
	_ = page.Close()

	return nil
}
