// Package text is the infra adapter that renders a domain.PlayerReport
// as plain-text turn report bytes, implementing app.ReportRenderer.
//
// SOUSA: rendering is an outer-layer concern. This package may import
// domain for the report shape, but must not contain game rules or
// reach for orders/state directly.
package text
