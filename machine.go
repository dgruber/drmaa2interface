package drmaa2interface

// Machine represents a compute instance implementing the
// extension interface.
type Machine struct {
	Extensible
	Extension      `json:"extension,omitempty"`
	Name           string  `json:"name,omitempty"`
	Available      bool    `json:"available,omitempty"`
	Sockets        int64   `json:"sockets,omitempty"`
	CoresPerSocket int64   `json:"coresPerSocket,omitempty"`
	ThreadsPerCore int64   `json:"threadsPerCore,omitempty"`
	Load           float64 `json:"load,omitempty"`
	PhysicalMemory int64   `json:"physicalMemory,omitempty"`
	VirtualMemory  int64   `json:"virtualMemory,omitempty"`
	Architecture   CPU     `json:"architecture,omitempty"`
	OSVersion      Version `json:"osVersion,omitempty"`
	OS             OS      `json:"os,omitempty"`
}

// OS is the operating system type.
type OS int

//go:generate stringer -type=OS
const (
	OtherOS OS = iota
	AIX
	BSD
	Linux
	HPUX
	IRIX
	MacOS
	SunOS
	TRU64
	UnixWare
	Win
	WinNT
)
