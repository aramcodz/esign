// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package templates contains text of service and model templates for
// generating docusign api packages
package templates

// Service is the default template for constructing service package files
const Service = `// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package {{.Package}} implements the DocuSign SDK
// category {{.Service}}.
// {{range .Comments}}
// {{.}}{{end}}
//
// {{$callVersion := .CallVersion}}{{$verPrefix := .VersionID}}{{$docPrefix := .DocPrefix}}{{if .AddDocLinks}}Service Api documentation may be found at:
// https://developers.docusign.com/docs/{{$docPrefix}}reference/{{.Service}}{{end}}
// Usage example:
//
//   import (
//       "{{.Directory}}"{{if .ModelIsPackage}}
//       "{{.Directory}}/{{.ModelPackagePath}}"{{end}}
//   ) 
//   ...
//   {{.Package}}Service := {{.Package}}.New(esignCredential)
package {{.Package}} // import "github.com/jfcote87/esign{{.PackagePath}}/{{.Package}}"

import ({{range .Packages}}
    {{.}}{{end}}
)

// Service implements DocuSign {{.Service}} Category API operations
type Service struct {
    credential esign.Credential 
}

// New initializes a {{.Package}} service using cred to authorize ops.
func New(cred esign.Credential) *Service {
    return &Service{credential: cred}
}

{{ range .Operations }}{{$accept:=.Accept}}{{range .CommentLines .FuncName $docPrefix .HasUploads .IsMediaUpload}}// {{.}}
{{end}}func (s *Service) {{.FuncName}}({{range $i, $p := .PathParams}}{{if $i}}, {{end}}{{$p.GoName}} string{{end}}{{if .OpPayload}}{{if len .PathParams}}, {{end}}{{if .IsMediaUpload}}media io.Reader, mimeType string{{else}}{{.OpPayload.GoName}} {{.OpPayload.Type}}{{end}}{{end}}{{if .HasUploads}}, uploads ...*esign.UploadFile{{end}}) *{{.FuncName}}Op {
    return &{{.FuncName}}Op{
        Credential: s.credential,
        Method:  "{{.HTTPMethod}}",
        Path: {{.OpPath2 $verPrefix .PathParams}},
        {{if .OpPayload}}Payload: {{if .IsMediaUpload}}&esign.UploadFile{ Reader: media, ContentType: mimeType }{{else}}{{.OpPayload.GoName}}{{end}},
        {{end}}{{if .HasUploads}}Files: uploads,
        {{end}}{{if $accept}}Accept: "{{$accept}}",
        {{end}}QueryOpts: make(url.Values),{{if $callVersion}}
        Version: {{$callVersion}},{{end}}
    }
}

// {{.FuncName}}Op implements DocuSign API SDK {{.SDK}}
type {{.FuncName}}Op esign.Op

// Do executes the op.  A nil context will return error.
func (op *{{.FuncName}}Op) Do(ctx context.Context)  {{if .Result}}({{.Result}}, error){{else}}error{{end}} {
    {{if .Result}}var res {{.Result}}
    {{end}}return {{if .Result}}res, {{end}}((*esign.Op)(op)).Do(ctx, {{if .Result}}&res{{else}}nil{{end}})
}

{{$funcName := .FuncName}}{{range .QueryOptions}}{{range .Comments}}// {{.}}
{{end}}func (op *{{$funcName}}Op) {{.GoName}}({{if ne .Type "bool"}}val {{.Type}}{{end}}) *{{$funcName}}Op {
    if op != nil {
        op.QueryOpts.Set("{{.Name}}", {{.Value}})
    }
    return op
}

{{end}}{{range .DownloadAdditions}}
{{ range .Comments}}// {{.}}
{{end}}func (op *{{$funcName}}Op) {{.Name}}(ctx context.Context) (*esign.Download, error) {
    var res *esign.Download
    if op == nil {
        return nil, esign.ErrNilOp
    }
    newOp := esign.Op(*op)
    newOp.Accept = "{{.MimeType}}"
    return res, (&newOp).Do(ctx, &res)
}

{{end}}{{end}}`

// Model is the default template for defining packages input and output structures
const Model = `// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT. {{$defMap := .DefMap}}{{$fldOverrides := .FldOverrides}}
{{if .IsPackage}}
// Package model provides definitions for all input
// and output parameters types found in DocuSign's 
// OpenAPI (swagger) file.
//
// Api documentation may be found at:
// https://developers.docusign.com/docs/{{$docPrefix := .DocPrefix}}{{$docPrefix}}reference
package {{.ModelPackage}} // import "github.com/jfcote87/esign/{{.ModelPackagePath}}{{else}}package {{.ModelPackage}}
{{end}}{{if .ModelImports}}

import ({{range .ModelImports}}
    "{{.}}"{{end}}
){{end}}
{{ range .Definitions }}{{if len .CommentLines}}{{ range .CommentLines}}
// {{.}}{{end}}{{end}}
type {{.StructName}} struct { {{range .StructFields $defMap $fldOverrides}}
    {{range .Comments}}// {{.}}
    {{end}}{{.Name}}{{if .Type}} {{.Type}}{{end}}{{if .JSON}} ` + "`json:\"{{.JSON}},omitempty\"`" + `{{end}}{{end }}
}
{{ end }}
{{.CustomCode}}`
