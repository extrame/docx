package docx

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

const (
	XMLNS_W = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	XMLNS_R = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
)

type Body struct {
	XMLName    xml.Name     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Paragraphs []*Paragraph `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main p"`
}

type Document struct {
	XMLName xml.Name       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main document"`
	XMLW    string         `xml:"xmlns:w,attr"`
	XMLR    string         `xml:"xmlns:r,attr"`
	Body    *Body          `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Styles  *DocumentStyle `xml:"-"`
}

type DocumentStyle struct {
	XMLName         xml.Name         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main styles"`
	XMLW            string           `xml:"xmlns:w,attr"`
	XMLR            string           `xml:"xmlns:r,attr"`
	DocumentDefault *DocumentDefault `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main docDefaults"`
	LatentStyles    *LatentStyles    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main latentStyles"`
	Styles          []*DefinedStyle  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main style"`
	styleMap        map[string]*DefinedStyle
}

func (d *DocumentStyle) GetStyleById(styleId string) *DefinedStyle {
	if d.styleMap == nil {
		d.styleMap = make(map[string]*DefinedStyle)
	}
	if s, ok := d.styleMap[styleId]; ok {
		return s
	}
	for _, s := range d.Styles {
		d.styleMap[s.StyleId] = s
		if s.StyleId == styleId {
			return s
		}
	}
	return nil
}

type DocumentDefault struct {
	XMLName    xml.Name    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main docDefaults"`
	RPrDefault *RPrDefault `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPrDefault"`
	PPrDefault *PPrDefault `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pPrDefault"`
}

type RPrDefault struct {
	XMLName xml.Name       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPrDefault"`
	RPr     *RunProperties `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr"`
}

type PPrDefault struct {
	XMLName xml.Name       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pPrDefault"`
	PPr     *ParagraphProp `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pPr"`
}

type ParagraphProp struct {
	XMLName xml.Name      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pPr"`
	Style   *StrValueNode `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pStyle,omitempty"`
}

type LatentStyles struct {
	XMLName           xml.Name        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main latentStyles"`
	LsdExceptions     []*LsdException `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main lsdException"`
	Count             int             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main count,attr"`
	DefQFormat        int             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defQFormat,attr"`
	DefUnhideWhenUsed int             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defUnhideWhenUsed,attr"`
	DefUIPriority     int             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defUIPriority,attr"`
	DefLockedState    int             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defLockedState,attr"`
	DefSemiHidden     int             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defSemiHidden,attr"`
	DefPrimaryStyle   int             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defPrimaryStyle,attr"`
}

type LsdException struct {
	XMLName        xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main lsdException"`
	Name           string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main name,attr"`
	Locked         int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main locked,attr"`
	SemiHidden     int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main semiHidden,attr"`
	UnhideWhenUsed int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main unhideWhenUsed,attr"`
	QFormat        int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main qFormat,attr"`
	UIPriority     int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main uiPriority,attr"`
	PrimaryStyle   int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main primaryStyle,attr"`
}

type DefinedStyle struct {
	XMLName      xml.Name             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main style"`
	Type         string               `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
	StyleId      string               `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main styleId,attr"`
	Name         *StrValueNode        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main name"`
	BasedOn      *StrValueNode        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main basedOn"`
	Next         *StrValueNode        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main next"`
	Link         string               `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main link"`
	RPr          *RunProperties       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr"`
	PPr          *ParagraphProperties `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pPr"`
	AutoRedefine *StrValueNode        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main autoRedefine"`
	SemiHidden   *StrValueNode        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main semiHidden"`
	QFormat      *StrValueNode        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main qFormat"`
	UiPriority   *StrValueNode        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main uiPriority"`
	TblPr        *TblPr               `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblPr"`
}

func (d *DefinedStyle) GetName() string {
	if d.Name != nil {
		return d.Name.Val
	}
	return ""
}

func (d *DefinedStyle) HeadingLevel() int {
	if d.PPr == nil {
		return 0
	}
	if strings.HasPrefix(d.Name.Val, "heading ") {
		level := d.Name.Val[8:]
		if level != "" {
			iLevel, err := strconv.Atoi(level)
			if err == nil {
				return iLevel
			}
		}
	}
	return 0
}

type StrValueNode struct {
	XMLName xml.Name
	Val     string `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

func (v *StrValueNode) String() string {
	return fmt.Sprintf("%s:'%s'", v.XMLName.Local, v.Val)
}

type IntValueNode struct {
	XMLName xml.Name
	Val     int64 `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

type TblPr struct {
	XMLName    xml.Name    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblPr"`
	TblStyle   string      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblStyle,attr"`
	TblCellMar *TblCellMar `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblCellMar"`
}

type TblCellMar struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblCellMar"`
	Top     int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main top,attr"`
	Left    int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main left,attr"`
	Bottom  int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bottom,attr"`
	Right   int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main right,attr"`
}
