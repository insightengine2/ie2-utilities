package ie2utilities

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func AWSGetAccountId(conf *aws.Config, ctx *context.Context) (string, error) {

	c := sts.NewFromConfig(*conf)

	res, err := c.GetCallerIdentity(*ctx, &sts.GetCallerIdentityInput{})

	if err != nil {
		return "", err
	}

	return *res.Account, nil
}
