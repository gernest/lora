package models

import "html/template"

func (s *SubSection) Sanitize() {
	s.BodyHtml = template.HTML(sanitizeHTMLField(s.Body))
}
