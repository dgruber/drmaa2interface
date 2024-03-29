// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/dgruber/drmaa2interface

// Package drmaa2interface implements the DRMAA2 Go interface.
// Actual Go DRMAA2 compatible implementations should use that
// interface to guarantee compatibility.
package drmaa2cue

import "time"

#ZeroTime: time.#Duration & 0

#InfiniteTime: time.#Duration & 9223372036854775807

#UnsetTime: -3

#UnsetNum: -1

#UnsetEnum: -1

#PlaceholderHomeDir: "$DRMAA2_HOME_DIR$"

#PlaceholderWorkingDir: "$DRMAA2_WORKING_DIR$"

#PlaceholderIndex: "$DRMAA2_INDEX$"
