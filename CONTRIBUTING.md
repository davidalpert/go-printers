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
  <a href="./README.md">README</a>
  ·
  <a href="./CHANGELOG.md">CHANGELOG</a>
  .
  <a href="./CONTRIBUTING.md"><strong>CONTRIBUTING</strong></a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-printers">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-printers/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-printers/issues">Request Feature</a>
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Table of contents</h2></summary>

- [Review existing issues](#review-existing-issues)
- [Setup for local development](#setup-for-local-development)
  - [Install prerequisites](#install-prerequisites)
  - [Get the code](#get-the-code)
  - [Visit the doctor](#visit-the-doctor)
  - [Run locally](#run-locally)
- [Useful Make targets](#useful-make-targets)
- [Development workflow](#development-workflow)
  - [Branch names](#branch-names)
  - [Commit message guidelines](#commit-message-guidelines)

</details>

Contributions make the open source community an great place to learn, inspire, and create.

Please review this contribution guide to streamline your experience.

## Review existing issues

Please review existing [issues](https://github.com/davidalpert/go-printers/issues) before reporting bug reports or requesting new features.

A quick discussion to coordinate a proposed change before you start can save hours of rework. 

## Setup for local development

### Install prerequisites

* [Make](https://www.gnu.org/software/make/manual/html_node/index.html#Top)  - often comes bundled with C compiler tools
* [Golang 1.18](https://golang.org/doc/manage-install)
  * with a working go installation:
    ```
    go install golang.org/dl/go1.18@latest
    go1.18 download
    ```
  * open a terminal with `go1.18` as the linked `go` binary

* [ruby 3.0.2](https://www.ruby-lang.org/en/documentation/installation/) [optional]

  * this project offers ruby and guard to provide real-time validation of grammar and spelling in key markdown files
  * use a ruby version manager like [rbenv](https://github.com/rbenv) or [asdf](https://asdf-vm.com/); or
  * install directly from [ruby-lang.org](https://www.ruby-lang.org/en/documentation/installation/)

### Get the code

1. [Fork the repository on Github](https://github.com/davidalpert/go-printers/fork)

1. Clone your fork
   ```sh
   git clone https://github.com/your-github-name/go-printers.git
   ```

### Visit the doctor

This repository includes a `doctor.sh` script which validates development dependencies.

1. Verify dependencies
    ```sh
    ./.tools/doctor.sh
    ```

This script attempts to fix basic issues, for example by running `go get` or `bundle install`.

If `doctor.sh` reports an issue that it can't resolve you may need to help it by taking action.

Please log any issues with the doctor script by [reporting a bug](https://github.com/davidalpert/go-printers/issues).

### Run locally

1. Build and run the tests
    ```sh
    make test
    ```

## Useful Make targets

This repository includes a `Makefile` for help running common tasks.

Run `make` with no arguments to list the available targets:
```
$ make

 go-printers 0.0.0 - available targets:

changelog                      Generate/update CHANGELOG.md
cit                            clean build and test-unit
clean                          clean build output
doctor                         run doctor.sh to sort out development dependencies
gen                            invoke go generate
guard                          run ruby-guard
preview-release-notes          preview release notes (generates RELEASE_NOTES.md)
preview-release                preview release (using goreleaser --snapshot)
test-unit                      run unit tests
test                           run unit tests
vale                           run linting rules against markdown files
version                        show current version
----------                     ------------------
release-major                  release major version
release-minor                  release minor version
release-patch                  release patch version
```

## Development workflow

This project follows a standard open source fork/pull-request workflow:

1. First, [fork the repository on Github](https://github.com/davidalpert/go-printers/fork)


1. Create your Feature Branch
   ```
   git checkout -b 123-amazing-feature
   ```
1. Commit your Changes
   ```
   git commit -m 'Add some AmazingFeature'
   ```
1. Make sure the code builds and all tests pass
   ```
   make cit
   ```
3. Push to the Branch
   ```
   git push origin 123-amazing-feature
   ```
4. Open a Pull Request

    https://github.com/davidalpert/go-printers/compare/123-amazing-feature

### Branch names

When working on a pull request to address or resolve a Github issue, prefix the branch name with the Github issue number.

In the preceding example, after picking up an issue with an id of 123, create a branch which starts with `GH-123` or just `123-` and a hyphenated description:

```
git checkout -b 123-amazing-feature
```

### Commit message guidelines

This project uses [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/#summary) to generate [CHANGELOG](CHANGELOG.md).

Format of a conventional commit:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

List of supported commit type tags include:
```yaml
  - "build"    # Changes that affect the build system or external dependencies
  - "ci"       # Changes to our CI configuration files and scripts 
  - "docs"     # Documentation only changes
  - "feat"     # A new feature
  - "fix"      # A bug fix
  - "perf"     # A code change that improves performance
  - "refactor" # A code change that neither fixes a bug nor adds a feature
  - "test"     # Adding missing tests or correcting existing tests
```

Prefix your commits with one of these type tags to automatically include the commit description in the [CHANGELOG](CHANGELOG.md) for the next release.

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