# go-project-template

[![codecov](https://codecov.io/gh/chyroc/go-project-template/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/go-project-template)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/go-project-template "go report card")](https://goreportcard.com/report/github.com/chyroc/go-project-template)
[![test status](https://github.com/chyroc/go-project-template/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/go-project-template/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/go-project-template)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Fgo-project-template.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Fgo-project-template)

![](./header.png)

## Install

```shell
go get github.com/chyroc/go-project-template
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/chyroc/go-project-template"
)

func main() {
	res := go_project_template.Incr(1)
	fmt.Println(res) // output: 2
}
```
