// Package mail is the infra adapter that delivers a rendered
// domain.Attachment to a recipient, implementing app.ReportDispatcher.
//
// SOUSA: transport choice (SMTP, SES, SendGrid, or a "drop in
// /outbox" file dispatcher for dev) is contained here. The dispatcher
// does not know how the attachment was produced and does not decide
// who gets reports — that is an app concern.
package mail
