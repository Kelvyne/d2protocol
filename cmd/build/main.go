package main

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"strings"

	"math"

	"github.com/kelvyne/kaykin/d2protocolparser"
)

func openFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Create: %v", err)
		os.Exit(3)
	}
	return f
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %v path.swf\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]
	p, err := d2protocolparser.Build(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "build: %v\n", err)
		os.Exit(2)
	}

	typesWriter := openFile("../../types.gen.go")
	if err = exportClasses(typesWriter, "types", p.Types); err != nil {
		fmt.Fprintf(os.Stderr, "types: %v\n", err)
		os.Exit(4)
	}

	messagesWriter := openFile("../../messages.gen.go")
	if err = exportClasses(messagesWriter, "messages", p.Messages); err != nil {
		fmt.Fprintf(os.Stderr, "messages: %v\n", err)
		os.Exit(4)
	}

	protocolWriter := openFile("../../protocol.gen.go")
	if err = exportProtocol(protocolWriter, p); err != nil {
		fmt.Fprintf(os.Stderr, "protocol: %v\n", err)
		os.Exit(4)
	}
}

func exportProtocol(w io.Writer, p *d2protocolparser.Protocol) error {
	const templateFile = "./protocol.template"
	tem := template.Must(template.ParseFiles(templateFile))
	if err := tem.Execute(w, p); err != nil {
		return err
	}
	return nil
}

func exportClasses(w io.Writer, name string, types []d2protocolparser.Class) error {
	funcMap := template.FuncMap{
		"MapName":                 func() string { return name },
		"ToFieldName":             strings.Title,
		"ToGolangType":            toGolangType,
		"TypeHasParent":           typeHasParent,
		"EffectiveType":           getEffectiveType,
		"UseBooleanByteWrapper":   useBooleanByteWrapper,
		"BBWIterate":              rangeBBW,
		"IsNativeType":            isNativeType,
		"DeserializeMethod":       getDeserializeMethod,
		"SerializeMethod":         getSerializeMethod,
		"DeserializeLengthMethod": getDeserializeLengthMethod,
		"SerializeLengthMethod":   getSerializeLengthMethod,
		"CastType":                getCastType,
	}
	const typesTemplateFile = "./class.template"
	tem := template.Must(template.New("class.template").Funcs(funcMap).ParseFiles(typesTemplateFile))
	if err := tem.Execute(w, types); err != nil {
		return err
	}
	return nil
}

func toGolangType(t string) string {
	return t
}

func typeHasParent(t d2protocolparser.Class) bool {
	return t.Parent != "" && t.Parent != "NetworkMessage"
}

func getEffectiveType(t d2protocolparser.Field) string {
	if isNativeType(t) {
		return t.Type
	}
	return "*" + t.Type
}

func useBooleanByteWrapper(t d2protocolparser.Class) bool {
	for _, f := range t.Fields {
		if f.UseBBW {
			return true
		}
	}
	return false
}

func rangeBBW(t d2protocolparser.Class) [][]d2protocolparser.Field {
	var bbwFields []d2protocolparser.Field
	for _, field := range t.Fields {
		if field.UseBBW {
			bbwFields = append(bbwFields, field)
		}
	}

	count := int(math.Ceil(float64(len(bbwFields)) / 8.0))
	r := make([][]d2protocolparser.Field, count)
	for i, field := range bbwFields {
		index := i / 8
		r[index] = append(r[index], field)
	}
	return r
}

func isNativeType(t d2protocolparser.Field) bool {
	if strings.HasPrefix(t.Type, "int") || strings.HasPrefix(t.Type, "uint") {
		return true
	} else if strings.HasPrefix(t.Type, "float") {
		return true
	} else if t.Type == "bool" {
		return true
	} else if t.Type == "string" {
		return true
	}
	return false
}

func getDeserializeMethod(t d2protocolparser.Field) string {
	return "Read" + t.Method
}

func getSerializeMethod(t d2protocolparser.Field) string {
	return "Write" + t.Method
}

var lengthTranslator = map[string]string{
	"writeShort":    "Int16",
	"writeInt":      "Int32",
	"writeVarShort": "VarInt16",
	"writeVarInt":   "VarInt32",
}

func getDeserializeLengthMethod(t d2protocolparser.Field) string {
	m, ok := lengthTranslator[t.WriteLengthMethod]
	if !ok {
		panic(fmt.Sprintf("%v length method is not valid", t.WriteMethod))
	}
	return "Read" + m
}

func getSerializeLengthMethod(t d2protocolparser.Field) string {
	m, ok := lengthTranslator[t.WriteLengthMethod]
	if !ok {
		panic(fmt.Sprintf("%v length method is not valid", t.WriteMethod))
	}
	return "Write" + m
}

func getCastType(method string) string {
	castMap := map[string]string{
		"writeShort":    "int16",
		"writeInt":      "int32",
		"writeVarShort": "int16",
		"writeVarInt":   "int32",
	}
	t, ok := castMap[method]
	if !ok {
		panic(fmt.Sprintf("cannot cast method %v", method))
	}
	return t
}
