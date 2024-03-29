// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/dgruber/drmaa2interface

package drmaa2cue

import "time"

// JobTemplate defines all fields for a job request defined by the DRMAA2 standard.
#JobTemplate: {
	extension?:     #Extension @go(Extension)
	remoteCommand?: string     @go(RemoteCommand)
	args?: [...string] @go(Args,[]string)
	submitAsHold?: bool @go(SubmitAsHold)
	reRunnable?:   bool @go(ReRunnable)
	jobEnvironment?: {[string]: string} @go(JobEnvironment,map[string]string)
	workingDirectory?: string @go(WorkingDirectory)
	jobCategory?:      string @go(JobCategory)
	email?: [...string] @go(Email,[]string)
	emailOnStarted?:    bool   @go(EmailOnStarted)
	emailOnTerminated?: bool   @go(EmailOnTerminated)
	jobName?:           string @go(JobName)
	inputPath?:         string @go(InputPath)
	outputPath?:        string @go(OutputPath)
	errorPath?:         string @go(ErrorPath)
	joinFiles?:         bool   @go(JoinFiles)
	reservationID?:     string @go(ReservationID)
	queueName?:         string @go(QueueName)
	minSlots?:          int64  @go(MinSlots)
	maxSlots?:          int64  @go(MaxSlots)
	priority?:          int64  @go(Priority)
	candidateMachines?: [...string] @go(CandidateMachines,[]string)
	minPhysMemory?: int64     @go(MinPhysMemory)
	machineOs?:     string    @go(MachineOs)
	machineArch?:   string    @go(MachineArch)
	startTime?:     time.Time @go(StartTime)
	deadlineTime?:  time.Time @go(DeadlineTime)
	stageInFiles?: {[string]: string} @go(StageInFiles,map[string]string)
	stageOutFiles?: {[string]: string} @go(StageOutFiles,map[string]string)
	resourceLimits?: {[string]: string} @go(ResourceLimits,map[string]string)
	accountingString?: string @go(AccountingID)
}
