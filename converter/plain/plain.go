package plain

import (
	"github.com/metafates/mangal/converter/webp"
	"github.com/metafates/mangal/filesystem"
	"github.com/metafates/mangal/source"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type Plain struct {
	converter *webp.Converter
}

func New() *Plain {
	return &Plain{
		converter: webp.New(),
	}
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
	page, err := p.converter.CheckAndConvert(page)
	if err != nil {
		return err
	}
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
