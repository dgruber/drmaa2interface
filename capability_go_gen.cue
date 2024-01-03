// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/dgruber/drmaa2interface

package drmaa2cue

// Capability is a type which represents the availability of optional
// functionality of the DRMAA2 implementation. Optional functionality
// is defined by the DRMAA2 standard but not mandatory to implement.
#Capability: int // #enumCapability

#enumCapability:
	#AdvanceReservation |
	#ReserveSlots |
	#Callback |
	#BulkJobsMaxParallel |
	#JtEmail |
	#JtStaging |
	#JtDeadline |
	#JtMaxSlots |
	#JtAccountingID |
	#RtStartNow |
	#RtDuration |
	#RtMachineOS |
	#RtMachineArch

#values_Capability: {
	AdvanceReservation:  #AdvanceReservation
	ReserveSlots:        #ReserveSlots
	Callback:            #Callback
	BulkJobsMaxParallel: #BulkJobsMaxParallel
	JtEmail:             #JtEmail
	JtStaging:           #JtStaging
	JtDeadline:          #JtDeadline
	JtMaxSlots:          #JtMaxSlots
	JtAccountingID:      #JtAccountingID
	RtStartNow:          #RtStartNow
	RtDuration:          #RtDuration
	RtMachineOS:         #RtMachineOS
	RtMachineArch:       #RtMachineArch
}

#AdvanceReservation:  #Capability & 0
#ReserveSlots:        #Capability & 1
#Callback:            #Capability & 2
#BulkJobsMaxParallel: #Capability & 3
#JtEmail:             #Capability & 4
#JtStaging:           #Capability & 5
#JtDeadline:          #Capability & 6
#JtMaxSlots:          #Capability & 7
#JtAccountingID:      #Capability & 8
#RtStartNow:          #Capability & 9
#RtDuration:          #Capability & 10
#RtMachineOS:         #Capability & 11
#RtMachineArch:       #Capability & 12