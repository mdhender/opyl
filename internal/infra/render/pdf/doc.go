// Package pdf is the infra adapter that renders a domain.PlayerReport
// as PDF bytes, implementing app.ReportRenderer.
//
// SOUSA: PDF library choice (gofpdf, typst, chromedp, etc.) is
// contained to this package. Swapping libraries must not require
// changes outside infra/render/pdf.
package pdf
