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
	AccountId        string
	Region           string
	ApiId            string
	ParentResourceId string
	ResourceId       string
	ResourceName     string
	Route            string
	Stage            string
	Integration      *LambdaIntegration
	Methods          []RESTMethod
}
