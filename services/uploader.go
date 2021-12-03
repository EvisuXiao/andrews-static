package services

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
	"upload-test/config"
	"upload-test/pkg/transfer"
	"upload-test/pkg/utils"
	"upload-test/types"
)

type Uploader struct {
	fileHeader *multipart.FileHeader
	tempFile   types.FileBrief
	filters    []*UploadFilter
}

type UploadFilter struct {
	name    string
	handler UploadFilterHandler
	File    types.FileBrief
}
type UploadFilterHandler func(types.FileBrief, []*UploadFilter) (types.FileBrief, error)

func NewUploader(fileHeader *multipart.FileHeader) *Uploader {
	uploader := &Uploader{fileHeader: fileHeader}
	uploader.tempFile.BaseDir = config.TempFilePath("")
	uploader.tempFile.Filename = fileHeader.Filename
	uploader.tempFile.FormatFromFilename()
	uploader.tempFile.Basename = fmt.Sprintf("%s%d", uploader.tempFile.Basename, time.Now().Unix())
	uploader.tempFile.FormatFromBasename()
	return uploader
}

func (u *Uploader) Upload() error {
	if err := u.check(); !utils.IsEmpty(err) {
		return err
	}
	if err := u.saveTmpFile(); !utils.IsEmpty(err) {
		return err
	}
	if err := u.hashFilename(); !utils.IsEmpty(err) {
		return err
	}
	for _, filter := range u.filters {
		filteredFile, err := filter.handler(u.tempFile, u.filters)
		if !utils.IsEmpty(err) {
			return err
		}
		filter.File = filteredFile
	}
	return u.transferFiles()
}

func (u *Uploader) RegisterFilter(name string, filter UploadFilterHandler) {
	u.filters = append(u.filters, &UploadFilter{name: name, handler: filter})
}

func (u *Uploader) check() error {
	if u.tempFile.InvalidMedia() {
		return fmt.Errorf("invalid media file: %s", u.fileHeader.Filename)
	}
	if !CheckFileSize(u.fileHeader.Size, u.tempFile.MediaType) {
		return fmt.Errorf("media file size is too big: %d", u.fileHeader.Size)
	}
	return nil
}

func (u *Uploader) saveTmpFile() error {
	src, err := u.fileHeader.Open()
	if !utils.IsEmpty(err) {
		return err
	}
	defer src.Close()
	out, err := os.Create(u.tempFile.FilePath)
	if !utils.IsEmpty(err) {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}

func (u *Uploader) hashFilename() error {
	h, err := utils.EncodeMd5File(u.tempFile.FilePath)
	if !utils.IsEmpty(err) {
		return err
	}
	tmpPath := u.tempFile.FilePath
	u.tempFile.Basename = h[:10]
	u.tempFile.FormatFromBasename()
	return os.Rename(tmpPath, u.tempFile.FilePath)
}

func (u *Uploader) transferFiles() error {
	for _, f := range u.filters {
		_ = u.transferFile(f.File)
	}
	return u.transferFile(u.tempFile)
}

func (u *Uploader) transferFile(file types.FileBrief) error {
	if utils.IsEmpty(file.FilePath) {
		return errors.New("empty file")
	}
	defer os.Remove(file.FilePath)
	return transfer.LocalToRemote(file.FilePath, WrapUploadPath(file.Filename, file.MediaType), false)
}

func WrapUploadPath(filename string, mediaType types.MEDIA_TYPE) string {
	return fmt.Sprintf("%s/%s", mediaType, filename)
}

func CheckFileSize(size int64, mediaType types.MEDIA_TYPE) bool {
	switch mediaType {
	case types.MEDIA_TYPE_IMAGE:
		return size <= config.UploadConfig.ImageMaxSize
	case types.MEDIA_TYPE_GIF:
		return size <= config.UploadConfig.ImageMaxSize
	case types.MEDIA_TYPE_VIDEO:
		return size <= config.UploadConfig.VideoMaxSize
	case types.MEDIA_TYPE_AUDIO:
		return size <= config.UploadConfig.AudioMaxSize
	default:
		return false
	}
}
