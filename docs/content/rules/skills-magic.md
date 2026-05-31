---
title: Skills, Magic & Religion
weight: 7
toc: true
---

## Skills

Skills represent knowledge that Olympian characters may know. Shipbuilding, thievery, kidnaping, training soldiers, castle building, mixing potions, and forging magical artifacts are just a few of the possible actions which skills allow a character to perform.

Skills are divided into category skills and sub-skills within those categories. The skill categories are:

| num | name         | time to learn          |
| --- | ------------ | ---------------------- |
| 600 | Shipcraft    | three weeks            |
| 610 | Combat       | three weeks            |
| 630 | Stealth      | four weeks             |
| 650 | Beastmastery | four weeks, 1 NP req'd |
| 670 | Persuasion   | four weeks             |
| 680 | Construction | three weeks            |
| 690 | Alchemy      | four weeks             |
| 700 | Forestry     | three weeks            |
| 720 | Mining       | three weeks            |
| 730 | Trade        | three weeks            |

There are also six schools of magic. For more details on learning and casting magical spells, see the _Magical Arts_ section.

The category skill must be learned before any of the sub-skills within the category may be known.

With each skill learned, the player will receive a lore sheet describing background information about the skill and how it may be used. Most skills are invoked with the `use skill` order. The skill lore sheets will give specific information about arguments to `use` and and requirements or limitations for using the skill.

The _lore sheets_ for the skill categories list some of the skills available for study within the category.

For instance, a noble wishing to undertake the study of Shipcraft would first learn the category skill with the `STUDY` command:

```orders
study 600 # study Shipcraft
```

The Shipcraft lore sheet lists some of the sub-skills available for building and sailing ships. One of these is Sailing [601], the skill required to control a ship on the ocean. The aspiring captain could then order:

```orders
study 601 # study Sailing
```

Once Sailing [601] is known, the lore sheet describing its usage will appear in the player's turn report.

To begin study of a skill requires the following:

- A source of instruction
- Payment of a fee for necessary materials
- Payment of Noble Points to learn advanced skills

### Sources of instruction

The source of instruction may be one of the following:

- The skill is commonly known, and may be studied anywhere once the category skill is known. When a category skill is learned, the player receives a lore sheet for the category in the turn report. The lore sheet lists the commonly known sub-skills which may be studied anywhere once the category skill is known.

Some sub-skills within a category may not be listed in the category lore sheet. These are not directly studyable, even once the parent skill is known. They must either be learned from a scroll, a rare book, or through research.

- The character is in a city which teaches the skill.

Many cities offer instruction in skills. Shipcraft [600], for example, is commonly taught in port cities.

The turn report lists skills taught by the location:

```report
Skills taught here:
    Alchemy [690]
```

- The skill was discovered through `RESEARCH`, and is listed as `Partially known` in the turn report.

```report
Partially known skills:
    Improve ship rigging [9999], 0/7
```

- The player has a book which teaches the skill. Items which teach skills are shown in the inventory listing:

```report
Inventory:
  qty  name
  ---  ----
    1  Old book [6001]

Old book [6001] permits study of the following skills:
  Alchemy [690]
```

### Fee to begin study

A fee of 100 gold is charged when study is first issued for a skill. The payment is used to acquire various materials needed for the study and eventual practice of the skill.

### Advanced skills

Some advanced skills require Noble Points to begin study. Each player begins the game with a number of NPs, and receives an additional one each year. Noble Points may be spent to acquire new nobles, or to learn advanced skills.

Most skills do not require Noble Points to learn. NPs are required for some heroic combat skills and for advanced magical spells.

### Study limit

Characters may STUDY up to 14 days per turn. Fast study days dont count towards this limit.

### Fast study

All new players are given 200+ "fast study" points. Each fast study point may be applied to a skill being studied in lieu of actually spending a day studying.

For example, the order `study 600 7` would apply 7 fast study days to learning Shipcraft. This study order would take 0 days to execute.

### Skill Experience

Experience is counted for each turn that a skill use is successfully completed. If a skill is used more than once per turn, only the first success will count towards experience.

Projects which take multiple turns to complete, such as shipbuilding or castle construction, only count towards experience when the project is finished.

| use   | level        |
| ----- | ------------ |
| 0-4   | apprentice   |
| 5-11  | journeyman   |
| 12-20 | adept        |
| 21-34 | master       |
| 35+   | grand master |

Experience will speed work with some skills. For example, a master shipbuilder will be able to construct a galley somewhat faster than an apprentice. Some skill uses benefit more from experience than others.

### Skills not rated for experience

A few skills are not rated for experience. These may be skills which are not directly used, or ones for which experience has no meaning. If a skill is not rated for experience, the skill level will not be shown in the skill listing.

For instance:

```report
Skills known:
  Shipcraft [600]
    Sailing [601], apprentice
    Shipbuilding [602], apprentice
    Fishing [603], apprentice
  Combat [610]
    Survive fatal wound [611]
    Fight to the death [612]
```

Since experience is not applicable to Survive fatal wound [611] and Fight to the death [612], its levels are not shown.

### Summary of studying

- The category skill must be learned before a sub-skill within the category may be studied.

- Beginning study of any skill costs 100 gold.

- Payment of Noble Points may be required to begin learning some advanced skills.

- A source of instruction must be available the first time `STUDY` is issued:
  - The skill is taught by the location.
  - Characters may `STUDY` up to 14 days per month.
  - The skill is commonly known, so it may be studied anywhere once its category skill is known.
  - The skill was discovered through `RESEARCH`, so it may now be studied.
  - The skill is taught by a book or a scroll.

- Once a skill is known, further `STUDY` of that skill has no effect.

### Teaching

Direct character-to-character teaching is not possible in Olympia. However, it is said that some magicians and alchemists possess the ability to record knowledge on scrolls, which other characters may study from.

### An example of study

This is a rough, heavily edited example of how some study commands might look in a turn report. All of the other details that a real turn report would have were omitted to focus on the study orders.

Turn one:

```report
Skills taught here:
    Shipcraft [600]

  1: > study 600
    1: Paid 100 gold to begin study.
    1: Will study Shipcraft for seven days.
    7: Learned Shipcraft [600].

Lore for Shipcraft [600]
------------------------
All skills concerning ocean travel fall under this category.
Shipcraft encompasses the building and repair of ships,
training of sailors, and navigation at sea.

The following skills may be studied directly once Shipcraft
is known:

    num   skill                              time to learn
    ---   -----                              -------------
    601   Sailing                            one week
```

Further skills may be found through research.

Turn two:

```report
1: > study 601
  1: Paid 100 gold to begin study.
  1: Will study Sailing for seven days.
  7: Learned Sailing [601].
```

(It doesn't matter where 601 is studied, since 600 offers it directly.)

Turn three:

```report
1: > study 690
  1: Instruction in [690] is not available here.
```

(Need to find a location or book that offers instruction in Alchemy)

Turn four:

```report
1: enter xxxx
  1: Arrival at City of Alchemists [xxxx]

Skills taught here:
    Alchemy [690]

  1: > study 690
  1: Paid 100 gold to begin study.
  1: Will study Alchemy for seven days.

Partially known skills:
    690  Alchemy, 7/14
```

(Alchemy requires 14 days of study to learn, seven of which we have completed.)

Turn five:

```report
1: > study 690
  1: Continue studying Alchemy.
  7: Learned Alchemy [690].

Skills known:
  600  Shipcraft
    601  Sailing, apprentice
  690  Alchemy
```

## Research

Research attempts to discover sub-skills which are not commonly known or made available when the category skill is learned.

Research is mostly used to discover new magical spells, as few spells are granted when a magic school is learned. However, even common skills such as Shipcraft, Combat and Construction may have hidden sub-skills which can be found through research.

Research for all skill categories except Religion [750] must be performed in a tower, by the tower's owner (the first character inside the tower). Towers make good laboratories for scholarly investigations, and minimize distractions. Other occupants of the tower may not use `RESEARCH`.

Research into Religion [750] must be performed in a temple, by the temple's owner.

Research by mages of 6th black circle level and above (maximum aura of 31 or higher) must be done in provinces with a civilization level of 1 or less.

When a category skill is learned, its lore sheet appears in the player's turn report, listing information about the skill as well as sub-skills which may be studied directly based on the parent skill.

An example:

```report
Lore for Shipcraft [600]
------------------------
All skills concerning ocean travel fall under this category.
Shipcraft encompasses the building and repair of ships,
training of sailors, and navigation at sea.

The following skills may be studied directly once Shipcraft
is known:

    num   skill                              time to learn
    ---   -----                              -------------
    601   Sailing                            one week
    ...

Further skills may be found through research.
```

The last line (`Further skills ...`) indicates that there are sub-skills of Shipcraft which are not mentioned in the lore sheet. These hidden skills may represent rare or hard-to-learn knowledge, or perhaps technology which has not yet been discovered.

Study of these skills is not possible based simply on knowledge of the parent skill. Sailing [601] can be learned by a character, no matter where he is, once Shipcraft is known. Hidden Shipcraft skills, however, must be learned in other ways, even if the character has learned their entity numbers from other players.

There are two possible ways such hidden skills may be learned:

1.  Through a rare book or scroll which offers instruction in the skill.
2.  Through research.

Research is the more difficult choice. However, to learn rare sub-skills there may be no alternative. Perhaps there are no players who already know the rare sub-skill, and so cannot record scrolls to instruct others. Or if there are, they choose to keep their knowledge secret.

Research incurs a fee of 25 gold to pay for miscellaneous materials and costs.

Each week of research yields a 25% chance that a new skill will be discovered.

If research is successful, the new sub-skill will be added to the character's partially known skill list. It then must be studied in order to become fully known and usable.

For example:

```research
1: > research 600
  1: Will research Shipcraft for seven days.
  7: Research uncovers a new skill:  Improve ship rigging [9999].
  7: To begin learning this skill, order 'study 9999'.

Partially known skills:
  9999  Improve ship rigging, 0/7
```

The `0/7` qualifying the new sub-skill in the partially known skill list indicates that seven full days of study are required to learn the skill, and none of them have yet been completed.

```research
1: > study 9999
  1: Paid 100 gold to begin study.
  1: Will study Improve ship rigging for seven days.
  7: Learned Improve ship rigging [9999].
```

The `RESEARCH` order is no longer used on the new skill once it becomes partially known. Research may continue to be used on Shipcraft [600], however, to seek out more hidden sub-skills.

Category skills may not be learned through research.

## Magic

```text
Magic is a dead cat in an oil-stained burlap bag.
Magic is a smelly old man, despised but feared by his neighbors.
Magic is what the king turns to, when his soldiers fail.
  --- a long-dead wise man of Areth Pirn
```

Magic is the dark art by which events are influenced outside of the normal boundaries of cause-and-effect. Rather than the glamorous ideal of shining wizards casting powerful fireballs at wicked foes, the reality of magic instead tends to be base, tedious work which earns few friends.

Hated and feared, the magician pursues his craft out of the sight of men. Like wisps of smoke rising from an ember cast into dry straw, so the mage's spells slowly take hold, woven with secret knowledge and foul ingredients.

The casting of a magical spell is accomplished with three ingredients: Knowledge of the spell, possession of any items necessary to fuel the spell, and a sufficient level of magical aura to perform the ritual or ceremony.

### Aura

Aura is a mystical force necessary to cast spells. Powerful mages will have a high aura rating, while apprentice sorcerers may only command a few points worth of aura.

Characters are rated for their current aura level and their maximum aura level. With each magical spell learned, a character will gain one point of maximum and current aura. Current aura is depleted by casting spells, and is naturally replenished at a rate of two points per turn. Current aura will increase until it reaches the maximum aura level. Other ways of gaining current aura may be found as the mage researches his craft.

Spells are rated on the amount of aura which the casting mage must possess. Minor spells may demand only one point of aura to cast. Powerful spells may require a current aura level of ten or higher.

### Magician status

Characters receive a rating in the turn reports based on their magical abiliity:

| maximum aura | label    |
| ------------ | -------- |
| 6-10         | conjurer |
| 11-15        | mage     |
| 16-20        | wizard   |
| 21-30        | sorcerer |
| 30+          | ??       |

For example:

```report
Osswid the Brave [5639], wizard, with three workers
```

The Basic magic [800] spell Appear common [803] allows magicians to prevent this label from displaying.

### Required Items

Many spells will require the magician to possess a rare or obscure item in order for the cast to succeed. Many of these items exist, of interest chiefly to sorcerers. Roots from plants found only in dense forests, bat's wings, and a dark blue powder which produces a brilliant cobalt flame when burned are only a few. Usually the required item is consumed by the attempt to cast the spell.

### Study of Spells

Magic is divided into six schools of study:

| num | name                  | time to learn          |
| --- | --------------------- | ---------------------- |
| 800 | Magic                 | four weeks, 1 NP req'd |
| 820 | Weather magic         | five weeks, 1 NP req'd |
| 840 | Scrying               | five weeks, 1 NP req'd |
| 860 | Gatecraft             | five weeks, 1 NP req'd |
| 880 | Artifact construction | six weeks, 1 NP req'd  |
| 900 | Necromancy            | six weeks, 2 NP req'd  |

Magical spells are simply sub-skills of one of the magical skill categories.

An aspiring mage will issue the `STUDY` order to learn the basics of a particular school of magic, known as the category skill. For example, one wishing to pursue knowledge of Magic [800] would order:

```orders
study 800
```

Once the category skill has been learned, the mage will receive a lore sheet listing some of the known spells of that school. The magician may then attempt to learn these spells through study.

A character must know the category skill for a school of magic before a spell in that school can be known.

Only some of the spells in the each school are commonly known. The more rare, obscure or powerful spells will need to be discovered via `RESEARCH` or by finding magic scrolls describing them.

Knowledge of individual spells in a school of magic is not possible without having learned the category skill.

### Schools of Magic

#### Magic [800]

The most common and well-known of the magical schools, Magic nonetheless has many useful and powerful spells.

Since most cities offer instruction in Magic, and several useful spells may be learned quickly, apprentice mages often begin their studies here.

#### Weather magic [820]

The study of spells to control the elements. Advanced weather magicians are said to research forgotten elf-lore in the hopes of finding tools to battle evil.

Weather magic is taught in Nimbus, Stratos and Aerovia, the three cities of the Cloudlands and randomly in non safe haven cities.

#### Scrying [840]

Scrying is the study of magical far-seeing, the ability to view images or learn information at a distance.

Scrying is taught in the Faery Cities, and randomly in non safe haven cities.

#### Gatecraft [860]

The study of ancient portals of teleportation which long ago connected distant regions of the land.

Knowledge of this art is taught in all safe haven cities, and in random non safe haven cities.

#### Artifact construction [880]

The realm of very advanced sorcerers, artifact construction concerns the making of physical items to focus, amplify or otherwise enhance a mage's power. Most spells in this school require the magician to possess high levels of aura.

Knowledge of this art is taught in random non safe haven cities and in some cities in Hades.

#### Necromancy [900]

The darkest of the Dark Arts, necromancy involves trafficking with undead or demon spirits to gain knowledge. Hated by civilized men, feared by other magicians, the necromancer seeks power and domination over the physical world.

Necromancy is taught in the City of the Dead, in Hades and randomly in non safe haven cities.

Monsters may sometimes be found guarding ancient books which teach the rare magical skill categories.

## Religion

Learning the skill category Religion [750] labels the character as a priest. Religion [750] requires 1 Noble Point and five weeks to learn. Religion [750] may be studied in any temple. Religion is not taught by cities. Skills in the Religion category are known as prayers.

Temples yield 100 gold in offerings per month to their owner, if the owner is a priest. Temples may be built anywhere, except inside another building.

Research into Religion [750] must be performed in a temple rather than a tower.

## Battles

Combat in Olympia occurs when one stack attacks another stack.

A unit will be defended by any other characters it is stacked with. Thus, a unit which is part of a stack cannot be attacked alone. No matter which unit is specified in the `ATTACK` order, the entire stack containing the unit will respond.

Only the leader of a stack may initiate combat. If another unit in a stack wants to issue `ATTACK`, it must first drop out of the stack.

Combat involves nobles, who possess strong heroic fighting abilities, and fighters who may accompany them. Fighters include soldiers, pikemen, swordsmen, knights, elite guard, crossbowmen, archers, and elite archers.

