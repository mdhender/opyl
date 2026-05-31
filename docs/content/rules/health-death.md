---
title: Health & Death
weight: 4
toc: true
---

## Health

Nobles have a health rating of 1-100, which indicates how wounded they are. A noble with health 100 has no wounds; a noble with a health of zero dies. Nobles also have a flag which indicates whether they are suffering from an illness. A sick noble will lose some health each week, while a wounded noble who is otherwise free of illness will recover somewhat each week.

Health is shown in the turn report for each noble that the player controls.

This is a noble in perfect health:

```report
Health: 100%
```

This is a rather sick noble:

```report
Health: 38% (getting worse)
```

If the noble were cured of illness, this would instead show:

```report
Health: 38% (getting better)
```

Medical technology is rather crude in the age of Olympia. Sanitation and hygiene are not the best. Even a minor wound runs a risk of developing into a serious, possibly life-threatening problem.

When a noble receives a new injury, their health is reduced by the amount of the injury, and a check is made to see whether they get sick. The chance that a character falls ill is _(100 - health)_. Thus, the more seriously the noble is wounded, the greater the probability that infection will set in.

Health is updated at the end of each game week (on days 7, 14, 21 and 28). Sick nobles lose 3-15 health each week. Healthy but wounded nobles recover by a like amount. Each week sick nobles have a 5% chance of fighting off their illness.

Nobles located in an inn benefit from the rest and relaxation that the inn provides. Inns increase the chance of fighting off infection to 10% each week.

Illness can also be cured by spells, special skills, or magical potions. Sick nobles would be wise to seek out those practiced in healing, and recover in a nearby inn.

Notes:

- Characters who recover from illness will not suddenly become sick again. The illness check is only made when a new wound is received.
- Since the chance that a character becomes sick is based on the unit's health after deducting a hit, an already-wounded noble stands a greater chance of becoming ill from a wound than a healthy one.
- Health is only tracked for nobles. Peasants, workers, sailors, soldiers, etc. are either dead or alive; they have no health rating.
- A wound received in combat will be randomly chosen between 1 and 100. Wounds received in other activities are scaled to match the level of danger involved.
- Some NPCs are not rated for health. Health for these characters appears as n/a in the turn report. For means of directly inflicting damage against a character such as terrorize, Aura blasts, and Lightning bolts, such a unit requires a hit of 50 or more points in order to be killed.

## Death

When a noble is killed in battle or dies, the body is moved into the province as an item which may be found with **EXPLORE**. An executed noble's body is given to the character who performed the execution.

A dead body exists for one and one half game years, at which point it fully decomposes, and the dead noble's spirit passes on.

Example: A dead body rots after twelve months have passed, so if a noble dies in turn 20, his body will decompose at the end of turn 32.

The bodies of nobles lost at sea will wash ashore somewhere.

Bodies decompose after one and one half years regardless of whether they remain in the province or are possesions of another.

Priests may learn a skill **LAY TO REST**, hastening the passing of the dead noble's spirit. Some exceptionally skilled priests possess the ability to resurrect dead characters.

In general, NP's invested in characters are returned to the character's player if the character deserts or his body decomposes.

Characters which renounce service becaused of lapsed contract or fear loyalty do not immediately return the NPs to the previous owner. The previous owner of these characters will not receive NPs for them until they swear to a new player faction or die.

When a body decomposes, the number of NP's which were invested in the character are returned to the original owner.

