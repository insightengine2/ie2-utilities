package ie2datatypes

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type ResearchArea struct {
	Name string `json:"name"`
}

type FileMetaData struct {
	OGFileName    string         `json:"ogfilename"`
	Synopsis      string         `json:"synopsis"`
	Authors       []Author       `json:"authors"`
	ResearchAreas []ResearchArea `json:"researchareas"`
	UploadedOn    string         `json:"uploadedon"`
	IngestedOn    string         `json:"ingestedon"`
}
