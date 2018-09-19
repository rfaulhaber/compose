package composefile

import "github.com/pkg/errors"

type FunctionType int

const (
	Go = iota
	JS
	Java
	CS
	Py
)

var functionTypeNames = [...]string{
	"go",
	"js",
	"java",
	"cs",
	"py",
}

var functionTypeNamesToValues = map[string]FunctionType{
	"go": Go,
	"js": JS,
	"java": Java,
	"cs": CS,
	"py": Py,
}

func (t FunctionType) String() string {
	return functionTypeNames[t]
}

func (t FunctionType) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *FunctionType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var out string

	err := unmarshal(&out)

	if err != nil {
		return errors.Wrap(err, "value could not be unmarshalled")
	}

	*t = functionTypeNamesToValues[out]

	return nil

}
