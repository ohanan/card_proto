package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) {
	var existOneOf bool
	var checkOneOf func(messages []*protogen.Message) bool
	checkOneOf = func(messages []*protogen.Message) bool {
		if existOneOf {
			return true
		}
		for _, message := range messages {
			if len(message.Oneofs) > 0 {
				return true
			}
			if checkOneOf(message.Messages) {
				return true
			}
		}
		return false
	}
	if !checkOneOf(file.Messages) && len(file.Services) == 0 {
		return
	}
	filename := file.GeneratedFilenamePrefix + "_custom.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	contextIdent := protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	}
	callOptionIdent := protogen.GoIdent{
		GoName:       "CallOption",
		GoImportPath: "google.golang.org/grpc",
	}
	g.P("// Code generated by protoc-gen-card-proto-plugin. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	for _, message := range file.Messages {
		for _, oneof := range message.Oneofs {
			for _, field := range oneof.Fields {
				goType, pointer := fieldGoType(g, field)
				if pointer {
					// goType = "*" + goType
				}
				g.P("func(x *", message.GoIdent, ") With", field.GoName, "(v ", goType, ")*", message.GoIdent, "{")
				g.P("x.", oneof.GoName, " = &", field.GoIdent, "{")
				g.P(field.GoName, ": v,")
				g.P("}")
				g.P("return x")
				g.P("}")
			}
		}
	}
	for _, service := range file.Services {
		serviceName := service.GoName
		isPlugin := serviceName == "Plugin"
		g.P("type ", serviceName, "XClient", " struct {")
		g.P("c ", serviceName+"Client")
		g.P("}")
		g.P("func New", serviceName, "XClient(c ", serviceName, "Client", ")*", serviceName, "XClient{")
		g.P("return &", serviceName, "XClient{c:c}")
		g.P("}")
		for _, method := range service.Methods {
			var signatures []any
			signatures = append(signatures, "func (x ", serviceName, "XClient)", method.GoName, "(")
			for i, field := range method.Input.Fields {
				if i > 0 {
					signatures = append(signatures, ",")
				}
				signatures = append(signatures, privateName(field.GoName), " ")
				goType, pointer := fieldGoType(g, field)
				if pointer {
					signatures = append(signatures, "*")
				}
				signatures = append(signatures, goType)
			}
			signatures = append(signatures, ")*", method.Output.GoIdent, "{")
			g.P(signatures...)
			g.P("r, err := x.c.", method.GoName, "(", protogen.GoIdent{
				GoName:       "Background",
				GoImportPath: "context",
			}, "(), &", method.Input.GoIdent, "{")
			for _, field := range method.Input.Fields {
				g.P(field.GoName, ": ", privateName(field.GoName), ",")
			}
			g.P("})")
			g.P("if err != nil {")
			g.P("panic(err)")
			g.P("}")
			g.P("return r")
			g.P("}")
			g.P("func(x ", serviceName, "XClient)", method.GoName, "0(req *", method.Input.GoIdent, ")*", method.Output.GoIdent, "{")
			g.P("r, err := x.c.", method.GoName, "(", protogen.GoIdent{
				GoName:       "Background",
				GoImportPath: "context",
			}, "(), req)")
			g.P("if err != nil {")
			g.P("panic(err)")
			g.P("}")
			g.P("return r")
			g.P("}")
		}
		g.P("type ", serviceName, "XServer interface {")
		for _, method := range service.Methods {
			var signatures []any
			signatures = append(signatures, method.GoName, "(")
			if isPlugin {
				signatures = append(signatures, "helper *Helper, ")
			}
			signatures = append(signatures, "req *", method.Input.GoIdent, ", resp *", method.Output.GoIdent, ")")
			g.P(signatures...)
		}
		g.P("}")
		g.P("func New", serviceName, "ClientFromServer(x ", serviceName, "Server) ", serviceName, "Client{")
		g.P("return &", privateName(serviceName), "ClientByServer{s: x}")
		g.P("}")
		serverToClientName := privateName(serviceName) + "ClientByServer"
		g.P("type ", serverToClientName, " struct {s ", serviceName, "Server}")
		for _, method := range service.Methods {
			g.P("func(x ", serverToClientName, ")", method.GoName, "(ctx ", contextIdent, ", req *", method.Input.GoIdent, ", opts...", callOptionIdent, ")(*", method.Output.GoIdent, ", error){")
			g.P("return x.s.", method.GoName, "(ctx, req)")
			g.P("}")
		}

		xServerName := privateName(serviceName) + "XServerAdapter"
		g.P("func New", serviceName, "ServerFromXServer(server ", serviceName, "XServer)", serviceName, "Server {")
		g.P("return &", xServerName, "{Server: server}")
		g.P("}")
		g.P("type  ", xServerName, " struct{ ")
		g.P("Server ", serviceName, "XServer")
		if isPlugin {
			g.P("*Helper")
		}
		g.P("}")
		for _, method := range service.Methods {
			g.P("func(x *", xServerName, ") ", method.GoName, "(ctx ", contextIdent, ", req *", method.Input.GoIdent, ")(*", method.Output.GoIdent, ", error){")
			g.P("resp := &", method.Output.GoIdent, "{}")
			caller := "(req, resp)"
			if isPlugin {
				caller = "(x.Helper, req, resp)"
			}
			g.P("x.Server.", method.GoName, caller)
			g.P("return resp, nil")
			g.P("}")
		}
	}
}

func privateName(v string) string {
	if v == "" {
		return ""
	}
	return strings.ToLower(v[:1]) + v[1:]
}
func fieldGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := fieldGoType(g, field.Message.Fields[0])
		valType, _ := fieldGoType(g, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}
