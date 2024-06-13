package ie2utilities

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	ie2datatypes "github.com/insightengine2/ie2-utilities/types"
)

func AWSLambdaExists(conf *aws.Config, ctx *context.Context, name string) (bool, error) {

	if len(name) <= 0 {
		e := errors.New("function name can not be empty")
		return false, e
	}

	if conf == nil {
		e := errors.New("aws.config can not be empty")
		return false, e
	}

	if ctx == nil {
		e := errors.New("context can not be empty")
		return false, e
	}

	c := lambda.NewFromConfig(*conf)

	if c == nil {
		e := errors.New("failed to create lambda client using provided config")
		return false, e
	}

	_, err := c.GetFunction(*ctx, &lambda.GetFunctionInput{
		FunctionName: aws.String(name),
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}

func AWSCreateLambda(
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

	_, e := c.CreateFunction(*ctx, &lambda.CreateFunctionInput{
		Architectures: []types.Architecture{types.Architecture(input.Architecture)},
		Code: &types.FunctionCode{
			S3Bucket: aws.String(input.S3Bucket),
			S3Key:    aws.String(input.S3Key),
		},
		FunctionName: aws.String(input.Name),
		Handler:      aws.String(input.Handler),
		Publish:      *aws.Bool(input.Publish),
		Role:         aws.String(input.RoleARN),
		Runtime:      types.Runtime(*aws.String(input.Runtime)),
	})

	if e != nil {
		return e
	}

	return nil
}

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

func AWSAddApiGatewayPermission(
	conf *aws.Config,
	ctx *context.Context,
	method string,
	sourcearn string,
	lambdaname string) error {

	if conf == nil {
		return errors.New("aws.config can not be empty")
	}

	if ctx == nil {
		return errors.New("context can not be empty")
	}

	if len(lambdaname) <= 0 {
		return errors.New("lambdaname can not be empty")
	}

	if len(method) <= 0 {
		return errors.New("method can not be empty")
	}

	c := lambda.NewFromConfig(*conf)

	_, e := c.AddPermission(*ctx, &lambda.AddPermissionInput{
		Action:       aws.String("lambda:InvokeFunction"),
		FunctionName: aws.String(lambdaname),
		Principal:    aws.String("apigateway.amazonaws.com"),
		StatementId:  aws.String(fmt.Sprintf("AllowAPIMethod%sOnFunction%s", method, lambdaname)),
		SourceArn:    aws.String(sourcearn),
	})

	if e != nil {
		return e
	}

	return nil
}
