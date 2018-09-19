package composefile

type ComposeFile struct {
	Region string `yaml:"region"`
	Stage string `yaml:"stage"`
	SaveToS3 bool `yaml:"saveToS3"`
	Functions []Function `yaml:"functions"`
}

type Function struct {
	Name string `yaml:"name"`
	Build []string `yaml:"build"`
	Stage string `yaml:"stage"`
}
