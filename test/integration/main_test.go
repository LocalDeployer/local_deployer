package integration

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	downloadService "github.com/LocalDeployer/local_deployer/internal/downloader/service"
	installService "github.com/LocalDeployer/local_deployer/internal/installer/service"
)

func TestIntegration(t *testing.T) {
	// 테스트 디렉토리 설정
	testDir, err := os.MkdirTemp("", "integration-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(testDir)

	logger, _ := zap.NewDevelopment()

	// 서비스 초기화
	downloader := downloadService.NewDownloader(downloadService.DownloaderOption{
		TempDir: filepath.Join(testDir, "temp"),
		Logger:  logger,
	})

	installer := installService.NewInstaller(installService.InstallerOption{
		InstallDir: filepath.Join(testDir, "install"),
		BackupDir:  filepath.Join(testDir, "backup"),
		Logger:     logger,
	})

	// 통합 테스트 시나리오
	t.Run("다운로드 후 설치", func(t *testing.T) {
		// Create mock server with test zip file
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/zip")
			w.Write([]byte("mock zip content"))
		}))
		defer mockServer.Close()

		testURL := mockServer.URL + "/test.zip"

		// 다운로드
		err := downloader.Download(testURL)
		assert.NoError(t, err)

		// 설치
		err = installer.Install(filepath.Join(testDir, "temp", "test.zip"))
		assert.NoError(t, err)

		// 상태 확인
		assert.Equal(t, "completed", installer.GetStatus())
	})
}
