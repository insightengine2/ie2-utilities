package ie2aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func S3GetObject(conf *aws.Config, ctx *context.Context, bucket string, key string) (*s3.GetObjectOutput, error) {

	log.Printf("Create S3 client from config.")
	client := s3.NewFromConfig(*conf)

	log.Printf("Attempting to retrieve s3 object: %s", bucket+"/"+key)

	res, err := client.GetObject(*ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Printf("Error reading from S3.")
		log.Printf("Bucket: %s", bucket)
		log.Printf("Key: %s", key)
		return nil, err
	}

	return res, nil
}
