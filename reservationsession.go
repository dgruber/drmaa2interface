package drmaa2interface

import (
	"time"
)

// ReservationInfo contains all details of the current state of
// a resource reservation.
type ReservationInfo struct {
	Extensible
	Extension            `json:"extension,omitempty"`
	ReservationID        string    `json:"reservationID,omitempty"`
	ReservationName      string    `json:"reservationName,omitempty"`
	ReservationStartTime time.Time `json:"reservationStartTime,omitempty"`
	ReservationEndTime   time.Time `json:"reservationEndTime,omitempty"`
	ACL                  []string  `json:"acl,omitempty"`
	ReservedSlots        int64     `json:"reservedSlots,omitempty"`
	ReservedMachines     []string  `json:"reservedMachines,omitempty"`
}

// ReservationTemplate contains ressource requests for a
// resource reservation.
type ReservationTemplate struct {
	Extensible
	Extension         `json:"extension,omitempty"`
	Name              string        `json:"name,omitempty"`
	StartTime         time.Time     `json:"startTime,omitempty"`
	EndTime           time.Time     `json:"endTime,omitempty"`
	Duration          time.Duration `json:"duration,omitempty"`
	MinSlots          int64         `json:"minSlots,omitempty"`
	MaxSlots          int64         `json:"maxSlots,omitempty"`
	JobCategory       string        `json:"jobCategory,omitempty"`
	UsersACL          []string      `json:"userACL,omitempty"`
	CandidateMachines []string      `json:"candidateMachines,omitempty"`
	MinPhysMemory     int64         `json:"minPhysMemory,omitempty"`
	MachineOs         string        `json:"machineOs,omitempty"`
	MachineArch       string        `json:"machineArch,omitempty"`
}

// Reservation implements all methods required to be
// a valid DRMAA2 compatible reservation, created by a
// ReservationSession.
type Reservation interface {
	GetID() (string, error)
	GetSessionName() (string, error)
	GetTemplate() (ReservationTemplate, error)
	GetInfo() (ReservationInfo, error)
	Terminate() error
}

// ReservationSession provides all methods required for a DRMAA2
// compatible reservation session.
type ReservationSession interface {
	Close() error
	GetContact() (string, error)
	GetSessionName() (string, error)
	GetReservation(string) (Reservation, error)
	RequestReservation(ReservationTemplate) (Reservation, error)
	GetReservations() ([]Reservation, error)
}
