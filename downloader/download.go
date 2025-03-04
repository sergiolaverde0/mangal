package downloader

import (
	"encoding/json"
	"fmt"
	"github.com/sergiolaverde0/mangal/color"
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/converter"
	"github.com/sergiolaverde0/mangal/filesystem"
	"github.com/sergiolaverde0/mangal/history"
	"github.com/sergiolaverde0/mangal/key"
	"github.com/sergiolaverde0/mangal/log"
	"github.com/sergiolaverde0/mangal/packer"
	"github.com/sergiolaverde0/mangal/source"
	"github.com/sergiolaverde0/mangal/style"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Download the chapter using given source.
func Download(chapter *source.Chapter, progress func(string)) (string, error) {
	log.Info("downloading " + chapter.Name)

	path, err := chapter.Path(false)
	if err != nil {
		return "", err
	}

	if viper.GetBool(key.DownloaderRedownloadExisting) {
		log.Info("chapter already downloaded, deleting and redownloading")
		err = filesystem.Api().Remove(path)
		if err != nil {
			log.Warn(err)
		}
	} else {
		log.Info("checking if chapter is already downloaded")
		if chapter.IsDownloaded() {
			log.Info("chapter already downloaded, skipping")
			return path, nil
		}
	}

	progress("Getting pages")
	err = chapter.Source().LoadPagesOf(chapter)
	if err != nil {
		log.Error(err)
		return "", err
	}
	log.Info("found " + fmt.Sprintf("%d", len(chapter.Pages)) + " pages")

	err = chapter.DownloadPages(false, progress)
	if err != nil {
		log.Error(err)
		return "", err
	}

	if viper.GetBool(key.MetadataFetchAnilist) {
		err := chapter.Manga.PopulateMetadata(progress)
		if err != nil {
			log.Warn(err)
		}
	}

	if viper.GetBool(key.MetadataSeriesJSON) {
		path, err := chapter.Manga.Path(false)
		if err != nil {
			log.Warn(err)
		} else {
			path = filepath.Join(path, "series.json")
			progress("Generating series.json")
			seriesJSON := chapter.Manga.SeriesJSON()
			buf, err := json.Marshal(seriesJSON)
			if err != nil {
				log.Warn(err)
			} else {
				err = filesystem.Api().WriteFile(path, buf, os.ModePerm)
				if err != nil {
					log.Warn(err)
				}
			}
		}
	}

	if viper.GetBool(key.DownloaderDownloadCover) {
		coverDir, err := chapter.Manga.Path(false)
		if err == nil {
			_ = chapter.Manga.DownloadCover(false, coverDir, progress)
		}
	}

	pack, err := packer.Get(viper.GetString(key.FormatsUse))
	if err != nil {
		log.Error(err)
		return "", err
	}

	conversionFormat := constant.ConversionFormat(viper.GetString(key.ImageConversion))
	if conversionFormat != constant.ImageFormatUnknown && conversionFormat != constant.ImageFormatNone && lo.Contains(pack.SupportedConversion(), conversionFormat) {
		log.Info("getting " + conversionFormat + " converter")
		progress(fmt.Sprintf(
			"Converting %d pages to %s",
			len(chapter.Pages),
			style.Fg(color.Yellow)(string(conversionFormat))),
		)
		conv, err := converter.Get(conversionFormat)
		if err != nil {
			log.Error(err)
			return "", err
		}
		chapter, err = conv.ConvertChapter(chapter, uint8(viper.GetUint(key.ImageConversionQuality)), progress)
		if err != nil {
			log.Error(err)
			return "", err
		}
	}

	log.Info("getting " + viper.GetString(key.FormatsUse) + " packer")
	progress(fmt.Sprintf(
		"Packing %d pages to %s %s",
		len(chapter.Pages),
		style.Fg(color.Yellow)(viper.GetString(key.FormatsUse)),
		style.Faint(chapter.SizeHuman())),
	)

	log.Info("converting " + viper.GetString(key.FormatsUse))
	path, err = pack.Save(chapter)
	if err != nil {
		log.Error(err)
		return "", err
	}

	if viper.GetBool(key.HistorySaveOnDownload) {
		go func() {
			err = history.Save(chapter)
			if err != nil {
				log.Warn(err)
			} else {
				log.Info("history saved")
			}
		}()
	}

	log.Info("downloaded without errors")
	progress("Downloaded")
	return path, nil
}
