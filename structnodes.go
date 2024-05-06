package docx

import (
	"encoding/xml"

	"github.com/sirupsen/logrus"
)

type Paragraph struct {
	XMLName    xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main p"`
	Properties *ParagraphProperties
	Links      []*Hyperlink `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	Runs       []*Run       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main r,omitempty"`
	file       *DocxFile
}

func (p *Paragraph) GetStyle() *DefinedStyle {
	if p.Properties == nil {
		return nil
	}
	var styleId = p.Properties.GetStyleId()
	return p.
		file.
		Document.
		Styles.
		GetStyleById(styleId)
}

func (p *Paragraph) Text() string {
	var text string
	for _, r := range p.Runs {
		if r.Text == nil {
			continue
		}
		text += r.Text.Text
	}
	return text
}

func (p *Paragraph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var elem Paragraph
	var itemCount int
	for {
		t, err := d.Token()
		if err != nil {
			break
		}

		switch tt := t.(type) {
		case xml.StartElement:
			if tt.Name.Local == "pPr" {
				var value ParagraphProperties
				d.DecodeElement(&value, &tt)
				elem.Properties = &value
			} else if tt.Name.Local == "r" {
				var value Run
				err := d.DecodeElement(&value, &tt)
				value.no = itemCount
				itemCount++
				elem.Runs = append(elem.Runs, &value)
				logrus.Debug("Run: ", value, err)
			} else if tt.Name.Local == "hyperlink" {
				var value Hyperlink
				d.DecodeElement(&value, &tt)
				value.no = itemCount
				itemCount++
				elem.Links = append(elem.Links, &value)
			} else {
				continue
			}
		}
	}

	*p = elem

	return nil
}
