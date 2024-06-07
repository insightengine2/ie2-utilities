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

/***
* Internal Functions
***/
func createApiGatewayClient(conf *aws.Config, ctx *context.Context) (*api.Client, error) {

	if conf == nil {
		e := errors.New("aws.config param can not be null")
		return nil, e
	}

	if ctx == nil {
		e := errors.New("context param can not be null")
		return nil, e
	}

	c := api.NewFromConfig(*conf)

	if c == nil {
		e := errors.New("failed to create apigatewayv2 client using provided config")
		return nil, e
	}

	return c, nil
}

func lambdaIntegrationExists(client *api.Client, ctx *context.Context, apiid string, resourceid string, in *ie2datatypes.RESTMethod) (bool, error) {

	if client == nil {
		return false, errors.New("client is null")
	}

	if ctx == nil {
		return false, errors.New("context is null")
	}

	if in == nil {
		return false, errors.New("input RESTMethod object is empty")
	}

	out, _ := client.GetIntegration(*ctx, &api.GetIntegrationInput{
		HttpMethod: aws.String(in.Name),
		ResourceId: aws.String(resourceid),
		RestApiId:  aws.String(apiid),
	})

	// if the integration does NOT exist, we get an error...
	// since I can't currently find a clean way to handle a 404 response when an integration for the given HttpMethod does NOT exist
	// i'm ignoring the error here as if it were a 404 and swallowing the error...ugh ;/

	if out != nil {
		return true, nil
	}

	return false, nil
}

func createLambdaIntegration(client *api.Client, ctx *context.Context, apiid string, resourceid string, uri string, in *ie2datatypes.RESTMethod) error {

	if client == nil {
		return errors.New("client is null")
	}

	if ctx == nil {
		return errors.New("context is null")
	}

	if in == nil {
		return errors.New("input RESTMethod object is null")
	}

	_, e := client.PutIntegration(*ctx, &api.PutIntegrationInput{
		HttpMethod:            aws.String(in.Name),
		IntegrationHttpMethod: aws.String("POST"),
		ResourceId:            aws.String(resourceid),
		RestApiId:             aws.String(apiid),
		Type:                  types.IntegrationTypeAwsProxy,
		PassthroughBehavior:   aws.String("WHEN_NO_MATCH"),
		Uri:                   aws.String(uri),
		RequestParameters:     in.ReqParams,
	})

	return e
}

func deleteLambdaIntegration(client *api.Client, ctx *context.Context, apiid string, resourceid string, in *ie2datatypes.RESTMethod) error {

	if client == nil {
		return errors.New("client is null")
	}

	if ctx == nil {
		return errors.New("context is null")
	}

	if in == nil {
		return errors.New("input RESTMethod object is null")
	}

	_, e := client.DeleteIntegration(*ctx, &api.DeleteIntegrationInput{
		HttpMethod: aws.String(in.Name),
		ResourceId: aws.String(resourceid),
		RestApiId:  aws.String(apiid),
	})

	return e
}

func createRESTMethod(client *api.Client, ctx *context.Context, apiid string, resourceid string, in *ie2datatypes.RESTMethod) error {

	if client == nil {
		return errors.New("client is null")
	}

	if ctx == nil {
		return errors.New("context is null")
	}

	if in == nil {
		return errors.New("input RESTMethod object is null")
	}

	_, e := client.PutMethod(*ctx, &api.PutMethodInput{
		ApiKeyRequired:    true,
		AuthorizationType: aws.String("NONE"),
		HttpMethod:        aws.String(in.Name),
		ResourceId:        aws.String(resourceid),
		RestApiId:         aws.String(apiid),
	})

	return e
}

func stageExists(client *api.Client, ctx *context.Context, apiid string, stage string) (bool, error) {

	if client == nil {
		return false, errors.New("client is null")
	}

	if ctx == nil {
		return false, errors.New("context is null")
	}

	if len(stage) <= 0 {
		return false, errors.New("stage value is empty")
	}

	out, e := client.GetStage(*ctx, &api.GetStageInput{
		RestApiId: aws.String(apiid),
		StageName: aws.String(stage),
	})

	if e != nil {
		return false, e
	}

	if out == nil {
		return false, nil
	}

	return true, nil
}

func createStage(client *api.Client, ctx *context.Context, apiid string, stage string, deploymentid string) error {

	if client == nil {
		return errors.New("client is null")
	}

	if ctx == nil {
		return errors.New("context is null")
	}

	if len(stage) <= 0 {
		return errors.New("stage value is empty")
	}

	_, e := client.CreateStage(*ctx, &api.CreateStageInput{
		DeploymentId: aws.String(deploymentid),
		RestApiId:    aws.String(apiid),
		StageName:    aws.String(stage),
	})

	return e
}

/***
* Exported Functions
***/
func AWSRESTMethodExists(conf *aws.Config, ctx *context.Context, apiid string, resourceid string, method *ie2datatypes.RESTMethod) (bool, error) {

	if method == nil {
		e := errors.New("method param can not be null")
		return false, e
	}

	c, e := createApiGatewayClient(conf, ctx)

	if e != nil {
		log.Print(e)
		return false, e
	}

	log.Printf("Checking if Method %s exists on API %s for Resource %s", method.Name, apiid, resourceid)
	o, e := c.GetMethod(*ctx, &api.GetMethodInput{
		HttpMethod: aws.String(method.Name),
		ResourceId: aws.String(resourceid),
		RestApiId:  aws.String(apiid),
	})

	if e != nil {
		log.Print(e)
		return false, e
	}

	if o == nil {
		return false, nil
	}

	return true, nil
}

func AWSRESTResourceExists(conf *aws.Config, ctx *context.Context, input *ie2datatypes.RESTEndpointInput) (bool, error) {

	if input == nil {
		e := errors.New("input param can not be null")
		return false, e
	}

	c, err := createApiGatewayClient(conf, ctx)

	if err != nil {
		log.Print(err)
		return false, err
	}

	_, err = c.GetResource(*ctx, &api.GetResourceInput{
		ResourceId: aws.String(input.ResourceId),
		RestApiId:  aws.String(input.ApiId),
	})

	if err != nil {
		return false, err
	}

	log.Printf("Resource %s exists...", input.ResourceId)

	return true, nil
}

func AWSRESTApiExists(conf *aws.Config, ctx *context.Context, input *ie2datatypes.RESTEndpointInput) (bool, error) {

	if input == nil {
		e := errors.New("input param can not be null")
		return false, e
	}

	c, err := createApiGatewayClient(conf, ctx)

	if err != nil {
		log.Print(err)
		return false, err
	}

	log.Printf("Checking if API exists using id %s", input.ApiId)

	_, err = c.GetRestApi(*ctx, &api.GetRestApiInput{
		RestApiId: aws.String(input.ApiId),
	})

	if err != nil {
		return false, err
	}

	log.Printf("API %s exists...", input.ApiId)

	return true, nil
}

func AWSCreateRESTResource(conf *aws.Config, ctx *context.Context, input *ie2datatypes.RESTEndpointInput) error {

	if input == nil {
		return errors.New("lambdaconfig can not be null")
	}

	if ctx == nil {
		return errors.New("context can not be null")
	}

	// does the resource already exist?
	// both the api and the resource should be present
	exists, e := AWSRESTApiExists(conf, ctx, input)

	if e != nil {
		log.Print(e)
		return e
	}

	if !exists {
		return errors.New("api does not exist")
	}

	c, e := createApiGatewayClient(conf, ctx)

	if e != nil {
		return e
	}

	exists, e = AWSRESTResourceExists(conf, ctx, input)

	if e != nil {
		return e
	}

	if !exists {

		// we need to create the REST resource
		_, e := c.CreateResource(*ctx, &api.CreateResourceInput{
			ParentId:  aws.String(input.ParentResourceId),
			PathPart:  aws.String(input.Route),
			RestApiId: aws.String(input.ApiId),
		})

		if e != nil {
			return e
		}
	}

	return nil
}

func AWSGetRESTApiIdFromName(conf *aws.Config, ctx *context.Context, name string) (string, error) {

	id := ""

	if len(name) <= 0 {
		log.Print("Resource name is empty!")
		return id, errors.New("resource name can not be empty")
	}

	if conf == nil {
		return id, errors.New("config can not be null")
	}

	if ctx == nil {
		return id, errors.New("context can not be null")
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

func AWSGetRESTResourceIdFromName(conf *aws.Config, ctx *context.Context, apiid string, name string) (string, error) {

	id := ""

	if len(apiid) <= 0 {
		log.Printf("ApiId value is empty!")
		return id, errors.New("apiid value can not be empty")
	}

	if len(name) <= 0 {
		log.Print("Resource name is empty!")
		return id, errors.New("resource name can not be empty")
	}

	if conf == nil {
		return id, errors.New("config can not be null")
	}

	if ctx == nil {
		return id, errors.New("context can not be null")
	}

	name = strings.ToLower(name)
	c := api.NewFromConfig(*conf)

	out, e := c.GetResources(*ctx, &api.GetResourcesInput{RestApiId: aws.String(apiid)})

	if e != nil {
		return id, e
	}

	log.Printf("Looking for a resource with a path part named %s", name)

	for _, item := range out.Items {

		if item.PathPart != nil {

			log.Printf("Looking for a path part named %s in resource %s", name, *item.PathPart)

			if *item.PathPart == name {
				log.Printf("Success! Found a resource with a path part named %s", name)
				id = *item.Id
				break
			}
		}
	}

	if len(id) <= 0 {
		log.Printf("Failed to find a resource with a path part named %s after searching %d resources", name, len(out.Items))
	}

	return id, nil
}

func AWSCreateLambdaIntegrations(conf *aws.Config, ctx *context.Context, input *ie2datatypes.RESTEndpointInput) error {

	if conf == nil {
		s := "config can not be null"
		log.Print(s)
		return errors.New(s)
	}

	if ctx == nil {
		s := "context can not be null"
		log.Print(s)
		return errors.New(s)
	}

	if input == nil {
		log.Printf("RESTEndpointInput value is null.")
		return errors.New("can not create lambda integration - input value is null")
	}

	if len(input.Integration.LambdaName) <= 0 {
		log.Printf("Lambda name is empty!")
		return errors.New("lambda name can not be empty")
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

	// create a client object
	c := api.NewFromConfig(*conf)

	// create the uri to the lambda function provided
	uri := fmt.Sprintf("arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:975050010293:function:%s/invocations", name)

	// iterate through each method
	// check if an integration exists
	// update if yes
	// create if no
	for _, method := range input.Methods {

		log.Printf("Checking if REST Method %s exists", method.Name)
		// does the method exist?
		exists, e := AWSRESTMethodExists(conf, ctx, input.ApiId, input.ResourceId, &method)

		if e != nil {
			log.Print(e)
			break
		}

		if !exists {

			// create the method
			e := createRESTMethod(c, ctx, input.ApiId, input.ResourceId, &method)

			if e != nil {
				log.Print(e)
				break
			}

			log.Printf("Successfully created REST method %s", method.Name)

		} else {

			log.Printf("REST Method %s exists!", method.Name)
		}

		log.Printf("Checking integration for ApiID %s ResourceId %s Method %s", input.ApiId, input.ResourceId, method.Name)
		exists, e = lambdaIntegrationExists(c, ctx, input.ApiId, input.ResourceId, &method)

		if e != nil {
			log.Print(e)
			break
		}

		if exists {

			log.Printf("%s method integration exists. Deleting existing integration.", method.Name)
			e := deleteLambdaIntegration(c, ctx, input.ApiId, input.ResourceId, &method)

			if e != nil {
				log.Print(e)
				return e
			}

			log.Printf("Succesfully deleted lambda integration for method %s", method.Name)

		} else {

			log.Printf("%s method integration does NOT exist.", method.Name)
		}

		log.Printf("Creating %s method integration for ApiID %s on Resource %s targeting Lambda %s using URI %s", method.Name, input.ApiId, input.ResourceId, name, uri)
		e = createLambdaIntegration(c, ctx, input.ApiId, input.ResourceId, uri, &method)

		if e != nil {
			log.Print(e)
			break
		}

		log.Printf("Successfully created an integration for method %s", method.Name)
	}

	log.Printf("Deploying API %s into environment %s", input.ApiId, input.Stage)
	log.Printf("Creating a new Deployment")
	newDeployment, e := c.CreateDeployment(*ctx, &api.CreateDeploymentInput{
		RestApiId: &input.ApiId,
	})

	if e != nil {
		log.Print(e)
		return e
	}

	log.Printf("Checking if stage %s exists", input.Stage)
	exists, e = stageExists(c, ctx, input.ApiId, input.Stage)

	if e != nil {
		log.Print(e)
		return e
	}

	if !exists {

		log.Printf("Stage %s does NOT exist", input.Stage)
		log.Printf("Creating stage %s", input.Stage)
		e = createStage(c, ctx, input.ApiId, input.Stage, *newDeployment.Id)

		if e != nil {
			log.Print(e)
			return e
		}

		log.Printf("Successfully created stage %s", input.Stage)

	} else {

		log.Printf("Stage %s exists!", input.Stage)
	}

	log.Printf("Updating API %s Resource %s Stage %s and DeploymentID %s", input.ApiId, input.ResourceName, input.Stage, *newDeployment.Id)
	_, e = c.UpdateRestApi(*ctx, &api.UpdateRestApiInput{
		RestApiId: &input.ApiId,
	})

	if e != nil {
		log.Print(e)
		return e
	}

	log.Printf("Successfully updated API.")

	return nil
}
