package service

import (
	"fmt"
	"go.uber.org/zap"
	"os"

	"github.com/LocalDeployer/local_deployer/pkg/installer"
)

type installerService struct {
	logger     *zap.Logger
	installDir string
	backupDir  string
	status     string
}

type InstallerOption struct {
	InstallDir string
	BackupDir  string
	Logger     *zap.Logger
}

func NewInstaller(opt InstallerOption) installer.Installer {
	if opt.Logger == nil {
		opt.Logger, _ = zap.NewProduction()
	}

	return &installerService{
		logger:     opt.Logger,
		installDir: opt.InstallDir,
		backupDir:  opt.BackupDir,
		status:     "ready",
	}
}

func (i *installerService) Install(path string) error {
	i.status = "installing"

	// 백업 디렉토리 생성
	if err := os.MkdirAll(i.backupDir, 0755); err != nil {
		i.status = "error"
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	// 설치 디렉토리 생성
	if err := os.MkdirAll(i.installDir, 0755); err != nil {
		i.status = "error"
		return fmt.Errorf("failed to create install directory: %w", err)
	}

	// 파일 복사 구현
	// ...

	i.status = "completed"
	return nil
}

func (i *installerService) Uninstall() error {
	i.status = "uninstalling"

	// 언인스톨 로직 구현
	// ...

	i.status = "uninstalled"
	return nil
}

func (i *installerService) GetStatus() string {
	return i.status
}
