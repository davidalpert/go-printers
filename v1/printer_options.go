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
	if !StringInSlice(o.SupportedFormats(), *o.OutputFormat) {
		return fmt.Errorf("invalid output format: %s\nvalid format values are: %v", *o.OutputFormat, strings.Join(o.SupportedFormatCategories(), "|"))
	}
	return nil
}

// FormatCategory returns the dereferenced format category
func (o *PrinterOptions) FormatCategory() string {
	ExitIfErr(o.Validate())

	if o.TablePopulateFN != nil {
		return supportedListPrinterFormatMap[*o.OutputFormat]
	}
	return supportedObjectPrinterFormatMap[*o.OutputFormat]
}

// AddPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) AddPrinterFlags(c *pflag.FlagSet) {
	if o.OutputFormat != nil {
		if o.TablePopulateFN != nil {
			o.addListPrinterFlags(c)
		} else {
			o.addObjectPrinterFlags(c)
		}
	}
}

// AddObjectPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) addObjectPrinterFlags(c *pflag.FlagSet) {
	if o.OutputFormat != nil {
		c.StringVarP(o.OutputFormat, "output", "o", *o.DefaultOutputFormat, fmt.Sprintf("output format: one of %s.", strings.Join(supportedObjectPrinterCategories, "|")))
	}
}

// AddListPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) addListPrinterFlags(c *pflag.FlagSet) {
	if o.OutputFormat != nil {
		c.StringVarP(o.OutputFormat, "output", "o", *o.DefaultOutputFormat, fmt.Sprintf("output format: one of %s.", strings.Join(supportedListPrinterCategories, "|")))
	}
}
