package composefile

type ComposeFile struct {
	Region string `yaml:"region"`
	Stage string `yaml:"stage"`
	SaveToS3 bool `yaml:"saveToS3"`
	Functions map[string]Function `yaml:"functions"`
}

type Function struct {
	Name string `yaml:"name"`
	Description string `yaml:"description"`
	Build []string `yaml:"build"`
	Stage string `yaml:"stage"`
	Type string `yaml:"type"`
	Handler string `yaml:"handler"`
	Path string `yaml:"path"`
	MemorySize int64 `yaml:"memorySize"`
	Timeout int64 `yaml:"timeout"`
	Tags map[string]string `yaml:"tags"`
}
