package map2xml

import (
	"strings"
	"testing"
)

func TestMarshalSimple(t *testing.T) {
	expected := `<person mood="happy">
  <people>cheer</people>
</person>`
	config := New(map[string]interface{}{
		"people": "cheer",
	})

	config.WithIndent("", "  ")
	config.WithRoot("person", map[string]string{"mood": "happy"})
	xmlBytes, err := config.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	if string(xmlBytes) != expected {
		t.Logf("%s", xmlBytes)
		t.Fail()
	}
}

func TestMarshalSimpleSort(t *testing.T) {
	expected := `<person mood="happy">
  <0>sad</0>
  <8>cry</8>
  <a>cheer</a>
  <b>happy</b>
</person>`
	config := New(map[string]interface{}{
		"8": "cry",
		"b": "happy",
		"0": "sad",
		"a": "cheer",
	})

	config.WithIndent("", "  ")
	config.WithSortedKeys()
	config.WithRoot("person", map[string]string{"mood": "happy"})
	xmlBytes, err := config.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	if string(xmlBytes) != expected {
		t.Logf("%s", xmlBytes)
		t.Fail()
	}
}
func TestMarshalSimpleNoSort(t *testing.T) {
	expected := []string{
		"<person mood=\"happy\">",
		"  <0>sad</0>",
		"  <8>cry</8>",
		"  <a>cheer</a>",
		"  <b>happy</b>",
		"</person>",
	}
	config := New(map[string]interface{}{
		"8": "cry",
		"b": "happy",
		"0": "sad",
		"a": "cheer",
	})

	config.WithIndent("", "  ")
	config.WithRoot("person", map[string]string{"mood": "happy"})
	xmlBytes, err := config.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	ans := string(xmlBytes)

	for _, line := range expected {
		if ok := strings.Contains(ans, line); !ok {
			t.Logf("Line: %s does not exist", line)
			t.Fail()
		}
	}
}

func TestAllTypes(t *testing.T) {
	var value = map[string]interface{}{
		"test": map[string]interface{}{
			"num":    2,
			"string": "hi",
			"nil":    nil,
			"slices": []map[string]interface{}{
				{"xml_child_name": "slice", "one": 1},
				{"two": 2},
				{"three": 3},
			},
			"multi": map[string]interface{}{
				"level": "value",
			},
		},
	}
	var expected = `<test>
  <multi>
    <level>value</level>
  </multi>
  <nil></nil>
  <num>2</num>
  <slices>
    <slice>
      <one>1</one>
    </slice>
    <slice>
      <two>2</two>
    </slice>
    <slice>
      <three>3</three>
    </slice>
  </slices>
  <string>hi</string>
</test>`

	xml := New(value)
	xml.WithIndent("", "  ")
	xml.WithSortedKeys()
	ans, err := xml.MarshalToString()
	if err != nil {
		t.Error(err)
	}
	if ans != expected {
		t.Logf("expected:\n---\n'%s'\n---\nvalue:\n+++\n%s\n+++\n", expected, ans)
		t.Fail()
	}
}
