package errors

type Model struct {
	Code    uint32   `json:"code,omitempty" xml:"code,omitempty" yaml:"code,omitempty"`
	Message string   `json:"message,omitempty" xml:"message,omitempty" yaml:"message,omitempty"`
	Details []string `json:"details,omitempty" xml:"details,omitempty" yaml:"details,omitempty"`
}
