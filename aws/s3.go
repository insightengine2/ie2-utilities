package ie2aws

import (
	"bytes"
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func S3ObjectToBuff(obj *s3.GetObjectOutput) (*bytes.Buffer, error) {

	if obj == nil {
		return nil, errors.New("unable to read s3 file contents, obj is empty")
	}

	log.Print("Attempting to read file contents from S3...")

	buffer := new(bytes.Buffer)

	if *obj.ContentLength > 0 {

		log.Printf("File found! Attempting to parse...")

		readBytes := int64(0)
		readBytes, err := buffer.ReadFrom(obj.Body)

		if err != nil {
			log.Printf("Error reading file!")
			return nil, err
		}

		log.Printf("Successfully read %d bytes.", readBytes)

	} else {
		log.Printf("Unable to read file contents!")
	}

	return buffer, nil
}

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
