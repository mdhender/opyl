---
title: Playing the Game
weight: 1
toc: true
---

## Introduction

Olympia is an open-ended computer moderated fantasy simulation. Characters move, battle, explore and study in the Olympian world. Each week, players submit orders for their units. After the turn runs, Olympia sends reports to the players detailing what happened.

Olympia is set in a low technology fantasy world. Characters do not have fixed goals. They may study whatever skills they believe will be useful and pursue goals as they find them. Olympia has no victory conditions; no winner is ever declared.

An aspiring heroic fighter may purchase weapons, study combat and slay monsters. Business-minded individuals can establish a profitable trading empire. The study of magic could furnish tools to advance dark plans. An explorer could trade his gold to a shipbuilder for a galley, and set out on the high seas to find adventure and treasure.

Olympia turn reports tend to be long, detailed and dense. Since your player character may hire other units which can study, move and fight on their own, large factions of characters may be organized. Supplying orders for many units in a large faction can be a lot of work!

A turn report for a new player controlling three characters is about five (66 line) pages. Average turn reports tend to be 15-25 pages. The size of the report will depend on whether a player chooses to concentrate development in a few characters, or to pursue an empire- building strategy of acquiring many units and controlling as much territory as possible.

### Overview

Each player begins with control of one character and a sum of gold, the Olympian currency. Characters may hire subordinates, study skills, trade items, build things, and explore. Each turn a player submits orders for all of the characters they control. These include the initial character, known as the Player Character, and any vassals of the player character, which belong to the player's faction.

The Olympia program will execute orders for characters in parallel. Each order may take no time or some number of game days to complete. As many orders will be processed for each character as game time allows. Orders not completed by the end of the turn may continue into the next. Each turn covers one month of game time. An Olympian month lasts 30 game days.

### Turn schedule

Turns are run on Mondays, at 12:00 noon, Australian Eastern Standard Time (AEST) or Australian Eastern Daylight Time (AEDT), depending on whether daylight savings is in effect in Sydney, Australia.

Players should receive their turn reports as soon as email can be delivered. Late orders will queue for the next turn. Beware! Even if mailed before the deadline, your message may take some time to get to the order scanner. Send your orders in early to avoid grief. This will also give you time to correct mistakes reported by the order scanner.

New player additions are performed when the turn is run.

If for technical reasons the turn cannot be run at the scheduled time, it will be run as soon as the technical issues have been dealt with. If a long delay is expected, players will be notified by email.

### Olympia Times

The Olympia Times is published with each turn and mailed to each player in the game. The Times has three sections: any comments or notices from the GM (if there are any), signed press, and rumors.

Players should submit items to the Times with the press and rumor commands.

### Diplomatic email forwarding service

Players can send mail to the player controlling an entity by mailing to <olympia@shadowlandgames.com>.

For example, suppose you saw the following characters:

```report
Seen here:
  Circe [2225], with three peasants, accompanied by:
      Hephaestus [3446], with 11 peasants
```

Circe's lord isn't identified, so you don't know what player is controlling her. However, it is still possible to send Circe's owner a message. Mail to <olympia@shadowlandgames.com> and include the following:

```orders
#forwardto: 2225
  Welcome to my lands! Prepare to die!
```

The message will be forwarded to whatever player is controlling Circe [2225].

Forwarding works for character and player entities. Mail sent to other entity numbers, or to unaffiliated characters, will be silently discarded.

Forwarded mail is not anonymous. The original headers on your message will be preserved.

### Compensation for errors

Bugs are inevitable in large, complex programs such as the Olympia game engine. While every effort is made to ferret out bugs, players must expect that from time to time something will go wrong.

If you cannot tolerate encountering bugs, please do not play Olympia.

There are usually more bugs at the start of a game than after it has been running for a while.

Because of the nature of how Olympia processes turns, it is not possible to re-run a turn for a single player if something goes wrong. Turns may only be re-run for all players. Therefore, re-runs will be considered for only the most serious bug-provoked disasters which affect many players.

Generally the preferred compensation for being the victim of a bug is some grant of CLAIM gold or other items. It is difficult and error-prone to attempt to move items around the game database, and there are fairness issues for other players who may be unpleasantly surprised by database edits occuring between turns.

All compensation given for bugs is solely at the discretion of the GM. Effects for minor bugs are assumed to even out across most players over the course of the game, so compensation is usually granted only for bugs which have seriously impacted a player.

Please concisely document any bugs you find, providing short excerpts from turn reports to show what went wrong.

If you feel you require compensation, briefly state how your position was irrevokably harmed by the bug, and suggest a suitable compensation (preferably in the form of CLAIM items which may be provided).

Bug reports should be posted on the forum or emailed to <admin@shadowlandgames.com>.

### Cheating

1. A player may not control more than one faction in the game. Also, multiple players from the same account are not allowed. Each player must have their own unique email address. Exceptionally, a holiday replacement may be assigned in case you will be unavailable to play for a short period of time. Please send an e-mail to <admin@shadowlandgames.com> to indicate this.

2. Players should not send in orders for another player's faction, to ruin that person's turn or otherwise benefit.

3. Players must inform the GM of any game bugs found. Send an email to <admin@shadowlandgames.com> or post a bug report on the forum.

4. Anti-social behavior, including harassing telephone calls, sending obscene/obnoxious unwanted communications, mail bombing, etc. will not be tolerated.

Punishment for serious violations is generally banishment from Olympia. So don't cheat, do play fair, and be a good sport.

All decisions of the GM are final.

### Olympian Calendar

The Olympia calendar has two months for each season, for a total of eight months per Olympian year. Each month is 30 game days long.

| Season | Month | Name             |
| ------ | ----- | ---------------- |
| Spring | 1     | Fierce Winds     |
| Spring | 2     | Snowmelt         |
| Summer | 3     | Blossom bloom    |
| Summer | 4     | Sunsear          |
| Fall   | 5     | Thunder and rain |
| Fall   | 6     | Harvest          |
| Winter | 7     | Waning days      |
| Winter | 8     | Dark night       |

### Game additions and rule changes

Features may be added to Olympia every so often, such as new areas in the map, NPC races, magical artifacts, skills, spells, etc. It is also sometimes necessary to make slight alterations to existing game mechanics to correct bugs or flaws in game balance.

While it is inevitable that some players will be affected by rule changes, every effort is made to not disrupt existing game positions.

This is mentioned only as a warning that the rules may evolve over time. Changes are announced on the forum.

### How do I join?

Visit the Olympia web site at <http://www.shadowlandgames.com/olympia> use the online signup form.

The Olympia web site also includes back issues of The Olympia Times, articles about Olympia written by players, and other useful information.

### Definition of Terms

#### Entity, Unit:

_Nobles, items, locations, and skills_

Everything in the Olympian world has a unique code. referenced with an "entity number". The code is shown in brackets after the name. Some examples:

- a player: Rich Skrenta [501]
- a character: Osswid the Destroyer [5499]
- a skill: Shipcraft [600]
- a place: City of the Lost [gx14]
- an item: Gold [1]
- an item: Scroll [yq12]

#### Noble, Character

Used interchangeably. These are the individuals under the control of players. All player orders are given to characters.

Players start with one character. Others may be hired or persuaded to join the player's faction.

Characters may possess items, travel through locations, learn skills, engage in combat, cast magical spells, etc.

#### Faction

All of the units controlled by a player are called the player's faction. A player starts with only one character, but the faction may grow to have many units.

#### Player character

The player character, or PC, is the character the player starts with. The PC begins with a loyalty of oath-2. The PC may later FORM other characters. Nothing is special about the PC other than being the player's first character; if the PC is killed, play continues with the player's other characters.

#### Item, possession

Characters may hold items, such as gold, scrolls, weapons, magic potions, jewels, lumber, rugs, etc.

#### Men

Characters may also have non-descript men in their employ. These men are represented as possessions for simplicity. They include peasants, workers, sailors, and different kinds of soldiers.

These men may not learn skills, hold any items, or act independently from the noble they are with.

For example, one might see:

```report
Seen here:
  Law Netexus [2020], with three peasants
```

Law Netexus is a character; the three peasants are non-descript men accompanying him.

Characters obtain peasants with the **RECRUIT** order.

"Men" may also include beast-fighters such as dragons (see the Beastmastery skill), but does not include work-animals such as horses and oxen which have no combat values.

#### Skills

Characters may learn skills, which are used to perform tasks. For instance, _Sailing [601]_ must be known in order to sail a ship.

Skills are grouped into the following categories:

- Alchemy
- Beastmastery
- Combat
- Construction
- Forestry
- Mining
- Persuasion
- Shipcraft
- Stealth
- Trade

There are also six schools of magic. See the **STUDY** and **RESEARCH** commands for information about learning skills.

#### Noble Points

A player starts with a certain amount of Noble Points (NP's). Each player gets an additional NP at fixed turns that are a multiple of eight (so at turns 8, 16, 24, 32, etc...). Players who join the game late, get additional starting NP's, known as _Catch-up NP's_. Ideally, all players will have an equal number of NP's at their disposal at any time.

NP's are used to buy nobles with the **FORM** command. They are also required to learn some advanced skills, and to swear characters to oath loyalty.

#### Stack

A group of characters joined such that they move and fight together.

#### Province

A location on the map. Provinces may have sub-locations within them, such as cities, bogs, caves, etc. Provinces are either forest, swamp, mountain, desert, plains, or ocean.

#### Month

Each turn is a game month, or 30 game days.

#### Safe haven

New players start in a Safe Haven city. Combat or magic are not permitted in safe havens. New players may acclimate themselves in safety before venturing out into the world.

