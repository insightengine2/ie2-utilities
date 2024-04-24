package ie2datatypes

type RESTMethod struct {
	Name      string
	ReqModel  map[string]string
	ReqParams map[string]string
}

type RESTEndpointInput struct {
	ApiId    string
	Resource string
	Route    string
	Stage    string
	Methods  []RESTMethod
}
