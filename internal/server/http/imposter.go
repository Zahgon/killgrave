package http

import (
	"io/fs"
)

// ImposterType allows to know the imposter type we're dealing with
type ImposterType int

const (
	jsonImposterExtension = ".imp.json"
	ymlImposterExtension  = ".imp.yml"
	yamlImposterExtension = ".imp.yaml"
)

const (
	// JSONImposter allows to know when we're dealing with a JSON imposter
	JSONImposter ImposterType = iota
	// YAMLImposter allows to know when we're dealing with a YAML imposter
	YAMLImposter
)

// ImposterConfig is used to load imposters based on which type they are
type ImposterConfig struct {
	Type     ImposterType
	FilePath string
}

// Imposter define an imposter structure
type Imposter struct {
	BasePath string    `json:"-" yaml:"-"`
	Path     string    `json:"-" yaml:"-"`
	Request  Request   `json:"request"`
	Response Responses `json:"response"`
	resIdx   int
}

// NextResponse returns the imposter's response.
// If there are multiple responses, it will return them sequentially.
func (i *Imposter) NextResponse() Response { _ = "STUB: not implemented"; return *new(Response) }

// CalculateFilePath calculate file path based on basePath of imposter's directory
func (i *Imposter) CalculateFilePath(filePath string) string { _ = "STUB: not implemented"; return "" }

// Request represent the structure of real request
type Request struct {
	Method     string             `json:"method"`
	Endpoint   string             `json:"endpoint"`
	SchemaFile *string            `json:"schemaFile"`
	Params     *map[string]string `json:"params"`
	Headers    *map[string]string `json:"headers"`
}

// Response represent the structure of real response
type Response struct {
	Status   int                `json:"status"`
	Body     string             `json:"body"`
	BodyFile *string            `json:"bodyFile" yaml:"bodyFile"`
	Headers  *map[string]string `json:"headers"`
	Delay    ResponseDelay      `json:"delay" yaml:"delay"`
}

// Responses is a wrapper for Response, to allow the use of either a single
// response or an array of responses, while keeping backwards compatibility.
type Responses []Response

func (rr *Responses) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (rr *Responses) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (rr *Responses) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (rr *Responses) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

type ImposterFs struct {
	path string
	fs   fs.FS
}

func NewImposterFS(path string) (ImposterFs, error) {
	_ = "STUB: not implemented"
	return *new(ImposterFs), nil
}

func (ifs ImposterFs) FindImposters(impostersCh chan []Imposter) error {
	_ = "STUB: not implemented"
	return nil
}

func (ifs ImposterFs) unmarshalImposters(imposterConfig ImposterConfig) ([]Imposter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
