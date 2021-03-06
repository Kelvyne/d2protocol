package main

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"strings"

	"math"

	"sort"

	"github.com/kelvyne/d2protocolparser"
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

	interfacesWriter := openFile("../../interfaces.gen.go")
	if err = exportInterfaces(interfacesWriter, p.Types, p.Messages); err != nil {
		fmt.Fprintf(os.Stderr, "interfaces: %v\n", err)
		os.Exit(4)
	}

	protocolWriter := openFile("../../protocol.gen.go")
	if err = exportProtocol(protocolWriter, p); err != nil {
		fmt.Fprintf(os.Stderr, "protocol: %v\n", err)
		os.Exit(4)
	}
}

func exportInterfaces(w io.Writer, types []d2protocolparser.Class, messages []d2protocolparser.Class) error {
	needingIntrf := map[string]bool{}
	for _, t := range types {
		for _, field := range t.Fields {
			if field.UseTypeManager {
				needingIntrf[field.Type] = true
			}
		}
	}
	for _, t := range messages {
		for _, field := range t.Fields {
			if field.UseTypeManager {
				needingIntrf[field.Type] = true
			}
		}
	}
	interfaces := []d2protocolparser.Class{}
	for _, t := range types {
		if f, _ := needingIntrf[t.Name]; f {
			interfaces = append(interfaces, t)
		}
	}

	funcMap := template.FuncMap{
		"ToFieldName":     strings.Title,
		"ToInterfaceName": toInterfaceName,
		"ToGetterName":    toGetterName,
		"EffectiveType":   getEffectiveType,
		"TypeHasParent":   typeHasParent,
	}
	const interfacesTemplateFile = "./interfaces.template"
	tem := template.Must(template.New("interfaces.template").Funcs(funcMap).ParseFiles(interfacesTemplateFile))
	if err := tem.Execute(w, interfaces); err != nil {
		return err
	}
	return nil
}

func exportProtocol(w io.Writer, p *d2protocolparser.Protocol) error {
	const templateFile = "./protocol.template"
	tem := template.Must(template.ParseFiles(templateFile))
	if err := tem.Execute(w, p); err != nil {
		return err
	}
	return nil
}

type byId []d2protocolparser.Class

func (a byId) Len() int           { return len(a) }
func (a byId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byId) Less(i, j int) bool { return a[i].ProtocolID < a[j].ProtocolID }

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
	sort.Sort(byId(types))
	if err := tem.Execute(w, types); err != nil {
		return err
	}
	return nil
}

func toGetterName(field string) string {
	return "Get" + field
}

func toInterfaceName(name string) string {
	return name + "Intrf"
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
	} else if t.UseTypeManager {
		return toInterfaceName(t.Type)
	}
	return t.Type
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
