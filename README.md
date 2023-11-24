<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- vale Google.Acronyms = NO -->
[![License: MIT v3][license-shield]][license-url]
<!-- vale Google.Acronyms = YES -->

<!-- [![Issues][issues-shield]][issues-url] -->
<!-- [![Forks][forks-shield]][forks-url] -->
<!-- ![GitHub Contributors][contributors-shield] -->
<!-- ![GitHub Contributors Image][contributors-image-url] -->

<!-- PROJECT LOGO -->
<br />
<!-- vale Google.Headings = NO -->
<h1 align="center">go-printers</h1>
<!-- vale Google.Headings = YES -->

<p align="center">
  A Golang module built on top of the <a href="https://github.com/spf13/pflag">spf13/pflag</a> library to assist with abstracting output formatting.
  <br />
  <a href="./README.md"><strong>README</strong></a>
  ·
  <a href="./CHANGELOG.md">CHANGELOG</a>
  .
  <a href="./CONTRIBUTING.md">CONTRIBUTING</a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-printers">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-printers/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-printers/issues">Request Feature</a>
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Table of contents</h2></summary>

- [About the project](#about-the-project)
  - [Built with](#built-with)
- [Getting started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

</details>

<!-- ABOUT THE PROJECT -->
## About the project

`go-printers` provides a set of helpers extending the [spf13/pflag](https://github.com/spf13/pflag) library to abstract away print formatting, making it easy to separate output formatting and support printing as text, table, json, or yaml.

### Built with

* [Golang 1.18](https://golang.org/)
<!-- vale Google.Acronyms = NO -->
* [spf13/pflag](https://github.com/spf13/pflag) - "a drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags"
<!-- vale Google.Acronyms = YES -->

<!-- GETTING STARTED -->
## Getting started

To get a local copy up and running follow these simple steps.

### Installation

```
go get https://github.com/davidalpert/go-printers
```

### Usage

1. Import the package

    ```
    import (
    	"github.com/davidalpert/go-printers/v1"
    )
    ```

1. Define an Options struct and a factory function to set defaults

    ```
    type MyCmdOptions struct {
      *printers.PrinterOptions
      printers.IOStreams
    }

    func NewMyCmdOptions(ioStreams printers.IOStreams) *MyCmdOptions {
      return &MyCmdOptions{
        IOStreams:      ioStreams,
        PrinterOptions: printers.NewPrinterOptions(),
      }
    }
    ```

1. Create an instance of streams, probably in your startup code
    ```
    ioStreams := printers.DefaultOSStreams() 
    ```

    In a unit test you might create test streams which exposes each stream as a `*bytes.Buffer`:
    ```
    testStreams, testOutBuf, testErrBuf, testInBuf := printers.NewTestIOStreams()
    ```
   
    See the [#NewTestIOStreams](./v1/v1_suite_test.go#29-81) test scenario for an example of using these buffers in unit tests.

1. Configure flags inside a `cobra.Command` factory:

    ```
    func NewCmdMyCmd(ioStreams printers.IOStreams) *cobra.Command {
      o := NewMyCmdOptions(ioStreams)
      var cmd = &cobra.Command{
        Use:     "mycmd",
        Args:    cobra.NoArgs,
        RunE: func(cmd *cobra.Command, args []string) error {
         if err := o.Complete(cmd, args); err != nil {
          return err
         }
         if err := o.Validate(); err != nil {
          return err
         }
         return o.Run()
        },
       }
      
       o.PrinterOptions.AddPrinterFlags(cmd.Flags())
      
       return cmd
      }
    ```

1. Works well with the `Complete/Validate/Run` pattern:

    ```
    // Complete the options
    func (o *MyCmdOptions) Complete(cmd *cobra.Command, args []string) error {
      return nil
    }

    // Validate the options
    func (o *MyCmdOptions) Validate() error {
      return o.PrinterOptions.Validate()
    }
    ```

1. Send any `interface{}` or struct to the printer:

    ```
    type Car struct {
      Make    string `json:"make"`
      Model   string `json:"model"`
      Mileage int    `json:"mileage"`
    }

    // String implements Stringer
    func (c *Car) String() string {
      return fmt.Sprintf("%s %s [%d miles]", c.Make, c.Model, c.Mileage)
    }

    func (o *MyCmdOptions) Run() error {
      newCar := Car{
        make:    "Ford",
        model:   "Taurus",
        mileage: 200000,
      }

      return o.IOStreams.WriteOutput(newCar, o.PrinterOptions)
    }
    ```

This adds an `-o`/`--output` flag to your command:
```
Flags:
  -h, --help            help for mycmd
  -o, --output string   output format: one of json|text|yaml. (default "text")
```

which can then print the same output in different formats:

- with `-o text` and a type which implements `Stringer`:
    ```
    Ford Taurus [200000 miles]
    ```
- with `-o json`:
    ```
    {
      "make": "Ford",
      "model": "Taurus",
      "mileage": 200000
    }
    ```
- with `-o yaml`:
    ```
    ---
    make: Ford
    model: Taurus
    mileage: 200000
    ```

<!-- ROADMAP -->
## Roadmap

See [open issues](https://github.com/davidalpert/go-printers/issues) for a list of known issues and up-for-grabs tasks.

## Contributing

See the [CONTRIBUTING](CONTRIBUTING.md) guide for local development setup and contribution guidelines.

<!-- LICENSE -->
## License

Distributed under the GPU v3 License. See [LICENSE](LICENSE) for more information.

<!-- CONTACT -->
## Contact

David Alpert - [@davidalpert](https://twitter.com/davidalpert)

Project Link: [https://github.com/davidalpert/go-printers](https://github.com/davidalpert/go-printers)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/davidalpert/go-printers
[contributors-image-url]: https://contrib.rocks/image?repo=davidalpert/go-printers
[forks-shield]: https://img.shields.io/github/forks/davidalpert/go-printers
[forks-url]: https://github.com/davidalpert/go-printers/network/members
[issues-shield]: https://img.shields.io/github/issues/davidalpert/go-printers
[issues-url]: https://github.com/davidalpert/go-printers/issues
[license-shield]: https://img.shields.io/badge/License-MIT-yellow.svg
[license-url]: https://opensource.org/licenses/MIT

