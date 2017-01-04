package main

import (
	"fmt"
	"os"
	"text/template"

	"strings"

	"github.com/kelvyne/kaykin/d2protocolparser"
	"math"
)

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

	if err := exportMessages(p.Messages); err != nil {
		fmt.Fprintf(os.Stderr, "messages: %v\n", err)
		os.Exit(3)
	}
	if err := exportTypes(p.Types); err != nil {
		fmt.Fprintf(os.Stderr, "types: %v\n", err)
		os.Exit(3)
	}
}

func exportMessages(m []d2protocolparser.Class) error {
	return nil
}

func exportTypes(types []d2protocolparser.Class) error {
	const typesTemplateFile = "./types.template"
	funcMap := template.FuncMap{
		"ToFieldName":   strings.Title,
		"ToGolangType":  toGolangType,
		"TypeHasParent": typeHasParent,
		"EffectiveType": getEffectiveType,
		"UseBooleanByteWrapper": useBooleanByteWrapper,
		"BBWIterate": rangeBBW,
		"IsNativeType": isNativeType,
		"DeserializeMethod": getDeserializeMethod,
		"DeserializeLengthMethod": getDeserializeLengthMethod,
	}
	tem := template.Must(template.New("types.template").Funcs(funcMap).ParseFiles(typesTemplateFile))
	if err := tem.Execute(os.Stdout, types); err != nil {
		return err
	}
	return nil
}

func toGolangType(t string) string {
	return t
}

func typeHasParent(t d2protocolparser.Class) bool {
	return t.Parent != ""
}

func getEffectiveType(t d2protocolparser.Field) string {
	if isNativeType(t) {
		return t.Type
	}
	return "*" + t.Type
}

func useBooleanByteWrapper(t d2protocolparser.Class) bool {
	for _, f := range t.Fields {
		return f.UseBBW
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
	} else if (t.Type == "bool") {
		return true
	} else if (t.Type == "string") {
		return true
	}
	return false
}

func getDeserializeMethod(t d2protocolparser.Field) string {
	return "Read" + t.Method
}

var lengthTranslator = map[string]string {
	"writeShort": "Int16",
	"writeInt": "Int32",
	"writeVarShort": "VarInt16",
	"writeVarInt": "VarInt32",
}

func getDeserializeLengthMethod(t d2protocolparser.Field) string {
	m, ok := lengthTranslator[t.WriteLengthMethod]
	if !ok {
		panic(fmt.Sprintf("%v length method is not valid", t.WriteMethod))
	}
	return "Read" + m
}