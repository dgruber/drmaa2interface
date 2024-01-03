// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/dgruber/drmaa2interface

package drmaa2cue

import "time"

// JobInfo represents the state of a job.
#JobInfo: {
	Extensible:         #Extensible
	extension?:         #Extension @go(Extension)
	id?:                string     @go(ID)
	exitStatus?:        int        @go(ExitStatus)
	terminationSignal?: string     @go(TerminatingSignal)
	annotation?:        string     @go(Annotation)
	state?:             #JobState  @go(State)
	subState?:          string     @go(SubState)
	allocatedMachines?: [...string] @go(AllocatedMachines,[]string)
	submissionMachine?: string         @go(SubmissionMachine)
	jobOwner?:          string         @go(JobOwner)
	slots?:             int64          @go(Slots)
	queueName?:         string         @go(QueueName)
	wallclockTime?:     time.#Duration @go(WallclockTime)
	cpuTime?:           int64          @go(CPUTime)
	submissionTime?:    time.Time      @go(SubmissionTime)
	dispatchTime?:      time.Time      @go(DispatchTime)
	finishTime?:        time.Time      @go(FinishTime)
}
