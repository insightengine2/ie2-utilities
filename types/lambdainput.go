package ie2datatypes

type LambdaInput struct {
	Architecture string
	DryRun       bool
	Name         string
	Handler      string
	Publish      bool
	RoleARN      string
	Runtime      string
	S3Bucket     string
	S3Key        string
}
