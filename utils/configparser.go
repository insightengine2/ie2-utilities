package ie2utilities

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ie2datatypes "github.com/insightengine2/ie2-datatypes/types"
	"gopkg.in/yaml.v3"
)

func ConfigParser(conf *aws.Config, ctx *context.Context, bucket string, key string) (ie2datatypes.LambdaConfig, error) {

	t := ie2datatypes.LambdaConfig{}

	log.Printf("Create S3 client from config.")
	client := s3.NewFromConfig(*conf)

	log.Printf("Attempting to parse config file: %s", bucket+"/"+key)
	log.Printf("Retrieving object from S3...")

	res, err := client.GetObject(*ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Printf("Error reading from S3.")
		log.Printf("Bucket: %s", bucket)
		log.Printf("Key: %s", key)
		return t, err
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
			return t, err
		}

		if readBytes > 0 {

			log.Printf("Successfully read %d bytes.", readBytes)
			log.Printf("%s", buffer.String())
			log.Printf("Preparing to unmarshal YAML.")
			err = yaml.Unmarshal(buffer.Bytes(), &t)

			if err != nil {
				log.Printf("Error Unmarshalling config file to JSON.")
				return t, err
			}

			log.Printf("Successfully unmarshalled file!")
			log.Print(&t)
		}

	} else {
		log.Printf("Unable to read file contents!")
	}

	return t, nil
}
