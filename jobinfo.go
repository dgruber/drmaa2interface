package drmaa2interface

import (
	"time"
)

// JobInfo represents the state of a job.
type JobInfo struct {
	Extensible
	Extension         `json:"extension,omitempty"`
	ID                string        `json:"id,omitempty"`
	ExitStatus        int           `json:"exitStatus,omitempty"`
	TerminatingSignal string        `json:"terminationSignal,omitempty"`
	Annotation        string        `json:"annotation,omitempty"`
	State             JobState      `json:"state,omitempty"`
	SubState          string        `json:"subState,omitempty"`
	AllocatedMachines []string      `json:"allocatedMachines,omitempty"`
	SubmissionMachine string        `json:"submissionMachine,omitempty"`
	JobOwner          string        `json:"jobOwner,omitempty"`
	Slots             int64         `json:"slots,omitempty"`
	QueueName         string        `json:"queueName,omitempty"`
	WallclockTime     time.Duration `json:"wallclockTime,omitempty"`
	CPUTime           int64         `json:"cpuTime,omitempty"`
	SubmissionTime    time.Time     `json:"submissionTime,omitempty"`
	DispatchTime      time.Time     `json:"dispatchTime,omitempty"`
	FinishTime        time.Time     `json:"finishTime,omitempty"`
}

// CreateJobInfo creates a JobInfo object where all values are initialized
// with UNSET (needed in order to differentiate if a value is
// not set or 0).
func CreateJobInfo() (ji JobInfo) {
	ji.ExitStatus = UnsetNum
	ji.Slots = UnsetNum
	ji.CPUTime = UnsetTime
	ji.State = Unset
	return ji
}
