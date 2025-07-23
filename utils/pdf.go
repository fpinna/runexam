package utils

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"net/url"
	"runexam/types"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// GeneratePDF - Generates a PDF from the result page data
func GeneratePDF(data types.ResultPage) []byte {
	// Render PDF Template
	tpl, err := template.New("pdf.html").Funcs(template.FuncMap{
		"add1": func(i int) int { return i + 1 },
	}).ParseFiles("templates/pdf.html")
	if err != nil {
		log.Println("Error parsing PDF template:", err)
		return nil
	}

	var htmlBuf bytes.Buffer
	if err := tpl.ExecuteTemplate(&htmlBuf, "pdf.html", data); err != nil {
		log.Println("Error rendering PDF template:", err)
		return nil
	}

	htmlStr := htmlBuf.String()

	// Usa chromedp para renderizar o HTML e exportar como PDF
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 15*time.Second)
	defer timeoutCancel()

	err = chromedp.Run(timeoutCtx,
		chromedp.Navigate("data:text/html,"+url.PathEscape(htmlStr)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return err
			}
			pdfBuf = buf
			return nil
		}),
	)
	if err != nil {
		log.Println("Error generating PDF:", err)
		return nil
	}
	return pdfBuf
}
