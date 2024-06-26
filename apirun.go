package docx

// Color allows to set run color
func (r *Run) Color(color string) *Run {
	r.RunProperties.Color = &StrValueNode{
		Val: color,
	}

	return r
}

// Size allows to set run size
func (r *Run) Size(size int) *Run {
	r.RunProperties.Size = &IntValueNode{
		Val: int64(size * 2),
	}
	return r
}

// AddText adds text to paragraph
func (p *Paragraph) AddText(text string) *Run {
	t := &Text{
		Text: text,
	}

	run := &Run{
		Text:          t,
		RunProperties: &RunProperties{},
	}

	// p.Data = append(p.Data, ParagraphChild{Run: run})

	return run
}
