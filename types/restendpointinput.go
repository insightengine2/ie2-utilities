package ie2datatypes

type LambdaIntegration struct {
	LambdaName string
}

type RESTMethod struct {
	Name      string
	ReqModel  map[string]string
	ReqParams map[string]string
}

type RESTEndpointInput struct {
	ApiId       string
	ResourceId  string
	Route       string
	Stage       string
	Integration *LambdaIntegration
	Methods     []RESTMethod
}
