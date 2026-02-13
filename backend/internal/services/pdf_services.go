package services

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

func ReadPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var builder strings.Builder
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() || p.V.Key("Contents").Kind() == pdf.Null {
			continue
		}
		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			for _, word := range row.Content {
				builder.WriteString(word.S + " ")
			}
			builder.WriteString("\n")
		}
	}
	return builder.String(), nil
}
