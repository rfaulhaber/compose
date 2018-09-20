package compose

type File struct {
	Region    string              `yaml:"region"`
	Functions map[string]Function `yaml:"functions"`
}

type Function struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Build       []string          `yaml:"build"`
	Stage       string            `yaml:"stage"`
	Runtime     string            `yaml:"runtime"`
	Handler     string            `yaml:"handler"`
	Path        string            `yaml:"path"`
	Include     []string          `yaml:"include"`
	MemorySize  int64             `yaml:"memorySize"`
	Timeout     int64             `yaml:"timeout"`
	Tags        map[string]string `yaml:"tags"`
}

var runtimes = map[string]struct{}{
	"nodejs":         {},
	"nodejs4.3":      {},
	"nodejs6.10":     {},
	"nodejs8.10":     {},
	"java8":          {},
	"python2.7":      {},
	"python3.6":      {},
	"dotnetcore1.0":  {},
	"dotnetcore2.0":  {},
	"dotnetcore2.1":  {},
	"nodejs4.3-edge": {},
	"go1.x":          {},
}

func isValidRuntime(runtime string) bool {
	_, ok := runtimes[runtime]
	return ok
}
