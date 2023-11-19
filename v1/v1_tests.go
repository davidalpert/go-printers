package printers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("go-printers", func() {
	Describe("#NewPrinterOptions", func() {
		var o *PrinterOptions

		BeforeEach(func() {
			o = NewPrinterOptions()
		})

		It("should set default output to text", func() {
			Ω(o.DefaultOutputFormat).ShouldNot(BeNil())
			Ω(*o.DefaultOutputFormat).Should(Equal("text"))
		})

		It("should set output to nil", func() {
			Ω(o.OutputFormat).Should(BeNil())
		})
	})
})
