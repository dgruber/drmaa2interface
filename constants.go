// Package drmaa2interface implements the DRMAA2 Go interface.
// Actual Go DRMAA2 compatible implementations should use that
// interface to guarantee compatibility.
package drmaa2interface

import (
	"time"
)

// ZeroTime is a special timeout value: Don't wait
const ZeroTime = time.Duration(0)

// InfiniteTime is a special timeout value: Wait probably infinitly
const InfiniteTime = time.Duration(1<<63 - 1)

// Now is always interpreted as the current time.
// const Now = -2

// UnsetTime is a special time value: Time or date not set
const UnsetTime = -3
const UnsetNum = -1
const UnsetEnum = -1

var UnsetList []interface{} = nil
var UnsetDict map[string]string = nil
var UnsetJobInfo *JobInfo = nil

const PlaceholderHomeDir = "$DRMAA2_HOME_DIR$"
const PlaceholderWorkingDir = "$DRMAA2_WORKING_DIR$"
const PlaceholderIndex = "$DRMAA2_INDEX$"
