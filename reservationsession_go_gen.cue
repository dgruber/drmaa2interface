// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/dgruber/drmaa2interface

package drmaa2cue

import "time"

// ReservationInfo contains all details of the current state of
// a resource reservation.
#ReservationInfo: {
	Extensible:            #Extensible
	extension?:            #Extension @go(Extension)
	reservationID?:        string     @go(ReservationID)
	reservationName?:      string     @go(ReservationName)
	reservationStartTime?: time.Time  @go(ReservationStartTime)
	reservationEndTime?:   time.Time  @go(ReservationEndTime)
	acl?: [...string] @go(ACL,[]string)
	reservedSlots?: int64 @go(ReservedSlots)
	reservedMachines?: [...string] @go(ReservedMachines,[]string)
}

// ReservationTemplate contains ressource requests for a
// resource reservation.
#ReservationTemplate: {
	Extensible:   #Extensible
	extension?:   #Extension     @go(Extension)
	name?:        string         @go(Name)
	startTime?:   time.Time      @go(StartTime)
	endTime?:     time.Time      @go(EndTime)
	duration?:    time.#Duration @go(Duration)
	minSlots?:    int64          @go(MinSlots)
	maxSlots?:    int64          @go(MaxSlots)
	jobCategory?: string         @go(JobCategory)
	userACL?: [...string] @go(UsersACL,[]string)
	candidateMachines?: [...string] @go(CandidateMachines,[]string)
	minPhysMemory?: int64  @go(MinPhysMemory)
	machineOs?:     string @go(MachineOs)
	machineArch?:   string @go(MachineArch)
}

// Reservation implements all methods required to be
// a valid DRMAA2 compatible reservation, created by a
// ReservationSession.
#Reservation: _

// ReservationSession provides all methods required for a DRMAA2
// compatible reservation session.
#ReservationSession: _
