// Code generated by "stringer -type=ErrorID"; DO NOT EDIT.

package drmaa2interface

import "fmt"

const _ErrorID_name = "SuccessDeniedByDrmsDrmCommunicationTryLaterSessionManagementTimeoutInternalInvalidArgumentInvalidSessionInvalidStateOutOfResourceUnsupportedAttributeUnsupportedOperationImplementationSpecificLastError"

var _ErrorID_index = [...]uint8{0, 7, 19, 35, 43, 60, 67, 75, 90, 104, 116, 129, 149, 169, 191, 200}

func (i ErrorID) String() string {
	if i < 0 || i >= ErrorID(len(_ErrorID_index)-1) {
		return fmt.Sprintf("ErrorID(%d)", i)
	}
	return _ErrorID_name[_ErrorID_index[i]:_ErrorID_index[i+1]]
}

func ErrorIDFromString(id string) (ErrorID, bool) {
	switch id {
	case "Success":
		return Success, true
	case "DeniedByDrms":
		return DeniedByDrms, true
	case "DrmCommunication":
		return DrmCommunication, true
	case "TryLater":
		return TryLater, true
	case "SessionManagement":
		return SessionManagement, true
	case "Timeout":
		return Timeout, true
	case "Internal":
		return Internal, true
	case "InvalidArgument":
		return InvalidArgument, true
	case "InvalidSession":
		return InvalidSession, true
	case "InvalidState":
		return InvalidState, true
	case "OutOfResource":
		return OutOfResource, true
	case "UnsupportedAttribute":
		return UnsupportedAttribute, true
	case "UnsupportedOperation":
		return UnsupportedOperation, true
	case "ImplementationSpecific":
		return ImplementationSpecific, true
	case "LastError":
		return LastError, true
	default:
		return Success, false
	}
}