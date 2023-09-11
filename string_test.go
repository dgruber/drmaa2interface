package drmaa2interface_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/dgruber/drmaa2interface"
)

var _ = Describe("String", func() {

	Context("String to constants", func() {

		It("should convert string to Capability", func() {
			// loop through all capabilities
			for i := drmaa2interface.AdvanceReservation; i <= drmaa2interface.RtMachineArch; i++ {
				cap, ok := drmaa2interface.CapabilityFromString(i.String())
				fmt.Printf("%s\n", i.String())
				Expect(ok).To(BeTrue())
				Expect(cap).To(Equal(i))
			}
		})

		It("should convert string to JobState", func() {
			// loop through all states
			for i := drmaa2interface.Unset; i <= drmaa2interface.Failed; i++ {
				state, ok := drmaa2interface.JobStateFromString(i.String())
				Expect(ok).To(BeTrue())
				Expect(state).To(Equal(i))
			}
		})

		It("should convert string to ErrorID", func() {
			// loop through all error ids
			for i := drmaa2interface.Success; i <= drmaa2interface.LastError; i++ {
				id, ok := drmaa2interface.ErrorIDFromString(i.String())
				Expect(ok).To(BeTrue())
				Expect(id).To(Equal(i))
			}
		})

		It("should convert string to OS", func() {
			// loop through all operating systems
			for i := drmaa2interface.OtherOS; i <= drmaa2interface.WinNT; i++ {
				os, ok := drmaa2interface.OSFromString(i.String())
				Expect(ok).To(BeTrue())
				Expect(os).To(Equal(i))
			}
		})

		It("should convert string to CPU", func() {
			// loop through all CPUs
			for i := drmaa2interface.OtherCPU; i <= drmaa2interface.SPARC64; i++ {
				cpu, ok := drmaa2interface.CPUFromString(i.String())
				Expect(ok).To(BeTrue())
				Expect(cpu).To(Equal(i))
			}
		})
	})
})
