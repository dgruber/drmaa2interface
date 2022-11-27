package drmaa2interface

// Queue implements all required elements of a queue.
type Queue struct {
	Extensible
	Extension `json:"extension,omitempty"`
	Name      string `json:"name,omitempty"`
}
