package printers

import (
	"fmt"
	"github.com/spf13/pflag"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// PrinterOptions contains options for printing
type PrinterOptions struct {
	OutputFormat        *string
	DefaultOutputFormat *string
	TableCaption        *string
	TablePopulateFN     *func(*tablewriter.Table)
	ItemsSelector       *func() interface{}
	IOStreams
}

// NewPrinterOptions defines new printer options
func NewPrinterOptions() *PrinterOptions {
	return (&PrinterOptions{IOStreams: DefaultOSStreams()}).WithDefaultOutput("text")
}

// WithStreams sets the IOStreams used by the PrinterOptions
func (o *PrinterOptions) WithStreams(s IOStreams) *PrinterOptions {
	o.IOStreams = s
	return o
}

// WithDefaultOutput sets a default output format if one is not provided through a flag value
func (o *PrinterOptions) WithDefaultOutput(output string) *PrinterOptions {
	o.DefaultOutputFormat = &output
	return o
}

// WithDefaultTableWriter sets a default table writer
func (o *PrinterOptions) WithDefaultTableWriter() *PrinterOptions {
	return o.WithDefaultOutput("table").WithTableWriter("n/a", func(t *tablewriter.Table) {})
}

// WithTableWriter decorates a PrinterOptions with table writer configuration
func (o *PrinterOptions) WithTableWriter(caption string, populateTable func(*tablewriter.Table)) *PrinterOptions {
	o.TableCaption = &caption
	o.TablePopulateFN = &populateTable
	return o
}

func (o *PrinterOptions) WithItemsSelector(selectItems func() interface{}) *PrinterOptions {
	o.ItemsSelector = &selectItems
	return o
}

// ActiveOutputFormat provides the resolved output format (falling back to configured default, then text)
// this allows us to reduce the use of * which assumes that we have a value; this function always returns
// a string, resolving in one place the possibility of either configured o.OutputFormat or o.DefaultOutputFormat
// to be nil
func (o *PrinterOptions) ActiveOutputFormat() string {
	if o.OutputFormat != nil && *o.OutputFormat != "" {
		return *o.OutputFormat
	}
	if o.DefaultOutputFormat != nil && *o.DefaultOutputFormat != "" {
		return *o.DefaultOutputFormat
	}
	return "text"
}

// SupportedFormats returns the list of supported formats
func (o *PrinterOptions) SupportedFormats() []string {
	if o.TablePopulateFN != nil {
		return supportedListPrinterKeys
	}
	return supportedObjectPrinterKeys
}

// SupportedFormatCategories returns the list of supported formats
func (o *PrinterOptions) SupportedFormatCategories() []string {
	if o.TablePopulateFN != nil {
		return supportedListPrinterCategories
	}
	return supportedObjectPrinterCategories
}

// Validate asserts that the printer options are valid
func (o *PrinterOptions) Validate() error {
	if !StringInSlice(o.SupportedFormats(), o.ActiveOutputFormat()) {
		return fmt.Errorf("invalid output format: %s\nvalid format values are: %v", o.ActiveOutputFormat(), strings.Join(o.SupportedFormatCategories(), "|"))
	}
	return nil
}

// FormatCategory returns the dereferenced format category
func (o *PrinterOptions) FormatCategory() string {
	ExitIfErr(o.Validate())

	if o.TablePopulateFN != nil {
		return supportedListPrinterFormatMap[o.ActiveOutputFormat()]
	}
	return supportedObjectPrinterFormatMap[o.ActiveOutputFormat()]
}

// AddPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) AddPrinterFlags(c *pflag.FlagSet) {
	if o.TablePopulateFN != nil {
		o.addListPrinterFlags(c)
	} else {
		o.addObjectPrinterFlags(c)
	}
}

// AddObjectPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) addObjectPrinterFlags(c *pflag.FlagSet) {
	if o.OutputFormat != nil {
		c.StringVarP(o.OutputFormat, "output", "o", o.ActiveOutputFormat(), fmt.Sprintf("output format: one of %s.", strings.Join(supportedObjectPrinterCategories, "|")))
	}
}

// AddListPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) addListPrinterFlags(c *pflag.FlagSet) {
	if o.OutputFormat != nil {
		c.StringVarP(o.OutputFormat, "output", "o", o.ActiveOutputFormat(), fmt.Sprintf("output format: one of %s.", strings.Join(supportedListPrinterCategories, "|")))
	}
}
