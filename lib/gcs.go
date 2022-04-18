package lib

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type GCS struct {
	client     *storage.Client
	projectId  string
	bucketName string
	uploadPath string
}

type GCSConfig struct {
	ProjectId      string
	BucketName     string
	CredentialPath string
	UploadPath     string
}

func NewGoogleCloudStorage(cfg *GCSConfig) *GCS {

	// get projecct path
	gcsCredentials, _ := os.Getwd()
	gcsCredentials = fmt.Sprintf("%s/%s", gcsCredentials, cfg.CredentialPath)

	// set google applicaton credentials
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", gcsCredentials)

	// init GCS client
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &GCS{
		client:     client,
		bucketName: cfg.BucketName,
		projectId:  cfg.ProjectId,
		uploadPath: cfg.UploadPath,
	}
}

func (g *GCS) UploadFile(file multipart.File, fileName string) error {

	// create context and failed if the execution pass the limit, the limit is 50 Seconds
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := g.client.Bucket(g.bucketName).Object(g.uploadPath + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	// Close the client bucket
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func (g *GCS) Close() {

	// close GCS client
	if g.client != nil {
		g.client.Close()
	}
}
