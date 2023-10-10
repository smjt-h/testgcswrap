package gcs

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// GCSClient represents a Google Cloud Storage client with a SignURL method.
type GCSClient interface {
	SignURL(bucket, object, customHost string, timeout int64) (string, error)
}

// concreteGCSClient is an implementation of the GCSClient interface.
type concreteGCSClient struct {
	client *storage.Client
}

// NewGCSClient creates a new GCS client with the provided JSON credentials file.
func NewGCSClient(ctx context.Context, credentialsFile string) (GCSClient, error) {

	// Create a GCS client with the credentials.
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}

	return &concreteGCSClient{client: client}, nil
}

// SignURL signs a URL for accessing a GCS object.
func (c *concreteGCSClient) SignURL(bucket, object, customHost string, timeout int64) (string, error) {
	url, err := c.client.Bucket(bucket).SignedURL(object, &storage.SignedURLOptions{
		Method:   "GET",
		Expires:  time.Now().Add(time.Minute * time.Duration(timeout)),
		Hostname: customHost,
	})
	if err != nil {
		return "", err
	}
	return url, nil
}
