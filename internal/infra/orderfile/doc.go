// Package orderfile is the infra adapter that parses flat-file player
// orders into domain.OrderBundle values, implementing app.OrderSource.
//
// SOUSA: input files are untrusted. Validate shape here at the
// boundary; never let malformed bytes reach app or domain. Parse and
// classify errors into cerr sentinels before returning.
//
// This package owns whatever flat-file format(s) opyl accepts. It must
// not contain game-rule decisions, persistence, rendering, or email
// concerns.
package orderfile
