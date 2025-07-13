package service

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestDownloaderService(t *testing.T) {
	// 임시 디렉토리 생성
	tempDir, err := os.MkdirTemp("", "downloader-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// 테스트 서버 설정
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test content"))
	}))
	defer ts.Close()

	// 다운로더 생성
	logger, _ := zap.NewDevelopment()
	downloader := NewDownloader(DownloaderOption{
		TempDir: tempDir,
		Logger:  logger,
	})

	tests := []struct {
		name          string
		url           string
		expectedError bool
	}{
		{
			name:          "정상 다운로드",
			url:           ts.URL + "/test.txt",
			expectedError: false,
		},
		{
			name:          "잘못된 URL",
			url:           "http://invalid-url",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := downloader.Download(tt.url)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "completed", downloader.GetStatus())

				// 파일 내용 확인
				content, err := os.ReadFile(filepath.Join(tempDir, "test.txt"))
				require.NoError(t, err)
				assert.Equal(t, "test content", string(content))
			}
		})
	}
}
