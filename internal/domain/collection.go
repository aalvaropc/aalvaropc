package domain

// Collection defines a group of HTTP-like requests and shared metadata.
type Collection struct {
	Name     string
	Requests []RequestSpec
}

// RequestSpec describes a single request, its payload, and expectations.
type RequestSpec struct {
	Name       string
	Method     string
	URL        string
	Headers    []Header
	Body       BodySpec
	Assertions AssertionsSpec
	Extract    []ExtractSpec
}

// Header represents a header name/value pair.
type Header struct {
	Name  string
	Value string
}

// BodySpec defines the possible payload variants for a request.
type BodySpec struct {
	JSON map[string]any
	Form map[string]string
	Raw  string
}

// AssertionsSpec groups the supported assertions for a request.
type AssertionsSpec struct {
	Status          int
	MaxMilliseconds int
	JSONPathExists  []string
}

// ExtractSpec defines a JSONPath extraction and destination variable.
type ExtractSpec struct {
	JSONPath string
	Var      string
}
