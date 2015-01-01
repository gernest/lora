package models

import (
	"html/template"

	"time"
)

func (s *Section) Sanitize() {

	s.BodyHtml = template.HTML(sanitizeHTMLField(s.Body))
}

func (s *Section) LastUpdate() string {
	return s.UpdatedAt.Format(time.RFC822)
}
