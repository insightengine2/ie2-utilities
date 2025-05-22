package ie2datatypes

type Paper struct {
	PageCnt       *int           `json:"pagecnt,omitempty"`
	Title         string         `json:"title"`
	ResearchAreas []ResearchArea `json:"researchareas"`
}

type Author struct {
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Papers    []Paper `json:"papers,omitempty"`
}

type ResearchArea struct {
	Name string `json:"name"`
}

type FileMetaData struct {
	Authors       []Author       `json:"authors"`
	IngestedOn    string         `json:"ingestedon"`
	OGFileName    string         `json:"ogfilename"`
	ResearchAreas []ResearchArea `json:"researchareas"`
	Synopsis      string         `json:"synopsis"`
	UploadedOn    string         `json:"uploadedon"`
}
