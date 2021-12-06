package map2xml

import (
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
