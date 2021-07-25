package map2xml

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

type Indentation struct {
	Prefix string
	Indent string
}

type StructMap struct {
	CData           bool
	XMLName         *xml.Name
	Map             map[string]interface{}
	Indent          *Indentation
	StartAttributes *[]xml.Attr
}

type xmlMapEntry struct {
	XMLName    xml.Name
	Value      interface{} `xml:",innerxml"`
	CDataValue interface{} `xml:",cdata"`
}

// func main() {
// 	ori := map[string]interface{}{
// 		"a": 1,
// 		"b": "abekat",
// 		"tivoli": map[string]interface{}{
// 			"int":  42,
// 			"crap": false,
// 		},
// 	}
// 	awesome := New(ori, "goodfellas").AsCData().WithIndent("", "  ").WithStartAttributes(map[string]string{"status": "awesome"})
// 	str, err := awesome.MarshalToString()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(str)
// }

func New(input map[string]interface{}, xmlName string) *StructMap {
	return &StructMap{Map: input, XMLName: &xml.Name{Local: xmlName}}
}

func (smap *StructMap) WithIndent(prefix string, indent string) *StructMap {
	smap.Indent = &Indentation{Prefix: prefix, Indent: indent}
	return smap
}

func (smap *StructMap) WithStartAttributes(xmlStartAttributes map[string]string) *StructMap {
	attr := []xml.Attr{}
	for k, v := range xmlStartAttributes {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: k}, Value: v})
	}
	smap.StartAttributes = &attr
	return smap
}

func (smap *StructMap) AsCData() *StructMap {
	smap.CData = true
	return smap
}

func (smap *StructMap) Print() {
	var indent interface{} = smap.Indent
	var attr interface{} = smap.StartAttributes
	if smap.Indent != nil {
		indent = map[string]int{"indent_spaces": len(*&smap.Indent.Indent), "prefix_spaces": len(*&smap.Indent.Prefix)}
	}
	var rootName interface{} = smap.XMLName
	if rootName == nil {
		rootName = "none"
	}
	fmt.Printf("XML root name: %v\nCDATA: %v\nIndentation: %v\nStart Attributes: %v\n",
		rootName, smap.CData, indent, attr)
}

func (smap *StructMap) Marshal() ([]byte, error) {
	if smap.Indent == nil {
		return xml.Marshal(smap)
	} else {
		return xml.MarshalIndent(smap, smap.Indent.Prefix, smap.Indent.Indent)
	}
}

func (smap *StructMap) MarshalToString() (string, error) {
	xmlBytes, err := smap.Marshal()
	return string(xmlBytes), err
}

func (smap StructMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(smap.Map) == 0 {
		return nil
	}
	if smap.XMLName != nil {
		start = xml.StartElement{Name: *smap.XMLName, Attr: start.Attr}
	}

	if smap.StartAttributes != nil {
		start.Attr = *smap.StartAttributes
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for k, v := range smap.Map {
		if err := handleChildren(e, k, v, smap.CData); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

func handleChildren(e *xml.Encoder, fieldName string, v interface{}, cdata bool) error {
	if reflect.TypeOf(v).Kind() == reflect.Map {
		e.EncodeToken(xml.StartElement{Name: xml.Name{Local: fieldName}})
		for key, val := range v.(map[string]interface{}) {
			handleChildren(e, key, val, cdata)
		}
		return e.EncodeToken(xml.EndElement{Name: xml.Name{Local: fieldName}})
	}
	if cdata {
		return e.Encode(xmlMapEntry{XMLName: xml.Name{Local: fieldName}, CDataValue: v})
	} else {
		return e.Encode(xmlMapEntry{XMLName: xml.Name{Local: fieldName}, Value: fmt.Sprintf("%v", v)})
	}
}
