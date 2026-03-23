package ie2datatypes

type AgenticFileMetaData struct {
	Abstract              string   `json:"abstract"`
	AIModel               string   `json:"aimodel"`
	Authors               []string `json:"authors"`
	CreatedOn             string   `json:"createdon"`
	Filename              string   `json:"filename"`
	Keywords              []string `json:"keywords"`
	Provider              string   `json:"provider"`
	PublishDate           string   `json:"publishdate"`
	ResearchSubCategories []string `json:"subcategories"`
	Title                 string   `json:"title"`
}
