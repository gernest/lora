package imgr

import (
	"errors"
	"path/filepath"

	"github.com/disintegration/imaging"
)

const (
	THUMNAIL_SIZE = 200
)

type Thumbnails struct {
	ImageManager
}

func (t *Thumbnails) Process() error {
	if len(t.Images) == 0 {
		return errors.New("There is nothing to process")
	}
	for _, img := range t.Images {
		logThis.Info("Processing %s", img.Name)
		err := createThumbnail(img, t.Destinalion, THUMNAIL_SIZE)
		if err != nil {
			return err
		}
	}
	return nil
}

func createThumbnail(img *Image, dest string, size int) error {
	pic, err := imaging.Open(img.Path)
	if err != nil {
		return err
	}
	destImg := imaging.Thumbnail(pic, size, size, imaging.Lanczos)
	destName := img.Name + "_thumbnail" + img.Ext
	destPath := filepath.Join(dest, destName)
	err = imaging.Save(destImg, destPath)
	return err
}
