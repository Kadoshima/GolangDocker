package service

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"mime"
	"path/filepath"
)

type MinioServiceImpl struct {
	client     *minio.Client
	bucketName string
}

// NewMinioService は MinioService の新しいインスタンスを返します
func NewMinioService(minioClient *minio.Client, bucketName string) MinioService {
	return &MinioServiceImpl{
		client:     minioClient,
		bucketName: bucketName,
	}
}

// ImageSave は画像を Minio に保存します
func (s *MinioServiceImpl) ImageSave(ctx context.Context, imageName string, imageData []byte) error {
	reader := bytes.NewReader(imageData)
	objectSize := int64(len(imageData))

	// ファイルの拡張子からコンテンツタイプを推測
	ext := filepath.Ext(imageName)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := s.client.PutObject(ctx, s.bucketName, imageName, reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Printf("画像 %s のアップロードに失敗しました: %v", imageName, err)
		return err
	}

	return nil
}

// ImageGet は Minio から画像を取得します
func (s *MinioServiceImpl) ImageGet(ctx context.Context, imageName string) ([]byte, error) {
	object, err := s.client.GetObject(ctx, s.bucketName, imageName, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("画像 %s の取得に失敗しました: %v", imageName, err)
		return nil, err
	}
	defer object.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, object); err != nil {
		log.Printf("画像 %s の読み取りに失敗しました: %v", imageName, err)
		return nil, err
	}

	return buf.Bytes(), nil
}
