package docx

// This contains internal functions needed to unpack (read) a zip file
import (
	"archive/zip"
	"encoding/xml"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// This receives a zip file (word documents are a zip with multiple xml inside)
// and parses the files that are relevant for us:
// 1.-Document
// 2.-Relationships
func unpack(zipReader *zip.Reader) (docx *DocxFile, err error) {
	var doc *Document
	var relations *Relationships
	for _, f := range zipReader.File {
		if f.Name == "word/_rels/document.xml.rels" {
			relations, err = processRelations(f)
			if err != nil {
				return nil, err
			}
		}
		if f.Name == "word/document.xml" {
			doc, err = processDoc(f)
			if err != nil {
				return nil, err
			}
		}
		if f.Name == "word/styles.xml" {
			styles, err := processStyles(f)
			if err != nil {
				return nil, err
			}
			doc.Styles = styles
		}
	}
	docx = &DocxFile{
		Document:    *doc,
		DocRelation: *relations,
	}
	for _, para := range doc.Body.Paragraphs {
		para.file = docx
	}
	return docx, nil
}

func processStyles(file *zip.File) (*DocumentStyle, error) {
	filebytes, err := readZipFile(file)
	if err != nil {
		logrus.Errorln("Error reading from internal zip file")
		return nil, err
	}
	logrus.Errorln("Doc:", string(filebytes))

	doc := DocumentStyle{
		XMLW:    XMLNS_W,
		XMLR:    XMLNS_R,
		XMLName: xml.Name{Space: XMLNS_W, Local: "styles"}}
	err = xml.Unmarshal(filebytes, &doc)
	if err != nil {
		logrus.Errorln("Error unmarshalling doc style", string(filebytes))
		return nil, err
	}
	return &doc, nil
}

// Processes one of the relevant files, the one with the actual document
func processDoc(file *zip.File) (*Document, error) {
	filebytes, err := readZipFile(file)
	if err != nil {
		logrus.Errorln("Error reading from internal zip file")
		return nil, err
	}
	logrus.Errorln("Doc:", string(filebytes))

	doc := Document{
		XMLW:    XMLNS_W,
		XMLR:    XMLNS_R,
		XMLName: xml.Name{Space: XMLNS_W, Local: "document"}}
	err = xml.Unmarshal(filebytes, &doc)
	if err != nil {
		logrus.Errorln("Error unmarshalling doc", string(filebytes))
		return nil, err
	}
	logrus.Debugln("Paragraph", doc.Body.Paragraphs)
	return &doc, nil
}

// Processes one of the relevant files, the one with the relationships
func processRelations(file *zip.File) (*Relationships, error) {
	filebytes, err := readZipFile(file)
	if err != nil {
		logrus.Errorln("Error reading from internal zip file")
		return nil, err
	}
	logrus.Errorln("Relations:", string(filebytes))

	rels := Relationships{Xmlns: XMLNS_R}
	err = xml.Unmarshal(filebytes, &rels)
	if err != nil {
		logrus.Errorln("Error unmarshalling relationships")
		return nil, err
	}
	return &rels, nil
}

// From a zip file structure, we return a byte array
func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
