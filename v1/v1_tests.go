package printers

import (
	"bufio"
	"bytes"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	flag "github.com/spf13/pflag"
)

type TestData struct {
	IntField    int
	StringField string
}

type TestDataWithStringer struct {
	TestData
}

func (d TestDataWithStringer) String() string {
	return fmt.Sprintf("stringed< %s is %d >", d.StringField, d.IntField)
}

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
			Ω(o.OutputFormat).Should(BeEmpty())
		})

		It("should set flags on a *pflag.FlagSet", func() {
			f := flag.NewFlagSet("test", flag.ContinueOnError)
			o.AddPrinterFlags(f)

			longFlag := f.Lookup("output")
			Ω(longFlag).ShouldNot(BeNil())
			Ω(longFlag.Name).Should(Equal("output"))
			Ω(longFlag.DefValue).Should(Equal("text"))

			shortFlag := f.ShorthandLookup("o")
			Ω(shortFlag).ShouldNot(BeNil())
			Ω(shortFlag.Name).Should(Equal("output"))
			Ω(longFlag.DefValue).Should(Equal("text"))
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

		Describe("with default printer options", func() {
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

		Describe("printing as txt", func() {
			BeforeEach(func() {
				o.WithDefaultOutput("txt")
			})

			It("should capture a struct without Stringer as %v", func() {
				testData := TestData{
					IntField:    42,
					StringField: "the answer to the question",
				}

				// write via custom helpers
				err := o.WriteOutput(testData)
				Ω(err).Should(BeNil())

				// read from the test buffer
				Ω(outBuf.String()).Should(Equal("{42 the answer to the question}"))
			})

			It("should capture a pointer to a struct without Stringer as %v", func() {
				ptrTestData := &TestData{
					IntField:    42,
					StringField: "the answer to the question",
				}

				// write via custom helpers
				err := o.WriteOutput(ptrTestData)
				Ω(err).Should(BeNil())

				// read from the test buffer
				Ω(outBuf.String()).Should(Equal("&printers.TestData{IntField:42, StringField:\"the answer to the question\"}"))
			})

			It("should capture a pointer to a struct without Stringer cast as interface{} as %v", func() {
				ptrTestData := &TestData{
					IntField:    42,
					StringField: "the answer to the question",
				}

				toEmptyInterface := func(v interface{}) interface{} {
					return v
				}

				v := toEmptyInterface(ptrTestData)

				// write via custom helpers
				err := o.WriteOutput(v)
				Ω(err).Should(BeNil())

				// read from the test buffer
				Ω(outBuf.String()).Should(Equal("&printers.TestData{IntField:42, StringField:\"the answer to the question\"}"))
			})

			It("should capture a struct with Stringer as %s", func() {
				testData := TestDataWithStringer{
					TestData: TestData{
						IntField:    42,
						StringField: "the answer to the question",
					},
				}

				// write via custom helpers
				err := o.WriteOutput(testData)
				Ω(err).Should(BeNil())

				// read from the test buffer
				Ω(outBuf.String()).Should(Equal("stringed< the answer to the question is 42 >"))
			})

			It("should capture a pointer to a struct with Stringer as %s", func() {
				ptrTestData := &TestDataWithStringer{
					TestData: TestData{
						IntField:    42,
						StringField: "the answer to the question",
					},
				}

				// write via custom helpers
				err := o.WriteOutput(ptrTestData)
				Ω(err).Should(BeNil())

				// read from the test buffer
				Ω(outBuf.String()).Should(Equal("stringed< the answer to the question is 42 >"))
			})
		})
	})
})
