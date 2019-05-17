<!--
  Attractive html formatting for rendering in github. sorry text editor
  readers! Besides the header and section links, everything should be clean and
  readable.
-->
<h1 align="center">vice</h1>
<p align="center"><i>Standard bad behaviours for <a href="https://golang.org">Go</a> error values</i></p>

<div align="center">
  <a href="https://godoc.org/github.com/jbowes/vice"><img src="https://godoc.org/github.com/jbowes/vice?status.svg" alt="GoDoc"></a>
  <img alt="Alpha Quality" src="https://img.shields.io/badge/status-ALPHA-orange.svg" >
  <a href="https://travis-ci.com/jbowes/vice"><img alt="Build Status" src="https://travis-ci.com/jbowes/vice.svg?branch=master"></a>
  <a href="https://github.com/jbowes/vice/releases/latest"><img alt="GitHub tag" src="https://img.shields.io/github/tag/jbowes/vice.svg"></a>
  <a href="./LICENSE"><img alt="BSD license" src="https://img.shields.io/badge/license-BSD-blue.svg"></a>
  <a href="https://codecov.io/gh/jbowes/vice"><img alt="codecov" src="https://img.shields.io/codecov/c/github/jbowes/vice.svg"></a>
  <a href="https://goreportcard.com/report/github.com/jbowes/vice"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/jbowes/vice"></a>
</div><br /><br />


## Introduction
Introduction | [Examples] | [Contributing] <br /><br />

🚧 ___Disclaimer___: _`vice` is alpha quality software. The API may change
without warning between revisions._ 🚧

`vice` provides bad behaviours for the new [Go] 2/1.13+ error values.


## Examples
[Introduction] | Examples | [Contributing] <br /><br />

For complete examples and usage, see the [GoDoc documentation](https://godoc.org/github.com/jbowes/vice).

### Building APIs on `vice`

`vice/skip` implements the `vice` API with an additional `skip` argument,
allowing creation of APIs on top of `vice` that will no report themselves in
error caller frames.


## Contributing
[Introduction] | [Examples] | [Usage] | Contributing <br /><br />

I would love your help!

`vice` is still a work in progress. You can help by:

- Opening a pull request to resolve an [open issue][issues].
- Adding a feature or enhancement of your own! If it might be big, please
  [open an issue][enhancement] first so we can discuss it.
- Improving this `README` or adding other documentation to `vice`.
- Letting [me] know if you're using `vice`.


<!-- Section link definitions -->
[introduction]: #introduction
[examples]: #examples
[usage]: #usage
[contributing]: #contributing

<!-- Other links -->
[go]: https://golang.org

[issues]: ./issues
[bug]: ./issues/new?labels=bug
[enhancement]: ./issues/new?labels=enhancement

[me]: https://twitter.com/jrbowes
