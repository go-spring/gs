package internal

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestProjectXml(t *testing.T) {

	result := `<?xml version="1.0" encoding="UTF-8"?>
<projects>
  <project>
    <name>spring-core</name>
    <dir>spring/spring-core</dir>
    <url>https://github.com/go-spring/spring-core.git</url>
    <branch>master</branch>
  </project>
  <project>
    <name>spring-boot</name>
    <dir>spring/spring-boot</dir>
    <url>https://github.com/go-spring/spring-boot.git</url>
    <branch>main</branch>
  </project>
</projects>`

	doc1 := ProjectXml{
		XMLName: xml.Name{Space: "", Local: "projects"},
		Projects: []Project{
			{
				Name:   "spring-core",
				Dir:    "spring/spring-core",
				Url:    "https://github.com/go-spring/spring-core.git",
				Branch: "master",
			},
			{
				Name:   "spring-boot",
				Dir:    "spring/spring-boot",
				Url:    "https://github.com/go-spring/spring-boot.git",
				Branch: "main",
			},
		},
	}

	bytes, _ := xml.MarshalIndent(doc1, "", "  ")
	data := xml.Header + string(bytes)
	if data != result {
		t.Errorf("%s != %s", data, result)
	}

	var doc2 ProjectXml
	_ = xml.Unmarshal([]byte(data), &doc2)
	if !reflect.DeepEqual(doc1, doc2) {
		t.Errorf("%v != %v", doc1, doc2)
	}

	if _, ok := doc2.Find("spring-boot"); !ok {
		t.Errorf("%s not found", "spring-boot")
	}

	if _, ok := doc2.Find("dummy"); ok {
		t.Errorf("%s not found", "dummy")
	}
}
