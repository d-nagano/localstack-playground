package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

func main() {
	awsEndpoint := "http://localstack:4566"
	// localstack のデフォルトリージョン
	awsRegion := "us-east-1"

	// デフォルトの AccessKey、SecretKey は "test"
	creds := credentials.NewStaticCredentialsProvider("test", "test", "")
	// AWSの設定を読み込む
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(creds),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/upload", func(c echo.Context) error {
		// S3クライアント作成
		client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
			o.UsePathStyle = true
			o.BaseEndpoint = aws.String(awsEndpoint)
		})

		bucketName := "sample-bucket"
		s3Key := "sample.html"
		// HTMLファイル読み込み
		html, err := os.Open("./sample.html")
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer html.Close()

		// S3にアップロード
		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket:      aws.String(bucketName),
			Key:         aws.String(s3Key),
			Body:        html,
			ContentType: aws.String("text/html"),
		})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "Success!")
	})
	e.Logger.Fatal(e.Start(":8001"))
}
