package service

import (
	"fmt"
	"path/filepath"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	"github.com/LocalDeployer/local_deployer/pkg/downloader"
)

type downloaderService struct {
	client  *resty.Client
	logger  *zap.Logger
	tempDir string
	status  string
}

type DownloaderOption struct {
	TempDir string
	Logger  *zap.Logger
}

func NewDownloader(opt DownloaderOption) downloader.Downloader {
	if opt.Logger == nil {
		opt.Logger, _ = zap.NewProduction()
	}

	return &downloaderService{
		client:  resty.New(),
		logger:  opt.Logger,
		tempDir: opt.TempDir,
		status:  "ready",
	}
}

func (d *downloaderService) Download(url string) error {
	d.status = "downloading"

	fileName := filepath.Base(url)
	filePath := filepath.Join(d.tempDir, fileName)

	resp, err := d.client.R().
		SetOutput(filePath).
		Get(url)

	if err != nil {
		d.status = "error"
		d.logger.Error("download failed", zap.Error(err))
		return fmt.Errorf("download failed: %w", err)
	}

	if resp.StatusCode() != 200 {
		d.status = "error"
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	d.status = "completed"
	return nil
}

func (d *downloaderService) GetStatus() string {
	return d.status
}
