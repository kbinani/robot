package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("ax_darwin.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(bufio.NewReader(f))
	reg := regexp.MustCompile(`^\s*([A-Za-z0-9]*)Attribute\s*=.*// (.*)$`)
	for s.Scan() {
		line := s.Text()
		result := reg.FindSubmatch([]byte(line))
		if len(result) == 0 {
			continue
		}
		attr := string(result[1])
		if strings.Contains(attr, "Parameterized") {
			continue
		}
		t := string(result[2])
		switch t {
		case "CFString":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() string {\n", attr)
			fmt.Printf("\tret, _ := ref.StringAttr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "CFBoolean":
			name := fmt.Sprintf("Is%s", attr)
			if attr[:2] == "Is" {
				name = attr
			}
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() bool {\n", name)
			fmt.Printf("\tret, _ := ref.BoolAttr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "AXUIElement":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() *UIElement {\n", attr)
			fmt.Printf("\tret, _ := ref.UIElementAttr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "CFArray<AXUIElement>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() []*UIElement {\n", attr)
			fmt.Printf("\treturn ref.SliceOfUIElementAttr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "CFArray<AXValue<CFRange>>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() []Range {\n", attr)
			fmt.Printf("\treturn ref.SliceOfRangeAttr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "CFNumber<SInt64>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() int64 {\n", attr)
			fmt.Printf("\tret, _ := ref.Int64Attr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "CFNumber<Float32>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() float32 {\n", attr)
			fmt.Printf("\tret, _ := ref.Float32Attr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "CFNumber<Float64>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() float64 {\n", attr)
			fmt.Printf("\tret, _ := ref.Float64Attr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "AXValue<CFRange>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() Range {\n", attr)
			fmt.Printf("\tret, _ := ref.RangeAttr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "AXValue<CGPoint>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() Point {\n", attr)
			fmt.Printf("\tret, _ := ref.PointAttr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "AXValue<CGSize>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() Size {\n", attr)
			fmt.Printf("\tret, _ := ref.SizeAttr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "AXValue<CGRect>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() Rect {\n", attr)
			fmt.Printf("\tret, _ := ref.RectAttr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "CFURL":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() *url.URL {\n", attr)
			fmt.Printf("\tret, _ := ref.URLAttr(%sAttribute)\n", attr)
			fmt.Printf("\treturn ret\n")
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "CFArray<CFString>":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() []string {\n", attr)
			fmt.Printf("\treturn ref.SliceOfStringAttr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
		case "[same as kAXValueAttribute]":
			fmt.Printf("// %s returns kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %s() interface{} {\n", attr)
			fmt.Printf("	return ref.attrAny(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
			fmt.Printf("// %sAsInt32 returns explicitly int32 casted kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %sAsInt32() (int32, error) {\n", attr)
			fmt.Printf("	return ref.Int32Attr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
			fmt.Printf("// %sAsInt64 returns explicitly int64 casted kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %sAsInt64() (int64, error) {\n", attr)
			fmt.Printf("	return ref.Int64Attr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
			fmt.Printf("// %sAsFloat32 returns explicitly float32 casted kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %sAsFloat32() (float32, error) {\n", attr)
			fmt.Printf("	return ref.Float32Attr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
			fmt.Printf("// %sAsFloat64 returns explicitly float64 casted kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %sAsFloat64() (float64, error) {\n", attr)
			fmt.Printf("	return ref.Float64Attr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
			fmt.Printf("// %sAsBool returns explicitly bool casted kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %sAsBool() (bool, error) {\n", attr)
			fmt.Printf("	return ref.BoolAttr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
			fmt.Printf("// %sAsString returns explicitly string casted kAX%sAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibility%sattribute\n", attr, attr, strings.ToLower(attr))
			fmt.Printf("func (ref *UIElement) %sAsString() (string, error) {\n", attr)
			fmt.Printf("	return ref.StringAttr(%sAttribute)\n", attr)
			fmt.Printf("}\n")
			fmt.Printf("\n")
		default:
			fmt.Fprintf(os.Stderr, "%s %s\n", attr, t)
		}
	}
}
