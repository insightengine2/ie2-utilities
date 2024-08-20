package ie2datatypes

type FileMetaData struct {
	OGFileName string `json:"ogfilename"`
	Synopsis   string `json:"synopsis"`
	Authors    []struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	} `json:"authors"`
	ResearchAreas []struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"researchareas"`
	UploadedOn string `json:"uploadedon"`
	IngestedOn string `json:"ingestedon"`
}
