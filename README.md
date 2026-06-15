<h1 align="center">slugigo</h1>

<p align="center">Fast, low-alloc Go slug generator with an elegant Fluent API.</p>

<p align="center">
    <a href="https://github.com/ivanbyone/slugigo/releases"><img src="https://img.shields.io/github/v/release/ivanbyone/slugigo?logo=github&sort=semver" alt="GH Release"></a>
    <a href="https://github.com/ivanbyone/slugigo/actions/workflows/slugigo.yml"><img alt="Slugigo CI Status" src="https://img.shields.io/github/actions/workflow/status/ivanbyone/slugigo/slugigo.yml?label=CI&logo=GitHub"></a>
    <a href="https://github.com/ivanbyone/slugigo/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-orange.svg" alt="Slugigo License"></a>
    <a href="https://pkg.go.dev/github.com/ivanbyone/slugigo"><img src="https://pkg.go.dev/badge/github.com/ivanbyone/slugigo.svg" alt="Go documentation"></a>
</p>

## Quick Start 

### Installation

```bash
go get -u github.com/ivanbyone/slugigo
```

### Usage example

```go
package main

import (
    "fmt"
    "github.com/ivanbyone/slugigo"
)

func main() {
    // Basic usage (ASCII, default compliance RFC-3986)
    slug := slugigo.Slug("Hello, Slugigo!").Build()
    fmt.Println(slug)
    // Output: hello-slugigo

    // Advanced custom configuration
    custom := slugigo.Slug("Hello, Fluent Slugigo").
        NoLowercase().
        MaxLength(12).
        Separator("_").
        SaveLeadingAndTrailingDash().
        Build()
    fmt.Println(custom)
    // Output: 
}
```

## Concept

Package `slugigo` provides functionality for generating URL-friendly slugs due to standard <b>[RFC-3986](https://datatracker.ietf.org/doc/html/rfc3986)</b>. So, <b>default</b> output complies with the following rules:

* Acceptable `ASCII` characters:
    * <b>only</b> lowercase letters: `a-z`
    * numbers: `0-9`
    * dash: `-`
    * underscore: `_`
    * dot: `.`
* Default separator: `-`
* Unacceptable leading and trailing dashes.

You can set custom configuration for use due to your needs. See more in docs.

## Roadmap

<b>v0.2.0</b>

* Improve library code style
* Add Go benchmarks
* Compile comprehensive documentation
* <b>Key feature</b>: custom replacements with Fluent style

<b>Stable release</b>

* <b>Major</b>: Multiple languages transliteration support
* <b>Minor</b>: 
    * support `Prefix()` & `Suffix()`
    * generating dates for slug (`UNIX`, `Datetime`)
    * emoji replacements 

## License

This project is licensed under the terms of the MIT license. See [LICENSE](./LICENSE) file for details.
