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
  <a href="./CHANGELOG.md"><strong>CHANGELOG<strong></a>
  .
  <a href="./CONTRIBUTING.md">CONTRIBUTING</a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-printers">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-printers/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-printers/issues">Request Feature</a>
</p>

## Changelog


<a name="v0.3.2"></a>
## [v0.3.2] - 2023-11-26
### Bug Fixes
- outsmarting the golang fmt package doesn't work
- marshalObjectToString ignored fmt.Stringer
- AddPrinterOptions broke in 0.3.1

### Build
- a .tools-version file sets golang locally for asdf

### Features
- use GoStringer when available


<a name="v0.3.1"></a>
## [v0.3.1] - 2023-11-24
### Bug Fixes
- default output format not set properly
- oldest one in the book

### Build
- don't bother with gen for a package
- add ci/cd actions
- replace make with task

### Chore
- downgrade module version to go 1.17

### Features
- add ActiveOutputFormat() method

### Test Coverage
- demonstrate using TestIOStreams
- reproduce issue [#1](https://github.com/davidalpert/go-git-mob/issues/1)
- add ginko and gomega as a testing framework

### Pull Requests
- Merge pull request [#2](https://github.com/davidalpert/go-git-mob/issues/2) from davidalpert/[GH-1](https://github.com/davidalpert/go-git-mob/issues/1)-default-output-format-not-properly-set
- Merge pull request [#6](https://github.com/davidalpert/go-git-mob/issues/6) from davidalpert/[GH-3](https://github.com/davidalpert/go-git-mob/issues/3)-add-cicd-actions
- Merge pull request [#5](https://github.com/davidalpert/go-git-mob/issues/5) from davidalpert/[GH-4](https://github.com/davidalpert/go-git-mob/issues/4)-replace-makefile-with-taskfile


<a name="v0.3.0"></a>
## [v0.3.0] - 2022-07-23
### Bug Fixes
- remove duplicate WriteOutput method

### Code Refactoring
- annonmize embeded streams


<a name="v0.2.0"></a>
## [v0.2.0] - 2022-07-22
### Code Refactoring
- embed streams inside printer options


<a name="v0.1.0"></a>
## v0.1.0 - 2022-07-17
### Features
- expose go-printers as a printers package


[Unreleased]: https://github.com/davidalpert/go-git-mob/compare/v0.3.2...HEAD
[v0.3.2]: https://github.com/davidalpert/go-git-mob/compare/v0.3.1...v0.3.2
[v0.3.1]: https://github.com/davidalpert/go-git-mob/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/davidalpert/go-git-mob/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/davidalpert/go-git-mob/compare/v0.1.0...v0.2.0
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