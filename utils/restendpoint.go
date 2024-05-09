package ie2utilities

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	api "github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	ie2datatypes "github.com/insightengine2/ie2-utilities/types"
)

func AWSRESTEndpointExists(conf *aws.Config, ctx *context.Context, input *ie2datatypes.RESTEndpointInput) (bool, error) {

	if conf == nil {
		e := errors.New("aws.config param can not be empty")
		return false, e
	}

	if ctx == nil {
		e := errors.New("context param can not be empty")
		return false, e
	}

	if input == nil {
		e := errors.New("input param can not be empty")
		return false, e
	}

	c := api.NewFromConfig(*conf)

	if c == nil {
		e := errors.New("failed to create apigatewayv2 client using provided config")
		return false, e
	}

	log.Printf("Checking if endpoint exists using route %s and apiid %s", input.Route, input.ApiId)

	_, err := c.GetRestApi(*ctx, &api.GetRestApiInput{
		RestApiId: aws.String(input.ApiId),
	})

	if err != nil {
		return false, err
	}

	log.Printf("API %s exists...", input.ApiId)

	/*
		c.GetResource(*ctx, &api.GetResourceInput{
			ResourceId: input.,
		})


		_, err := c.GetRoute(*ctx, &api.GetRouteInput{
			ApiId:   aws.String(input.ApiId),
			RouteId: aws.String(input.Route),
		})
	*/

	if err != nil {
		return false, err
	}

	return true, nil
}

func AWSCreateRESTEndpoint(
	conf *aws.Config,
	ctx *context.Context,
	input *ie2datatypes.RESTEndpointInput) error {

	if conf == nil {
		return errors.New("aws.config can not be empty")
	}

	if ctx == nil {
		return errors.New("context can not be empty")
	}

	if input == nil {
		return errors.New("lambdaconfig can not be empty")
	}

	/*
		c := api.NewFromConfig(*conf)

		_, e := c.CreateRoute(*ctx, &api.CreateRouteInput{

		})

		if e != nil {
			return e
		}
	*/

	return nil
}

/*
func AWSUpdateRESTEndpoint(
	conf *aws.Config,
	ctx *context.Context,
	input *ie2datatypes.LambdaInput) error {

	if conf == nil {
		return errors.New("aws.config can not be empty")
	}

	if ctx == nil {
		return errors.New("context can not be empty")
	}

	if input == nil {
		return errors.New("lambdaconfig can not be empty")
	}

	c := lambda.NewFromConfig(*conf)

	_, e := c.UpdateFunctionCode(*ctx, &lambda.UpdateFunctionCodeInput{
		Architectures: []types.Architecture{types.Architecture(input.Architecture)},
		DryRun:        input.DryRun,
		FunctionName:  aws.String(input.Name),
		Publish:       input.Publish,
		S3Bucket:      aws.String(input.S3Bucket),
		S3Key:         aws.String(input.S3Key),
	})

	if e != nil {
		return e
	}

	return nil
}

func AWSDeleteRESTEndpoint(
	conf *aws.Config,
	ctx *context.Context,
	name string) error {

	if conf == nil {
		return errors.New("aws.config can not be empty")
	}

	if ctx == nil {
		return errors.New("context can not be empty")
	}

	if len(name) <= 0 {
		return errors.New("lambda name can not be empty")
	}

	c := lambda.NewFromConfig(*conf)

	_, e := c.DeleteFunction(*ctx, &lambda.DeleteFunctionInput{
		FunctionName: aws.String(name),
	})

	if e != nil {
		return e
	}

	return nil
}
*/

func AWSGetRESTApiIdFromName(conf *aws.Config, ctx *context.Context, name string) (string, error) {

	id := ""

	if len(name) <= 0 {
		log.Print("Resource name is empty!")
		return id, errors.New("resource name is empty")
	}

	if conf == nil {
		return id, errors.New("config is empty")
	}

	if ctx == nil {
		return id, errors.New("context is empty")
	}

	name = strings.ToLower(name)
	c := api.NewFromConfig(*conf)

	out, e := c.GetRestApis(*ctx, &api.GetRestApisInput{})

	if e != nil {
		return id, e
	}

	log.Printf("Looking for a REST api named %s", name)

	for _, item := range out.Items {
		if *item.Name == name {
			log.Printf("Success! Found REST api named %s", name)
			id = *item.Id
			break
		}
	}

	if len(id) <= 0 {
		log.Printf("Failed to find a REST api named %s after searching through %d REST APIs", name, len(out.Items))
	}

	return id, nil
}

func AWSGetRESTResourceIdFromName(conf *aws.Config, ctx *context.Context, name string) (string, error) {

	id := ""

	if len(name) <= 0 {
		log.Print("Resource name is empty!")
		return id, errors.New("resource name is empty")
	}

	if conf == nil {
		return id, errors.New("config is empty")
	}

	if ctx == nil {
		return id, errors.New("context is empty")
	}

	name = strings.ToLower(name)
	c := api.NewFromConfig(*conf)

	out, e := c.GetResources(*ctx, &api.GetResourcesInput{})

	if e != nil {
		return id, e
	}

	log.Printf("Looking for a resource with a path part named %s", name)
	for _, item := range out.Items {

		log.Printf("Looking for a path part named %s in resource %s", name, *item.PathPart)

		if *item.PathPart == name {
			log.Printf("Success! Found a resource with a path part named %s", name)
			id = *item.Id
			break
		}
	}

	if len(id) <= 0 {
		log.Printf("Failed to find a resource with a path part named %s after searching %d resources", name, len(out.Items))
	}

	return id, nil
}

func AWSCreateLambdaIntegration(conf *aws.Config, ctx *context.Context, input *ie2datatypes.RESTEndpointInput) error {

	if input == nil {
		log.Printf("RESTEndpointInput value is empty.")
		return errors.New("can not create lambda integration - input value is empty")
	}

	if len(input.Integration.LambdaName) <= 0 {
		log.Printf("Lambda name is empty!")
		return errors.New("lambda name is empty")
	}

	name := input.Integration.LambdaName

	log.Printf("Attempting to create an integration for Lambda Function: %s", name)
	log.Printf("Making sure lambda '%s' exists.", name)

	exists, e := AWSLambdaExists(conf, ctx, name)

	if e != nil {
		log.Print("Failure calling AWSLambdaExists.")
		return e
	}

	if !exists {
		msg := fmt.Sprintf("lambda '%s' does NOT exist.", name)
		log.Printf("Can not create integration. Lambda function '%s' does NOT exist.", name)
		return errors.New(msg)
	}

	c := api.NewFromConfig(*conf)
	uri := fmt.Sprintf("arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:975050010293:function:%s/invocations", name)

	_, error := c.PutIntegration(*ctx, &api.PutIntegrationInput{
		HttpMethod:          aws.String("POST"),
		ResourceId:          aws.String(""),
		RestApiId:           aws.String(""),
		Type:                types.IntegrationTypeAwsProxy,
		PassthroughBehavior: aws.String("WHEN_NO_MATCH"),
		Uri:                 aws.String(uri),
	})

	return error
}
