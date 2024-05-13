package docx

import (
	"encoding/xml"
)

const (
	HYPERLINK_STYLE = "a1"
)

// A Run is part of a paragraph that has its own style. It could be
// a piece of text in bold, or a link
type Run struct {
	XMLName       xml.Name       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main r,omitempty"`
	RunProperties *RunProperties `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr,omitempty"`
	InstrText     string         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main instrText,omitempty"`
	Text          *Text
	no            int
}

// The Text object contains the actual text
type Text struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main t"`
	XMLSpace string   `xml:"xml:space,attr,omitempty"`
	Text     string   `xml:",chardata"`
}

// The hyperlink element contains links
type Hyperlink struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID      string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	Run     Run      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main r,omitempty"`
	no      int
}

// RunProperties encapsulates visual properties of a run
type RunProperties struct {
	XMLName  xml.Name      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr,omitempty"`
	Color    *StrValueNode `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,omitempty"`
	Size     *IntValueNode `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz,omitempty"`
	RunStyle *StrValueNode `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rStyle,omitempty"`
	Style    *StrValueNode `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pStyle,omitempty"`
	Fonts    *Fonts        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rFonts,omitempty"`
}

// Fonts contains the font family
type Fonts struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rFonts,omitempty"`
	Ascii    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main ascii,attr"`
	HAnsi    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hAnsi,attr"`
	EastAsia string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main eastAsia,attr"`
	Complex  string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main complex,attr"`
	Cs       string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main cs,attr"`
}

// Size contains the font size

type ParagraphProperties struct {
	XMLName xml.Name      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pPr"`
	Style   *StrValueNode `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pStyle,omitempty"`
	Spacing *Spacing      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main spacing,omitempty"`
	Ind     *Indent       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main ind,omitempty"`
	Jc      *StrValueNode `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main jc,omitempty"`
	Outline *StrValueNode `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main outlineLvl,omitempty"`
}

func (p *ParagraphProperties) GetStyleId() string {
	if p.Style != nil {
		return p.Style.Val
	}
	return ""
}

type Indent struct {
	XMLName    xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main ind,omitempty"`
	First      int      `xml:"http://schemas.openxmlformats.org/2006/main first,attr"`
	Hanging    int      `xml:"http://schemas.openxmlformats.org/2006/main hanging,attr"`
	Left       int      `xml:"http://schemas.openxmlformats.org/2006/main left,attr"`
	Right      int      `xml:"http://schemas.openxmlformats.org/2006/main right,attr"`
	LeftChars  int      `xml:"http://schemas.openxmlformats.org/2006/main leftChars,attr"`
	RightChars int      `xml:"http://schemas.openxmlformats.org/2006/main rightChars,attr"`
}

type Spacing struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main spacing,omitempty"`
	After   int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main after,attr"`
	Before  int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main before,attr"`
	Line    int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main line,attr"`
}

func getAtt(atts []xml.Attr, name string) string {
	for _, at := range atts {
		if at.Name.Local == name {
			return at.Value
		}
	}
	return ""
}
