#!/bin/bash

# 버전 정보
VERSION="1.0.0"
BINARY_NAME="local_deployer"

# 빌드 디렉토리 생성
mkdir -p build

# 빌드 함수
build() {
    local GOOS=$1
    local GOARCH=$2
    local OUTPUT=$3

    echo "Building for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "build/$OUTPUT" cmd/main.go

    if [ $? -eq 0 ]; then
        echo "✅ Successfully built for $GOOS/$GOARCH"
    else
        echo "❌ Failed to build for $GOOS/$GOARCH"
        exit 1
    fi
}

# 이전 빌드 결과물 삭제
rm -rf build/*

# 각 플랫폼 별 빌드
echo "🚀 Starting multi-platform build..."

# macOS (Intel, M1)
build darwin amd64 "${BINARY_NAME}-${VERSION}-darwin-amd64"
build darwin arm64 "${BINARY_NAME}-${VERSION}-darwin-arm64"

# Linux (64-bit, ARM)
build linux amd64 "${BINARY_NAME}-${VERSION}-linux-amd64"
build linux arm64 "${BINARY_NAME}-${VERSION}-linux-arm64"

# Windows (64-bit)
build windows amd64 "${BINARY_NAME}-${VERSION}-windows-amd64.exe"

# 압축파일 생성
echo "📦 Creating archives..."
cd build
for file in *; do
    if [[ $file == *.exe ]]; then
        zip "${file%.exe}.zip" "$file"
        rm "$file"
    else
        tar czf "${file}.tar.gz" "$file"
        rm "$file"
    fi
done
cd ..

echo "✨ Build complete! Check the build directory for the binaries."