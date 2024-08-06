package ie2datatypes

type LambdaConfig struct {
	Name         string `yaml:"name"`
	RoleName     string `yaml:"rolename"`
	Architecture string `yaml:"architecture"`
	Runtime      string `yaml:"runtime"`
	Handler      string `yaml:"handler"`
	Filename     string `yaml:"filename"`
	Endpoint     []struct {
		Version  int    `yaml:"version"`
		Resource string `yaml:"resource"`
		Methods  []struct {
			Name string `yaml:"name"`
			Req  string `yaml:"req"`
			Res  string `yaml:"res"`
		} `yaml:"methods"`
	} `yaml:"endpoint"`
}
