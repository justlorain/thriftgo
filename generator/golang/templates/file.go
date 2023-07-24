// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templates

// File .
var File = `// Code generated by thriftgo ({{Version}}). DO NOT EDIT.
{{InsertionPoint "bof"}}

package {{.FilePackage}}

import (
	{{InsertionPoint "imports"}}
	{{- if Features.GenerateReflectionInfo}}thriftreflection "github.com/cloudwego/kitex/pkg/reflection/thrift"{{end}}
)

{{template "Constant" .}}

{{- range .Enums}}
{{template "Enum" .}}
{{- end}}

{{- range .Typedefs}}
{{template "Typedef" .}}
{{- end}}

{{- range .Structs}}
{{template "StructLike" .}}
{{- end}}

{{- range .Unions}}
{{template "StructLike" .}}
{{- end}}

{{- range .Exceptions}}
{{template "StructLike" .}}
{{- end}}

{{- range .Services}}
{{template "Service" .}}
{{template "Client" .}}
{{- end}}

{{- range .Services}}
{{template "Processor" .}}
{{- end}}

{{- if Features.GenerateReflectionInfo}}
	var file_{{.IDLName}}_rawDesc = {{.IDLMeta}}
	func init(){
		thriftreflection.RegisterIDL(file_{{.IDLName}}_rawDesc)
	}
{{end}}

{{- if Features.WithReflection}}
{{- UseStdLibrary "thrift_reflection"}}

{{$IDLName := .IDLName}}
{{$IDLPath := .AST.Filename}}
{{$FilePackage := .FilePackage}}

var file_{{$IDLName}}_thrift_go_types = []interface{}{
	{{- range $index, $element := .Structs}}
	(*{{.GoName}})(nil),	// Struct {{$index}}: {{$FilePackage}}.{{.Name}}
	{{- end}}
	{{- range $index, $element := .Unions}}
	(*{{.GoName}})(nil),	// Union {{$index}}: {{$FilePackage}}.{{.Name}}
	{{- end}}
	{{- range $index, $element := .Exceptions}}
	(*{{.GoName}})(nil),	// Exception {{$index}}: {{$FilePackage}}.{{.Name}}
	{{- end}}
	{{- range $index, $element := .Enums}}
	(*{{.GoName}})(nil),	// Enum {{$index}}: {{$FilePackage}}.{{.Name}}
	{{- end}}
}
var file_{{$IDLName}}_thrift *thrift_reflection.FileDescriptor
var file_idl_{{$IDLName}}_rawDesc = {{.MarshalDescriptor}}

func init() { 
	if file_{{$IDLName}}_thrift != nil {
		return
	}
	file_{{$IDLName}}_thrift = thrift_reflection.BuildFileDescriptor(file_idl_{{$IDLName}}_rawDesc,file_{{$IDLName}}_thrift_go_types)
}

func GetFileDescriptorFor{{ToCamel $IDLName}}() *thrift_reflection.FileDescriptor{
	return file_{{$IDLName}}_thrift
}

{{- range .Structs}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetStructDescriptor("{{.Name}}")
}
{{- end}}
{{- range .Enums}}
func (p {{.GoName}}) GetDescriptor() *thrift_reflection.EnumDescriptor{
	return file_{{$IDLName}}_thrift.GetEnumDescriptor("{{.Name}}")
}
{{- end}}
{{- range .Typedefs}}
func GetTypeDescriptorFor{{.GoName}}() *thrift_reflection.TypedefDescriptor{
	return file_{{$IDLName}}_thrift.GetTypedefDescriptor("{{.Alias}}")
}
{{- end}}
{{- range .Constants.GoConstants}}
func GetConstDescriptorFor{{.GoName}}() *thrift_reflection.ConstDescriptor{
	return file_{{$IDLName}}_thrift.GetConstDescriptor("{{.Name}}")
}
{{- end}}
{{- range .Unions}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetUnionDescriptor("{{.Name}}")
}
{{- end}}
{{- range .Exceptions}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetExceptionDescriptor("{{.Name}}")
}
{{- end}}
{{- range .Services}}
{{$ServiceName := .GoName}}
func GetServiceDescriptorFor{{.GoName}}() *thrift_reflection.ServiceDescriptor{
	return file_{{$IDLName}}_thrift.GetServiceDescriptor("{{.Name}}")
}
{{- range .Functions}}
func GetMethodDescriptorFor{{$ServiceName}}{{.GoName}}() *thrift_reflection.MethodDescriptor{
	return file_{{$IDLName}}_thrift.GetMethodDescriptor("{{$ServiceName}}","{{.Name}}")
}
{{- end}}
{{- end}}

{{end}}
{{- InsertionPoint "eof"}}
`
