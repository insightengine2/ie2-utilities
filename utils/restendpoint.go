package ie2utilities

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	api "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
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

	_, err := c.GetRoute(*ctx, &api.GetRouteInput{
		ApiId:   aws.String(input.ApiId),
		RouteId: aws.String(input.Route),
	})

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
func AWSUpdateLambda(
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

func AWSDeleteLambda(
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
