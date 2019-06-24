package xproj

import (
	"github.com/typical-go/typical-go/appx"
	"github.com/typical-go/typical-rest-server/typical"
)

// TODO: readme template should outside of binary

var readmeData = struct {
	Context       appx.Context
	Configuration string
}{
	Context:       typical.Context,
	Configuration: configurationTable(),
}

const readmeTemplate = `<!-- Autogenerated by Typical-Go; Modify 'typical/readme_template.go' to add more content -->
# {{ .Context.Name}}

{{ .Context.Description}}

## Configuration

{{ .Configuration}}
`