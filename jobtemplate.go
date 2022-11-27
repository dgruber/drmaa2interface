package drmaa2interface

import (
	"time"
)

// JobTemplate defines all fields for a job request defined by the DRMAA2 standard.
type JobTemplate struct {
	Extensible        `xml:"-" json:"-"`
	Extension         `json:"extension,omitempty"`
	RemoteCommand     string            `json:"remoteCommand,omitempty"`
	Args              []string          `json:"args,omitempty"`
	SubmitAsHold      bool              `json:"submitAsHold,omitempty"`
	ReRunnable        bool              `json:"reRunnable,omitempty"`
	JobEnvironment    map[string]string `json:"jobEnvironment,omitempty"`
	WorkingDirectory  string            `json:"workingDirectory,omitempty"`
	JobCategory       string            `json:"jobCategory,omitempty"`
	Email             []string          `json:"email,omitempty"`
	EmailOnStarted    bool              `json:"emailOnStarted,omitempty"`
	EmailOnTerminated bool              `json:"emailOnTerminated,omitempty"`
	JobName           string            `json:"jobName,omitempty"`
	InputPath         string            `json:"inputPath,omitempty"`
	OutputPath        string            `json:"outputPath,omitempty"`
	ErrorPath         string            `json:"errorPath,omitempty"`
	JoinFiles         bool              `json:"joinFiles,omitempty"`
	ReservationID     string            `json:"reservationID,omitempty"`
	QueueName         string            `json:"queueName,omitempty"`
	MinSlots          int64             `json:"minSlots,omitempty"`
	MaxSlots          int64             `json:"maxSlots,omitempty"`
	Priority          int64             `json:"priority,omitempty"`
	CandidateMachines []string          `json:"candidateMachines,omitempty"`
	MinPhysMemory     int64             `json:"minPhysMemory,omitempty"`
	MachineOs         string            `json:"machineOs,omitempty"`
	MachineArch       string            `json:"machineArch,omitempty"`
	StartTime         time.Time         `json:"startTime,omitempty"`
	DeadlineTime      time.Time         `json:"deadlineTime,omitempty"`
	StageInFiles      map[string]string `json:"stageInFiles,omitempty"`
	StageOutFiles     map[string]string `json:"stageOutFiles,omitempty"`
	ResourceLimits    map[string]string `json:"resourceLimits,omitempty"`
	AccountingID      string            `json:"accountingString,omitempty"`
}
