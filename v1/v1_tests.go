package printers

import (
	"bufio"
	"bytes"
	"fmt"
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

	Describe("#NewTestIOStreams", func() {
		var o *PrinterOptions
		var ioStreams IOStreams
		var inBuf *bytes.Buffer
		var outBuf *bytes.Buffer
		var errBuf *bytes.Buffer

		BeforeEach(func() {
			// create TestIOStreams, capturing the buffers bound to the given ioStreams
			ioStreams, inBuf, outBuf, errBuf = NewTestIOStreams()

			// configure your printer options to use the given ioStreams
			o = NewPrinterOptions().WithStreams(ioStreams)
		})

		It("should capture stdout as text", func() {
			// write via custom helpers
			err := o.WriteOutput("hello")
			Ω(err).Should(BeNil())

			// write via IOStreams.Out
			_, err = fmt.Fprintf(o.Out, " world")
			Ω(err).Should(BeNil())

			// read from the test buffer
			Ω(outBuf.String()).Should(Equal("hello world"))
		})

		It("should capture stderr as text", func() {
			// write via custom helpers
			err := o.WriteErr("hello")
			Ω(err).Should(BeNil())

			// write via IOStreams.ErrOut
			_, err = fmt.Fprintf(o.ErrOut, " world")
			Ω(err).Should(BeNil())

			// read from the test buffer
			Ω(errBuf.String()).Should(Equal("hello world"))
		})

		It("should capture feeding input via inBuf as text", func() {
			// feed inBuf
			_, err := inBuf.WriteString("hello world\n")
			Ω(err).Should(BeNil())

			// read input from IOStreams
			reader := bufio.NewReader(o.In)
			line, err := reader.ReadString('\n')

			Ω(err).Should(BeNil())
			Ω(line).Should(Equal("hello world\n"))
		})
	})
})
