package main

import (
	"bytes"
	"os"
	"text/template"

	"github.com/yoheimuta/go-protoparser/v4"
	parser2 "github.com/yoheimuta/go-protoparser/v4/parser"
)

func getMethods(filename string) []string {
	reader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = reader.Close() }()

	got, err := protoparser.Parse(reader)
	if err != nil {
		panic(err)
	}
	var methods []string
	for _, v := range got.ProtoBody {
		m, ok := v.(*parser2.Message)
		if !ok {
			continue
		}
		var req, resp bool
		for _, vv := range m.MessageBody {
			mm, ok := vv.(*parser2.Message)
			if !ok {
				continue
			}
			if !req && mm.MessageName == "Req" {
				req = true
			}
			if !resp && mm.MessageName == "Resp" {
				resp = true
			}
		}
		if req && resp {
			methods = append(methods, m.MessageName)
		}
	}
	return methods
}

func main() {
	var bb bytes.Buffer
	if err := template.Must(template.New("rpc_proto").Parse(`syntax = "proto3";
package proto;

import "proto/card_methods.proto";
import "proto/plugin_methods.proto";

option go_package = "./protoservice";

service Card {
{{range .Cards}}  rpc {{.}}({{.}}.Req) returns ({{.}}.Resp);
{{end}}}

service Plugin {
{{range .Plugins}}  rpc {{.}}({{.}}.Req) returns ({{.}}.Resp);
{{end}}}
`)).Execute(&bb, map[string]interface{}{
		"Cards":   getMethods("proto/card_methods.proto"),
		"Plugins": getMethods("proto/plugin_methods.proto"),
	}); err != nil {
		panic(err)
	}
	if err := os.WriteFile("service.proto", bb.Bytes(), 0644); err != nil {
		panic(err)
	}
}
