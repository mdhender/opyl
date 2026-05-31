# TODO — opyl game design

Open questions and provisional decisions from [GAME-DESIGN.md](GAME-DESIGN.md), captured to
resume in a new session.

**Where we are:** §1 concept, §2 map, §3 entities, §4 `move` (accepted), and §5 turn
phases + simultaneity are decided. The design has a complete spine — concept, map,
entities, one full order, and the turn loop. Legend: 🟡 provisional / needs tuning ·
❓ undecided.

---

## Start here (priority order)

1. **§5 intra-phase determinism** ❓ — the canonical tie-break for simultaneous conflicts
   (two forces want one hex; multi-party combat order): by character id, or a seeded RNG
   recorded in the turn ledger? This *is* "resolve turns deterministically" and **blocks
   any combat-bearing verb**.
2. **§5 Heal & Combat rules** ❓ — what heals and by how much (GURPS HT-based?); what
   *triggers* combat (co-location + hostility?) and how it resolves.
3. **§4 next verbs** ❓ — spec one at a time: recruit, build, attack, hire (NPC), scout,
   trade, train. (`attack` needs determinism + combat rules first.)
4. **Economy model** ❓ — resource production/consumption underpinning Economy (phase 3),
   growth points, recruiting costs, and monster raids. Several systems already lean on it.

---

## Backlog by section

### §1 Concept
- 🟡 **Scoring** — exists but deferred; design later.

### §2 Map
- 🟡 **Strawman numbers** — MP budgets (walk 12 / ride 24 / fly 36) and default terrain
  entry costs; tune.
- ❓ **Terrain effects** — beyond movement: resource yield, defence, visibility/sighting?
  Which water types are passable by which travel modes?
- ❓ **Resource model** — are resources quantities held by the hex? Do they regenerate?
  How harvested into a nation's economy? (Overlaps the Economy item above.)
- ❓ **Structure attributes** — what settlement / castle / tower / ruins carry (owner,
  garrison, growth thresholds); what hides ruins and how they're discovered.

### §3 Entities
- ❓ **Nation** — how is "territory controlled" determined (occupied / adjacent hexes)?
- ❓ **Character** — cap on character count?
- ❓ **NPC** — hiring mechanics (cost, duration, loyalty); does an NPC belong to a nation
  or stand alone?
- ❓🟡 **Monster** — monster-agent order-generation logic (TBD); exact attraction /
  aggression model and raid triggers.
- 🟡 **Skills** — full skill list, levels, and which orders each gates/modifies.
- 🟡 **Inventory** — other item kinds beyond minions and money (equipment, carried
  resources).
- 🟡 **GURPS dependency** — write an *explanation* page on how much GURPS is adopted vs.
  simplified, plus a **licensing** note. (GURPS now load-bearing in 3 places: attributes,
  skills, City Stats growth.)

### §4 Orders
- ❓ **Order file format** — custom DSL vs YAML vs structured email (also an AGENTS.md
  decision). Provisional line form `<character>: <verb> <args>`.
- ❓ **Full catalog** — see "next verbs" above.

### §5 Turn resolution
- ❓ Intra-phase determinism — see priority #1.
- ❓ Heal & Combat rules — see priority #2.
- ❓ Economy phase depends on the resource/economy model.

### §6 Player report
- ❓ **Report detail** — beyond character sightings + minion reports: character/minion
  status, order outcomes (success/failure), economy, events/messages.
- ❓ **Sighting** — what "see" covers (current hex + adjacent? range by terrain); how
  stale knowledge of previously-seen hexes is presented. (Referenced by `move` and §6.)

### §7 Carried from AGENTS.md
- ❓ **Order file format** (= §4 above).
- ❓ **State storage** — SQLite vs per-turn JSON. Data point: the sparse `(q,r,direction)`
  movement-override table favours a keyed table / sparse map.
- ❓ **Concurrency** — turns serial per game, multiple games in parallel; confirm.

---

## Then: first reference page

Once a slice freezes and is reflected in `internal/domain`, write the first **reference**
page (the *Game entities* roster is closest to frozen). The GURPS note becomes an
**explanation** page. Per Diataxis, don't write reference ahead of the implementation.
