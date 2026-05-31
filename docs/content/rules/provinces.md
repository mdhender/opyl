---
title: Provinces & Relics
weight: 10
toc: true
---

## Province control and ownership

Each province and city generates an amount of gold each month. This gold is known as its _tax base_. The size of the tax base is determined by the civilization level of the province. This gold may be pillaged, taxed by a castle, or taxed by a garrison. It does not accumulate if left uncollected at the end of a turn.

A castle automatically collects all of the gold in its province.

The `GARRISON CASTLE` order installs a group of at least ten soldiers in a province to claim it and guard against pillaging. Garrisons must be bound to a castle.

A garrison pays maintenance for its members from the province tax base, then forwards 1/2 of the remaining gold to its castle.

Garrisons with fewer than 10 fighting men will pay maintenance for themselves, but will not be able to forward tax to a castle, guard against pillaging or obey decree orders.

The castle owner gains status from the number of provinces under control:

| provinces | rank                                          |
| --------- | --------------------------------------------- |
| 1-5       | lord                                          |
| 6-12      | knight                                        |
| 13-25     | baron                                         |
| 25-37     | count                                         |
| 38-50     | earl                                          |
| 51-63     | marquess                                      |
| 64+       | duke                                          |
| region    | king (region must have at least 15 provinces) |

A noble may `PLEDGE` to another noble, granting status and control of owned provinces. The status of a noble who pledges is the smaller of the original status or one below the rank of the pledge target:

```
new status = min(original status, one below rank of pledge target)
```

Control of a province allows one to change its name or the name of any of its sublocations, take items from the garrison, and issue decrees to watch for certain units, or to attack specified units on sight.

The castle continues to receive the income from garrisoned provinces, even if the castle's owner is pledged to another noble.

Every noble in the pledge chain shares control of the garrisoned provinces. In other words, a castle owner may pledge to a noble, who in turn may pledge to a third noble, etc. Thus a province may have any number of rulers.

Visitors to a province are informed of the castle to which the garrison is bound, and the top-most ruler in the pledge chain (which may simply be the owner of the castle):

```report
Province controlled by Amber Keep [0909], castle, in Forest
  [cj12]
Ruled by Erekosse [5210], baron
...
Seen here:
  Garrison [780], garrison, on guard, with ten soldiers
```

### Tax base

Each province generates a tax base each month. The amount of gold fed into the tax base is determined by the civilization level of the province:

| civ level  | tax gold |
| ---------- | -------- |
| wilderness | 50       |
| civ-1      | 100      |
| civ-2      | 150      |
| civ-3      | 200      |
| civ-4      | 250      |
| civ-5      | 300      |
| civ-6      | 350      |
| civ-7      | 400      |

Cities add a flat 100 gold/month to their province's tax base.

The tax base support garrisons, can be collected by castles, or seized through pillaging.

Pillaging and opium consumption reduce the future tax base of a province.

A city's tax base is added to the province's tax base at the end of the turn. If the city is pillaged during the month, the amount transferred to the province will be diminished.

Gold left in the province at the end of the turn does not accumulate.

### Castles

Castles are the foundation of land ownership. A castle provides its owner with taxes from the province it is located in, as well as from garrisons in other provinces which are bound to the castle.

The owner of a castle automatically receives half of the remaining tax base from the castle's province at the end of each month. If a garrison is stationed in the same province as a castle, the garrison will first pay maintenance from the province's tax base, then the castle will collect half of whatever is left.

Each province may contain only one castle. The castle must be built in the outer province or in a city, if the province has one. (Tax revenue for the castle is the same no matter where it is built.) Castles may not be built inside other sublocations.

A castle alone is not sufficient to rule a province. A garrison must be stationed outside the castle in the province to protect it.

### Garrisons

Garrisons are groups of men who are stationed in provinces to protect them, and collect taxes in the name of a castle. Garrisons must be created with the `GARRISON` order, and must be bound to a castle located in the same region.

For example, suppose that the region Lesser Atnos had 20 provinces. One of these provinces contains Amber Keep [0909]. A garrison bound to Amber Keep could be stationed in each of the 20 provinces (including the province containing the castle itself).

Continuing the example, garrison units not in the Lesser Atnos region could not be bound to Amber Keep. The castle a garrison is bound to must be in the same region.

Garrisons can be bound to any castle in the region. If Lesser Atnos had two castles, some of the garrisons could be bound to one, and the rest to the second castle.

### More about garrisons

A garrison in a province containing a castle must be bound to that castle.

A garrison may only be installed in a province adjoining a province which already contains a garrison bound to the same castle, or the province the castle is in.

Garrisons are established with the `GARRISON CASTLE` order. Ten soldiers are required to create a garrison. The `GARRISON` order must be issue at the outer level of a province; one can't establish a garrison while inside a city, building or other sublocation.

The garrison pays the maintenance cost of its men directly from the tax base of the province. Half of the remaining tax base is forwarded to the castle the garrison is bound to.

For example, a garrison of ten soldiers would require 20 gold per month to support. This would leave 280 gold remaining in a typical province. 50% of this, or 140 gold, would be forwarded to the owner of the garrison's castle.

Example:

```report
> garrison cy09
Installed Garrison [780], garrison, on guard, with ten soldiers
```

Visitors to this province would see:

```
Province controlled by Amber Keep [0909], castle, in Forest
  [cj12]
Ruled by Erekosse [5210], baron
```

Note that Erekosse may be located inside the castle, or the castle's owner may have pledged service to him, in which case Erekosse could be anywhere.

### Garrison reports

Garrisons do not provide full location reports to their owners. They do notice any resource depletion activity, such as timber cutting or mining, as well as any large or unusual parties which enter their province. This includes any stack of five units or more, any party of 20 or more men, and most monsters or wild beasts.

Garrisons do not monitor activity in hidden locations, even if the players who rule over the garrisons have discovered the hidden locations.

The `DECREE WATCH WHO` order may be given by a ruler to instruct all garrisons to watch for a particular unit. This is useful for locating individuals who would otherwise go unnoticed by the garrisons.

### Referring to garrisons

Since a province may only have one garrison, garrisons may be referred to without knowing their entity number. The keyword `GARRISON` will match the province's garrison, if there is one.

Examples:

```orders
give garrison 12 5
attack garrison
```

### Status

The number of provinces a noble controls determines his status or rank:

| provinces | rank                                          |
| --------- | --------------------------------------------- |
| 1-5       | lord                                          |
| 6-12      | knight                                        |
| 13-25     | baron                                         |
| 25-37     | count                                         |
| 38-50     | earl                                          |
| 51-63     | marquess                                      |
| 64+       | duke                                          |
| region    | king (region must have at least 15 provinces) |

In addition, if a character has control over every province in a region, and the region contains at least 15 provinces, then the character is given the rank of king.

Provinces may be directly owned, if the noble is the owner of a castle, or indirectly, through other pledged nobles.

### Pledging land

A noble may `PLEDGE` his lands to another noble. This grants the pledge target status by increasing the number provinces he may rule over.

For example, suppose there are two castle owners, Osswid and Feasel. Osswid has garrisoned six provinces, and Feasel has three. Osswid is therefore a baron, and Feasel is a lord.

If Feasel and Osswid both pledge to Candide, Candide would attain the rank of Count. Osswid and Feasel would remain at the same rank in this example.

Candide would receive garrison reports for all provinces which Osswid and Feasel control. He would have the same privileges in the controlled provinces: he could take items from the garrisons, alter the names of the provinces or their sublocations, and issue watch and hostile decrees.

However, the income generated by the provinces would continue to be forwarded to the castles. No extra income goes to the pledge target.

### Status after pledging

The status of a noble `A` who is pledged to another noble `B` will be either `A`'s original status, as determined by how many provinces he controls, or one rank below `B`, whichever is lower.

For example, a noble with 5 provinces who pledges to a king will remain a baron. However, if pledged to another baron, the noble's rank would fall to lord.

## Relics

Relics are unique item artifacts which are introduced into the game via quests with monsters. All relics except the Throne return to the netherworld after use or some delay to be given out to a new adventurer via QUEST.

### Imperial Throne [401]

Long ago, the emperor of Olympia sat on his throne in the emperor's palace, high amidst Mt. Olympus. It is said that whoever rebuilds the famed castle on Mt. Olympus, and sits upon the throne, will be titled Emperor of Olympia.

### Crown of Prosperity [402]

The Crown of Prosperity was once worn by the most prosperous mortal to ever rule in Olympia, King Damar. Damar now wears his crown in the underworld and reflects on his past adventures of long ago, before the first Great Ending swept away all that he knew, and carried him from his beloved city of Kircarth to the land of the dead.

Sometimes nobles are able to acquire the Crown and hold it for a time. The crown infuses whatever province it ends each turn in with a measure of prosperity and economic health, equivalent to a +2 increase in the province's civilization level.

However, this prosperity does not last forever, for King Damar's ghostly hand invariably will reach out from the underworld to reclaim his relic. The Crown can be expected to return to its rightful owner 12-24 turns after its appearance in the mortal world.

### Skull of Bastrestric [403]

The most feared of the ancients was a wizard of terrible power known as Bastrestric ther Archymonaged. Bastrestric routinely incinerated his foes (or anyone who offended him) with bursts of raw aura energy directed from his black tower in the castle built on Mt. Olympus.

Though he was the most powerful living mage in the known world, and by any measure a fearsome, unholy force, Bastrestric yearned for ever greater abilities. He felt increasingly constrained by the limits of his mortal body, and thus in time resolved to abandon it. BtA's spirit jumped free of his body, plunging directly into the aura rivers which bind together the deepest structures of the very world itself. Only a scorched, empty body was left behind in his tower.

BtA's spirit has passed on, but the remants of his mortal body continue to radiate intense power.

Use of BtA's skull (USE 403) causes an intense aura burst which a mage will attempt to absorb. If successful, the mage will gain a 50-75 boost to current aura (within the limit of 5 times the mage's maximum aura).

There is a 25% chance that the mage will be killed by use of the skull. Non-mages who use the skull will be instantly killed.

BtA's skull will vanish with use, or within 10-20 turns after appearing in the mortal world.

