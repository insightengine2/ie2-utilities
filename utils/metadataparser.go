package ie2utilities

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ie2datatypes "github.com/insightengine2/ie2-utilities/types"
	"gopkg.in/yaml.v3"
)

// TODO: This function is so similar to configparser, can we abstract and consolidate the functionality? -gb 8.21.24
func MetaDataParser(conf *aws.Config, ctx *context.Context, bucket string, key string) (ie2datatypes.FileMetaData, error) {

	ret := ie2datatypes.FileMetaData{}

	log.Printf("Create S3 client from config.")
	client := s3.NewFromConfig(*conf)

	log.Printf("Attempting to parse metadata file: %s", bucket+"/"+key)
	log.Printf("Retrieving object from S3...")

	res, err := client.GetObject(*ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Printf("Error reading from S3.")
		log.Printf("Bucket: %s", bucket)
		log.Printf("Key: %s", key)
		return ret, err
	}

	log.Print("GetObject result...")
	log.Print(*res)
	log.Printf("Unstarred val: %d", res.ContentLength)
	log.Printf("File content length: %d", *res.ContentLength)

	if *res.ContentLength > 0 {

		log.Printf("File found! Attempting to parse...")

		readBytes := int64(0)
		buffer := new(bytes.Buffer)
		readBytes, err = buffer.ReadFrom(res.Body)

		if err != nil {
			log.Printf("Error reading file!")
			return ret, err
		}

		if readBytes > 0 {

			log.Printf("Successfully read %d bytes.", readBytes)
			log.Printf("%s", buffer.String())
			log.Printf("Preparing to unmarshal JSON.")
			err = yaml.Unmarshal(buffer.Bytes(), &ret)

			if err != nil {
				log.Printf("Error Unmarshalling config file to JSON.")
				return ret, err
			}

			log.Printf("Successfully unmarshalled file!")
			log.Print(&ret)
		}

	} else {
		log.Printf("Unable to read file contents!")
	}

	return ret, nil
}
