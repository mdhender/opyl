---
title: Add a new order file format
weight: 1
prev: /how-to
---

{{< callout type="info" >}}
Placeholder. This guide will document the steps to add a new flat-file
format that produces `domain.OrderBundle` values.
{{< /callout >}}

This guide is for developers extending opyl with a new order file format
(for example, a structured email body, or a new DSL).

## Before you start

You should already understand:

- How `app.OrderSource` is declared (see [reference](../../reference/ports))
- The SOUSA placement rule for infra adapters (see
  [explanation](../../explanation/sousa-in-opyl))

## Steps

1. Create a new package under `internal/infra/orderfile/<format>/`.
2. Implement `app.OrderSource` against that format.
3. Validate input shape and surface `cerr.ErrInvalidOrders` on failure.
4. Wire the adapter selection in `cmd/opyl/main.go` based on a CLI flag.
5. Add fixture files and a parser test under the same package.
6. Run `go test ./...` and the SOUSA conformance checks.

_(to be filled in once the first format is implemented)_
