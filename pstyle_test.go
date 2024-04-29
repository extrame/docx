package docx

import (
	"encoding/xml"
	"testing"
)

func TestParsePstyle(t *testing.T) {
	var style = `<w:pStyle w:val="2" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"/>`
	var pstyle = &StrValueNode{}
	err := xml.Unmarshal([]byte(style), pstyle)
	if err != nil {
		t.Errorf("Error parsing pstyle: %s", err)
	}
	if pstyle.Val != "2" {
		t.Errorf("Expected pstyle val to be 2, got %s", pstyle.Val)
	}

}
