package main

import (
	"log"

	downloadService "github.com/LocalDeployer/local_deployer/internal/downloader/service"
	installService "github.com/LocalDeployer/local_deployer/internal/installer/service"
	"github.com/spf13/viper"
)

type Config struct {
	Download struct {
		URL     string `mapstructure:"url"`
		TempDir string `mapstructure:"temp_dir"`
	} `mapstructure:"download"`
	Install struct {
		Path string `mapstructure:"path"`
	} `mapstructure:"install"`
}

func main() {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}

	downloader := downloadService.NewDownloader(downloadService.DownloaderOption{
		TempDir: config.Download.TempDir,
	})
	installer := installService.NewInstaller(installService.InstallerOption{
		InstallDir: config.Install.Path,
	})

	err := downloader.Download(config.Download.URL)
	if err != nil {
		log.Fatal(err)
	}

	err = installer.Install(config.Install.Path)
	if err != nil {
		log.Fatal(err)
	}
}
