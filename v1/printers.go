package printers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"

	"github.com/go-xmlfmt/xmlfmt"
	"github.com/gocarina/gocsv"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

// WriteOutput writes output in the configured format with default printer options
// retained for backwards compat with verb-noun commands which assume STDOUT
func WriteOutput(v interface{}, caption string, populateTable func(*tablewriter.Table)) error {
	return NewPrinterOptions().WithTableWriter(caption, populateTable).WriteOutput(v)
}

// DescribeObject writes a single object in the default format with default printer options
// retained for backwards compat with verb-noun commands which assume STDOUT
func DescribeObject(v interface{}) error {
	return NewPrinterOptions().WriteOutput(v)
}

// WriteErr writes output to s.Err in the configured format
func (o *PrinterOptions) WriteErr(v interface{}) error {
	return o.FWriteOutput(o.ErrOut, v)
}

// WriteOutput writes output to s.Out in the configured format
func (o *PrinterOptions) WriteOutput(v interface{}) error {
	return o.FWriteOutput(o.Out, v)
}

// FWriteOutput writes output to given ioWriter in the configured format
func (o *PrinterOptions) FWriteOutput(s io.Writer, v interface{}) error {
	ExitIfErr(o.Validate())

	output, _, err := o.FormatOutput(v)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(s, output)
	return err
}

// FormatOutput writes a single object in the configured format
func FormatOutput(v interface{}, outputFormat string) (string, string, error) {
	return NewPrinterOptions().FormatOutput(v)
}

// FormatOutput writes a single object in the configured format
func (o *PrinterOptions) FormatOutput(v interface{}) (string, string, error) {
	outputFormat := o.ActiveOutputFormat()
	formatCategory := o.FormatCategory()
	output := ""

	// tables are special and only work if the PrinterOptions have defined a TablePopulateFN function
	if formatCategory == "table" {
		if o.TablePopulateFN == nil {
			ExitIfErr(fmt.Errorf("output format is %s but TablePopulateFN is not defined", outputFormat))
		}

		outputStringBuilder := &strings.Builder{}
		table := tablewriter.NewWriter(outputStringBuilder)
		if *o.TableCaption != "" {
			table.SetCaption(true, *o.TableCaption)
		}
		(*o.TablePopulateFN)(table)

		table.Render()

		return outputStringBuilder.String(), formatCategory, nil
	} else if formatCategory != "" {
		// the rest of the format categories have category-specific marshalers
		return o.marshalObjectToString(v, formatCategory)
	}

	return output, formatCategory, fmt.Errorf("could not map output=%s to a format category", outputFormat)
}

var supportedObjectPrinterFormatMap = map[string]string{
	"j":    "json",
	"js":   "json",
	"json": "json",
	"t":    "text",
	"txt":  "text",
	"text": "text",
	"y":    "yaml",
	"yml":  "yaml",
	"yaml": "yaml",
}

var supportedListPrinterFormatMap = map[string]string{
	"c":     "csv",
	"csv":   "csv",
	"table": "table",
	"x":     "xml",
	"xml":   "xml",
}

var supportedObjectPrinterKeys = []string{}
var supportedObjectPrinterCategories = []string{}

var supportedListPrinterKeys = []string{}
var supportedListPrinterCategories = []string{}

func (o *PrinterOptions) marshalObjectToString(v interface{}, formatCategory string) (string, string, error) {
	output := ""
	if formatCategory == "text" {
		switch vv := v.(type) {
		case fmt.Stringer:
			output = vv.String()
			break
		case fmt.GoStringer:
			output = vv.GoString()
			break
		default:
			output = fmt.Sprintf("%v", v)
			break
		}
	} else if formatCategory == "yaml" {
		oB, _ := yaml.Marshal(v)
		output = string(oB)
	} else if formatCategory == "json" {
		oB, _ := json.MarshalIndent(v, "", "  ")
		output = string(oB) + "\n"
	} else if formatCategory == "xml" {
		if o.ItemsSelector != nil {
			v = (*o.ItemsSelector)()
		}
		oB, _ := xml.Marshal(v)
		if reflect.ValueOf(v).Kind() == reflect.Struct {
			// the interface is a struct, render it raw
			output = xmlfmt.FormatXML(string(oB), "", "  ")
		} else {
			// the interface is probably a slice; we need to wrap it in a root element
			output = fmt.Sprintf("<Result>%s\n</Result>\n", xmlfmt.FormatXML(string(oB), "  ", "  "))
		}
	} else if formatCategory == "csv" {
		if o.ItemsSelector != nil {
			v = (*o.ItemsSelector)()
		}
		oB, _ := gocsv.MarshalString(v)
		output = string(oB)
	} else {
		return output, formatCategory, fmt.Errorf("do not support format category %s", formatCategory)
	}

	return output, formatCategory, nil
}

// parseOutput reads a single object in the configured format
func parseOutput(in []byte, outputFormat string, v interface{}) error {
	return NewPrinterOptions().WithDefaultOutput(outputFormat).ParseOutput(in, v)
}

// ParseOutput reads a single object in the configured format
func (o *PrinterOptions) ParseOutput(in []byte, v interface{}) error {
	formatCategory := o.FormatCategory()

	if formatCategory == "yaml" {
		return yaml.Unmarshal(in, v)
	} else if formatCategory == "json" {
		return json.Unmarshal(in, v)
	} else if formatCategory == "xml" {
		return xml.Unmarshal(in, v)
	} else if formatCategory == "csv" {
		return gocsv.UnmarshalBytes(in, v)
	}

	return fmt.Errorf("do not know how to parse output=%s or category %s", o.ActiveOutputFormat(), formatCategory)
}

func containsString(list []string, q string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == q {
			return true
		}
	}
	return false
}

func init() {
	// object printer formats are also list printer formats
	for k, v := range supportedObjectPrinterFormatMap {
		supportedListPrinterFormatMap[k] = v
	}

	// extract and sort supportedObjectPrinterKeys and Categories
	for k, c := range supportedObjectPrinterFormatMap {
		supportedObjectPrinterKeys = append(supportedObjectPrinterKeys, k)

		if !containsString(supportedObjectPrinterCategories, c) {
			supportedObjectPrinterCategories = append(supportedObjectPrinterCategories, c)
		}
	}

	sort.Slice(supportedObjectPrinterKeys, func(i, j int) bool {
		return supportedObjectPrinterKeys[i] < supportedObjectPrinterKeys[j]
	})

	sort.Slice(supportedObjectPrinterCategories, func(i, j int) bool {
		return supportedObjectPrinterCategories[i] < supportedObjectPrinterCategories[j]
	})

	// extract and sort supportedListPrinterKeys and Categories
	for k, c := range supportedListPrinterFormatMap {
		supportedListPrinterKeys = append(supportedListPrinterKeys, k)

		if !containsString(supportedListPrinterCategories, c) {
			supportedListPrinterCategories = append(supportedListPrinterCategories, c)
		}
	}

	sort.Slice(supportedListPrinterKeys, func(i, j int) bool {
		return supportedListPrinterKeys[i] < supportedListPrinterKeys[j]
	})

	sort.Slice(supportedListPrinterCategories, func(i, j int) bool {
		return supportedListPrinterCategories[i] < supportedListPrinterCategories[j]
	})
}
