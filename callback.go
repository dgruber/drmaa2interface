package drmaa2interface

type EventChannel <-chan Notification

// Event specifies the type of the event
type Event int

const (
	NewState Event = iota
	Migrated
	AttributeChange
)

// Notification is the argument of the callback function
// automatically called for an event.
type Notification struct {
	Evt         Event    `json:"event"`
	JobId       string   `json:"jobId"`
	SessionName string   `json:"sessionName"`
	State       JobState `json:"jobState"`
}
