// Code generated by "stringer -type=CPU"; DO NOT EDIT.

package drmaa2interface

import "fmt"

const _CPU_name = "OtherCPUAlphaARMARM64CellPARISCPARISC64x86x64IA64MIPSMIPS64PowerPCPowerPC64SPARCSPARC64"

var _CPU_index = [...]uint8{0, 8, 13, 16, 21, 25, 31, 39, 42, 45, 49, 53, 59, 66, 75, 80, 87}

func (i CPU) String() string {
	if i < 0 || i >= CPU(len(_CPU_index)-1) {
		return fmt.Sprintf("CPU(%d)", i)
	}
	return _CPU_name[_CPU_index[i]:_CPU_index[i+1]]
}

func CPUFromString(cpu string) (CPU, bool) {
	switch cpu {
	case "OtherCPU":
		return OtherCPU, true
	case "Alpha":
		return Alpha, true
	case "ARM":
		return ARM, true
	case "ARM64":
		return ARM64, true
	case "Cell":
		return Cell, true
	case "PARISC":
		return PARISC, true
	case "PARISC64":
		return PARISC64, true
	case "x86":
		return X86, true
	case "x64":
		return X64, true
	case "IA64":
		return IA64, true
	case "MIPS":
		return MIPS, true
	case "MIPS64":
		return MIPS64, true
	case "PowerPC":
		return PowerPC, true
	case "PowerPC64":
		return PowerPC64, true
	case "SPARC":
		return SPARC, true
	case "SPARC64":
		return SPARC64, true
	default:
		return OtherCPU, false
	}
}