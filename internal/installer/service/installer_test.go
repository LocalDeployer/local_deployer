package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestInstallerService(t *testing.T) {
	// 테스트용 디렉토리 생성
	testDir, err := os.MkdirTemp("", "installer-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(testDir)

	installDir := filepath.Join(testDir, "install")
	backupDir := filepath.Join(testDir, "backup")

	// 테스트 파일 생성
	testFilePath := filepath.Join(testDir, "test.txt")
	err = os.WriteFile(testFilePath, []byte("test content"), 0644)
	require.NoError(t, err)

	// 인스톨러 생성
	logger, _ := zap.NewDevelopment()
	installer := NewInstaller(InstallerOption{
		InstallDir: installDir,
		BackupDir:  backupDir,
		Logger:     logger,
	})

	t.Run("설치 테스트", func(t *testing.T) {
		err := installer.Install(testFilePath)
		assert.NoError(t, err)
		assert.Equal(t, "completed", installer.GetStatus())

		// 설치 디렉토리 확인
		assert.DirExists(t, installDir)
	})

	t.Run("언인스톨 테스트", func(t *testing.T) {
		err := installer.Uninstall()
		assert.NoError(t, err)
		assert.Equal(t, "uninstalled", installer.GetStatus())
	})
}
