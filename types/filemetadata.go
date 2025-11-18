package ie2datatypes

type Paper struct {
	Id            *int           `json:"id,omitempty" db:"id"`
	Title         string         `json:"title" db:"title"`
	Abstract      *string        `json:"abstract,omitempty" db:"abstract"`
	Url           *string        `json:"url,omitempty" db:"url"`
	CreatedOn     *string        `json:"createdon,omitempty" db:"createdOn"`
	UpdatedOn     *string        `json:"updatedon,omitempty" db:"updatedOn"`
	DeletedOn     *string        `json:"deletedon,omitempty" db:"deletedOn"`
	PageCnt       *int           `json:"pagecnt,omitempty"`
	ResearchAreas []ResearchArea `json:"researchareas"`
}

type Author struct {
	Id         *int    `json:"id,omitempty" db:"id"`
	FirstName  string  `json:"firstname" db:"fname"`
	MiddleName *string `json:"middlename,omitempty" db:"mname"`
	LastName   string  `json:"lastname" db:"lname"`
	Title      *string `json:"title,omitempty" db:"title"`
	IsActive   bool    `json:"isactive,omitempty" db:"isactive"`
	CreatedOn  *string `json:"createdon,omitempty" db:"createdOn"`
	UpdatedOn  *string `json:"updatedon,omitempty" db:"updatedOn"`
	DeletedOn  *string `json:"deletedon,omitempty" db:"deletedOn"`
	Papers     []Paper `json:"papers,omitempty" db:"-"`
}

type AuthorPaper struct {
	AuthorId int `json:"authorid" db:"authorid"`
	PaperId  int `json:"paperid" db:"paperid"`
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
