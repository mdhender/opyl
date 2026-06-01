---
title: Map coordinate compression
weight: 4
prev: /explanation/use-cases
---

A province has two names. One is its **identity** — the numeric `(row, col)`
pair the engine stores and reasons about. The other is its **label** — the
short code a player sees in a report, like `aa1` or `ae48`. This page explains
why opyl keeps those two separate, and how the label is "compressed" into
something compact and unambiguous.

## Two layers, on purpose

The engine never identifies a province by its printed code. Internally a
province is a one-based `(row, col)` pair: `(1, 1)` is the top-left (northwest)
corner, rows increase to the south, columns increase to the east, out to
`(n, n)`. Arithmetic — neighbours, distances, edge checks — is just integer
math on that pair.

The code is **display formatting layered on top**. This split is the same
discipline the rest of opyl follows: the domain holds plain values, and turning
them into something human-facing is a rendering concern. It also means we can
change how codes *look* without touching a single rule, because no rule depends
on the spelling of a code.

So `[ae48]` in a turn report is cosmetic. The province *is* `(row 27, col 48)`.

## What "compressed" means

The rulebook opyl is derived from labelled provinces with a fixed, padded
`aa00` scheme — always two letters, always two digits, counting columns from
zero. opyl drops that for a **compressed** code: variable width, no leading
zeros, counting from one. The top-left province is `a1`, not `aa00`.

"Compressed" buys two things:

- **Short labels stay short.** Early rows and columns don't carry padding they
  don't need. `a1`, `h8`, `z1` are complete codes.
- **No ambiguous characters.** Codes are read aloud, retyped from email, and
  squinted at in monospaced reports, so the letters that look like digits are
  removed entirely (see below).

## The row alphabet

The row part of a code is spelled from a 22-letter alphabet — the Latin
alphabet with **`i`, `j`, `l`, and `o` removed**, because they are too easily
confused with `1` and `0`:

```
a b c d e f g h k m n p q r s t u v w x y z
```

This is a deliberate change from the rulebook's stated `abcdfghjkmnpqrstvwxz`
sequence, which dropped `e` yet used `[ae48]` as its running example — an
internal contradiction. Keeping `e` and dropping the look-alikes resolves it
and leaves the rulebook's worked examples valid.

The column part is a plain decimal number. Digits never collide with the row
letters, so the column needs no special alphabet.

## Bijective counting: why `z` is followed by `aa`

The row letters form a **bijective base-22** numeral. "Bijective" means there
is **no zero digit** — every position uses `1`–`22` (`a`–`z`), never `0`. That
is exactly what you want for a label: there is no awkward `a0`, and every row
has precisely one spelling.

Counting goes:

```
a  = 1      h  = 8       s  = 15
b  = 2      k  = 9       t  = 16
c  = 3      m  = 10      u  = 17
d  = 4      n  = 11      v  = 18
e  = 5      p  = 12      w  = 19
f  = 6      q  = 13      x  = 20
g  = 7      r  = 14      y  = 21
                         z  = 22
```

After `z` (22) there is no `z+1` digit, so the count rolls into a second
letter: **`aa` = 23**, `ab` = 24, and so on. Two-letter rows therefore run from
23 (`aa`) to 506 (`zz`); a third letter appears only past that (`aaa` = 507).
The scheme has no hard ceiling — very large maps just grow a letter.

This is the same counting a spreadsheet uses for columns (`Z` then `AA`), only
in base 22 instead of base 26.

## Worked examples

| Code   | `(row, col)`   | Notes                                          |
| ------ | -------------- | ---------------------------------------------- |
| `a1`   | `(1, 1)`       | Top-left corner of the whole coordinate space  |
| `h8`   | `(8, 8)`       | Single-letter row, single-digit column         |
| `z1`   | `(22, 1)`      | Last single-letter row                          |
| `aa1`  | `(23, 1)`      | First two-letter row, immediately after `z`     |
| `ae48` | `(27, 48)`     | `ae` = 1×22 + 5 = 27                             |
| `zz99` | `(506, 99)`    | Last two-letter row                             |

## Converting in both directions

Decoding a row string to a number walks the letters left to right,
multiplying by 22 each step:

```
row = 0
for each letter L in the row string:
    row = row * 22 + position(L)      # a=1 … z=22
```

`ae` → `(0×22 + 1)×22 + 5` = `27`.

Encoding a number back to letters peels off one digit at a time, with the `-1`
that makes the system bijective:

```
while n > 0:
    n, r = divmod(n - 1, 22)
    prepend alphabet[r] to the string    # alphabet is 0-indexed here
```

`27` → first step `divmod(26, 22)` = `(1, 4)` → `e`; second step
`divmod(0, 22)` = `(0, 1)`… handled as the leading `a` → `ae`.

The column is just the decimal number; no conversion needed.

## Where the map starts

The coordinate space always begins at `(1, 1)` = `a1`, but a game master is
free to lay out the playable world wherever inside it they like. By convention,
most maps put **`aa1` at the upper-left**, which is `(row 23, col 1)`. That
leaves the 22 single-letter rows `a`–`z` as empty margin to the north — handy
breathing room above the inhabited world.

None of this matters to the engine. It stores `(row, col)`, computes on
`(row, col)`, and only ever spells out a code when it is about to show one to a
person. The convention is for the people drawing maps, not for the simulation.

## See also

- [Geography & Movement]({{< relref "/rules/geography" >}}) — the player-facing
  rules where these codes appear.
- `GAME-DESIGN.md` §2.2 — the design decision this page expands on.
