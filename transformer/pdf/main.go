package pdf

import (
	"io"
	"os"

	"github.com/extrame/docx"
	"github.com/jung-kurt/gofpdf"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font/sfnt"
)

func Trans(name string, doc *docx.DocxFile, dest string, fontFile ...string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetMargins(10, 10, 10)
	pdf.SetAutoPageBreak(true, 10)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(0.1)

	if len(fontFile) > 0 {

		_, err := os.Stat(fontFile[0])
		if err != nil {
			if fontFile[0] == "" {
				pdf.SetFont("Arial", "", 12)
				goto parsing
			}
			if os.IsNotExist(err) {
				//treat as font name
				logrus.WithField("font", fontFile[0]).Info("Set Font")
				pdf.SetFont(fontFile[0], "", 12)
				goto parsing
			}
			logrus.WithError(err).Error("Error getting font file info")
			return err
		}

		logrus.WithField("font", fontFile[0]).Info("Adding font")
		ttfFile := fontFile[0]
		file, err := os.Open(ttfFile)
		if err != nil {
			logrus.WithError(err).Error("Error opening font file")
			return err
		}
		defer file.Close()
		var bts []byte
		bts, err = io.ReadAll(file)
		if err != nil {
			logrus.WithError(err).Error("Error reading font file")
			return err
		}
		font, err := sfnt.Parse(bts)
		if err != nil {
			logrus.WithError(err).Error("Error parsing font file")
			return err
		}
		name, err := font.Name(nil, sfnt.NameIDFamily)

		if err != nil {
			logrus.WithError(err).Error("Error getting font name")
			return err
		}

		pdf.AddUTF8FontFromBytes(name, "", bts)
		logrus.WithField("font", name).Info("Font added")
		pdf.SetFont(name, "", 12)
	} else {
		pdf.SetFont("Arial", "", 12)
	}

parsing:

	pdf.AddPage()
	pdf.SetY(10)
	pdf.SetX(10)
	pdf.CellFormat(190, 10, name, "0", 0, "C", true, 0, "")
	pdf.Ln(10)
	pdf.SetY(20)
	pdf.SetX(10)

	for _, para := range doc.Paragraphs() {

		style := para.GetStyle()

		if style != nil {
			var headingLevel = style.HeadingLevel()

			if headingLevel > 0 {
				logrus.Debugf("Heading %d -> %s\n", headingLevel, para.Text())
				pdf.SetX(10)
				pdf.MultiCell(190, 10, para.Text(), "0", "L", true)
				pdf.Ln(10)
			}
		} else {
			logrus.Debug(para.Text())
			pdf.SetX(10)
			pdf.MultiCell(190, 10, para.Text(), "0", "L", true)
			pdf.Ln(10)
		}
	}

	err := pdf.OutputFileAndClose(dest)
	if err != nil {
		logrus.Error(err)
	}
	return err
}
