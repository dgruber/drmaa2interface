// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/dgruber/drmaa2interface

package drmaa2cue

// Version is a DRMAA2 version type which holds the major version
// and minor version.
#Version: {
	major?: string @go(Major)
	minor?: string @go(Minor)
}