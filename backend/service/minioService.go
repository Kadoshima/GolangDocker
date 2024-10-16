package service

import (
	"context"
)

type MinioService interface {
	// ImageSave... 画像を保存するインターフェース
	ImageSave(ctx context.Context, imageName string, imageData []byte) error
	// ImageGet... 画像を取得するインターフェース
	ImageGet(ctx context.Context, imageName string) ([]byte, error)
}
