package infrastructure

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func NewMinio(endpoint, accessKeyID, secretAccessKey, bucketName string, useSSL bool) (*minio.Client, error) {

	// MinIOクライアントを初期化
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	if errBucketExists != nil {
		log.Fatalln(errBucketExists)
	}

	// バケットが存在しない場合、作成
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("バケット %s が作成されました。\n", bucketName)
	}

	return minioClient, nil
}
