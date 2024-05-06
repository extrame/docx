package docx

// AddParagraph adds a new paragraph
func (f *DocxFile) AddParagraph() *Paragraph {
	p := &Paragraph{
		// Data: make([]ParagraphChild, 0),
		file: f,
	}

	f.Document.Body.Paragraphs = append(f.Document.Body.Paragraphs, p)
	return p
}

func (f *DocxFile) Paragraphs() []*Paragraph {
	return f.
		Document.
		Body.
		Paragraphs
}

// func (p *Paragraph) Children() (ret []ParagraphChild) {
// 	return p.Data
// }
