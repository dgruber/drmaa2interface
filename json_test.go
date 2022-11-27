package drmaa2interface_test

import (
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/dgruber/drmaa2interface"
)

var _ = Describe("Json", func() {

	Context("JobTemplate", func() {

		It("should marshal and unmarshal JobTemplate", func() {
			jt := drmaa2interface.JobTemplate{
				RemoteCommand: "echo",
				Args:          []string{"hello", "world"},
				Extension: drmaa2interface.Extension{
					ExtensionList: map[string]string{
						"foo": "bar",
					},
				},
			}

			json, err := jt.MarshalJSON()
			Expect(err).To(BeNil())
			Expect(json).ToNot(BeNil())

			log.Printf("json: %s", string(json))

			var jt2 drmaa2interface.JobTemplate
			err = jt2.UnmarshalJSON(json)
			Expect(err).To(BeNil())
			Expect(jt2.ExtensionList["foo"]).To(Equal("bar"))

			Expect(jt).To(Equal(jt2))
		})

		It("should parse unset times as zero times", func() {
			e := `{"extension":{"extensionList":{"foo":"bar"}},"remoteCommand":"echo","args":["hello","world"]}`
			var jt drmaa2interface.JobTemplate
			err := jt.UnmarshalJSON([]byte(e))
			Expect(err).To(BeNil())
			Expect(jt.ExtensionList["foo"]).To(Equal("bar"))
			Expect(jt.StartTime.IsZero()).To(BeTrue())
		})

	})

})
