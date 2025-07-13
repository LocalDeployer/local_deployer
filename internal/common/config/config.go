package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DownloaderConfig DownloaderConfig `yaml:"downloader"`
	InstallerConfig  InstallerConfig  `yaml:"installer"`
}

type DownloaderConfig struct {
	TempDir     string `yaml:"temp_dir"`
	Concurrency int    `yaml:"concurrency"`
}

type InstallerConfig struct {
	InstallDir string `yaml:"install_dir"`
	BackupDir  string `yaml:"backup_dir"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("설정 파일 읽기 실패: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("YAML 파싱 실패: %w", err)
	}

	return &cfg, nil
}
