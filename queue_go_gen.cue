// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/dgruber/drmaa2interface

package drmaa2cue

// Queue implements all required elements of a queue.
#Queue: {
	Extensible: #Extensible
	extension?: #Extension @go(Extension)
	name?:      string     @go(Name)
}