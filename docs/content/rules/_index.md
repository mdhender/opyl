---
title: Olympia Rules
weight: 5
toc: true
---

These are the rules of the Olympian world that opyl is built to simulate. They
describe the game itself, not the engine. For engine documentation, see
[Reference](../reference) and [How-to](../how-to).


## Introduction

Olympia is an open-ended computer moderated fantasy simulation.  Characters move, battle, explore and study in the Olympian world. Each week, players submit orders for their units. After the turn runs, Olympia sends reports to the players detailing what happened.

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

If for technical reasons the turn cannot be run at the scheduled time, it will be run as soon as the technical issues have been dealt with.  If a long delay is expected, players will be notified by email.

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

1. A player may not control more than one faction in the game. Also, multiple players from the same account are not allowed. Each player must have their own unique email address.  Exceptionally, a holiday replacement may be assigned in case you will be unavailable to play for a short period of time.  Please send an e-mail to <admin@shadowlandgames.com> to indicate this.

2. Players should not send in orders for another player's faction, to ruin that person's turn or otherwise benefit.

3. Players must inform the GM of any game bugs found.  Send an email to <admin@shadowlandgames.com> or post a bug report on the forum.

4. Anti-social behavior, including harassing telephone calls, sending obscene/obnoxious unwanted communications, mail bombing, etc. will not be tolerated.

Punishment for serious violations is generally banishment from Olympia. So don't cheat, do play fair, and be a good sport.

All decisions of the GM are final.

### Olympian Calendar

The Olympia calendar has two months for each season, for a total of eight months per Olympian year. Each month is 30 game days long. 

| Season | Month | Name             |
|--------|-------|------------------|
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

- Entity, Unit:

Nobles, items, locations, and skills

Everything in the Olympian world has a unique code. referenced with an "entity number". The code is shown in brackets after the name. Some examples:

- a player: Rich Skrenta [501]
- a character: Osswid the Destroyer [5499]
- a skill: Shipcraft [600]
- a place: City of the Lost [gx14]
- an item: Gold [1]
- an item: Scroll [yq12]

<!-- -->

- Noble, Character

Used interchangeably. These are the individuals under the control of players. All player orders are given to characters.

Players start with one character. Others may be hired or persuaded to join the player's faction.

Characters may possess items, travel through locations, learn skills, engage in combat, cast magical spells, etc.

- Faction

All of the units controlled by a player are called the player's faction. A player starts with only one character, but the faction may grow to have many units.

- Player character

The player character, or PC, is the character the player starts with. The PC begins with a loyalty of oath-2. The PC may later FORM other characters. Nothing is special about the PC other than being the player's first character; if the PC is killed, play continues with the player's other characters.

- Item, possession

Characters may hold items, such as gold, scrolls, weapons, magic potions, jewels, lumber, rugs, etc.

- Men

Characters may also have non-descript men in their employ. These men are represented as possessions for simplicity. They include peasants, workers, sailors, and different kinds of soldiers.

These men may not learn skills, hold any items, or act independently from the noble they are with.

For example, one might see:

**Seen here:**

**Law Netexus [2020], with three peasants**

Law Netexus is a character; the three peasants are non-descript men accompanying him.

Characters obtain peasants with the RECRUIT order.

"Men" may also include beast-fighters such as dragons (see the Beastmastery skill), but does not include work-animals such as horses and oxen which have no combat values.

- Skills

Characters may learn skills, which are used to perform tasks. For instance, *Sailing [601]* must be known in order to sail a ship.

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

There are also six schools of magic. See the STUDY and RESEARCH commands for information about learning skills.

- Noble Points

A player starts with a certain amount of Noble Points (NP's). Each player gets an additional NP at fixed turns that are a multiple of eight (so at turns 8, 16, 24, 32, etc...). Players who join the game late, get additional starting NP's, known as Catch-up NP's.  Ideally, all players will have an equal number of NP's at their disposal at any time.  

NP's are used to buy nobles with the FORM command. They are also required to learn some advanced skills, and to swear characters to oath loyalty.

- Stack

A group of characters joined such that they move and fight together.

- Province

A location on the map. Provinces may have sub-locations within them, such as cities, bogs, caves, etc. Provinces are either forest, swamp, mountain, desert, plains, or ocean.

- Month

Each turn is a game month, or 30 game days.

- Safe haven

New players start in a Safe Haven city. Combat or magic are not permitted in safe havens. New players may acclimate themselves in safety before venturing out into the world.

### Command summary

| Command    | Arguments                                 | Time     | Priority |
|------------|-------------------------------------------|----------|----------|
| ACCEPT     | \<from-who\> \<item\> [qty]             | 0 days   | 0        |
| ADMIT      | \<who or what\> [ALL] [units]         | 0 days   | 0        |
| ATTACK     | \<target\> [flag]                       | 1 day    | 3        |
| BANNER     | [unit] "message"                        | 0 days   | 1        |
| BEHIND     | \<number\>                                | 0 days   | 1        |
| BOARD      | \<ship\> [maximum fee]                  | 0 days   | 2        |
| BRIBE      | \<who\> \<amount\> [flag]               | 7 days   | 3        |
| BUILD      | \<structure\> "Name" [max days] [id]  | varies   | 3        |
| BUY        | \<item\> \<qty\> \<price\> [have-left]  | 0 days   | 1        |
| CATCH      | [number of horses] [days]             | as given | 3        |
| CLAIM      | \<item\> [number]                       | 0 days   | 1        |
| COLLECT    | \<item\> [number] [days]              | as given | 3        |
| CONTACT    | \<who\>                                   | 0 days   | 0        |
| DECREE     | \<decree\>                                | 0 days   | 0        |
| DEFAULT    | \<who\>                                   | 0 days   | 0        |
| DEFEND     | \<who\>                                   | 0 days   | 0        |
| DIE        |                                           | 0 days   | 1        |
| DROP       | \<item\> \<qty\> [have-left]            | 0 days   | 1        |
| EMAIL      | \<new email address\>                     |          |          |
| EXECUTE    | [prisoner]                              | 0 days   | 1        |
| EXHUME     | [body]                                  | 7 days   | 3        |
| EXPLORE    |                                           | 7 days   | 3        |
| FEE        | [gold per 100 wt]                       | 0 days   | 1        |
| FERRY      |                                           | 0 days   | 1        |
| FISH       | [number of fish] [days]               | as given | 3        |
| FLAG       | string                                    | 0 days   | 1        |
| FLY        | \<direction or destination\> [...]      | varies   | 2        |
| FORGET     | \<skill\>                                 | 0 days   | 1        |
| FORM       | \<unit\> "Name of new character"          | 7 days   | 3        |
| FORMAT     | \<number\>                                |          |          |
| GARRISON   | \<castle\>                                | 1 day    | 3        |
| GET        | \<who\> \<item\> [qty] [have-left]    | 0 days   | 1        |
| GIVE       | \<to-who\> \<item\> [qty] [have-left] | 0 days   | 1        |
| GUARD      | \<flag\>                                  | 0 days   | 1        |
| HONOR      | \<amount\>                                | 1 day    | 3        |
| HOSTILE    | \<who\>                                   | 0 days   | 0        |
| IMPROVE    | [days]                                  | varies   | 3        |
| LORE       | \<lore sheet\>                            |          |          |
| MAKE       | \<item\> [qty]                          | as given | 3        |
| MESSAGE    | \<# of lines of text\> \<to-who\>         | 1 day    | 3        |
| MOVE       | \<direction or destination\> [...]      | varies   | 2        |
| NAME       | [unit] "new name for unit"              | 0 days   | 1        |
| NEUTRAL    | \<who\>                                   | 0 days   | 0        |
| NOTAB      | \<number\>                                |          |          |
| OATH       | \<level\>                                 | 1 day    | 3        |
| PASSWORD   | ["password"]                            |          |          |
| PAY        | \<to-who\> [amount] [have-left]       | 0 days   | 1        |
| PILLAGE    | [flag]                                  | 7 days   | 3        |
| PLAYERS    |                                           |          |          |
| PLEDGE     | \<who\>                                   | 0 days   | 1        |
| POST       | \<# of lines of following text\>          | 1 day    | 3        |
| PRESS      | \<# of lines of text\>                    | 0 days   | 1        |
| PROMOTE    | \<who\>                                   | 0 days   | 1        |
| PUBLIC     |                                           |          |          |
| QUARRY     | [number of stones] [days]             | as given | 3        |
| QUEST      |                                           | 7 days   | 3        |
| QUIT       |                                           |          |          |
| RAZE       | [building]                              | varies   | 3        |
| RECRUIT    | [days]                                  | as given | 3        |
| REPAIR     | [days]                                  | as given | 3        |
| RESEARCH   | \<skill\>                                 | 7 days   | 3        |
| RESEND     | [turn]                                  |          |          |
| RUMOR      | \<# lines of text\>                       | 0        | 1        |
| SAIL       | \<direction or destination\> [...]      | varies   | 4        |
| SEEK       | [who]                                   | 7 days   | 3        |
| SELL       | \<item\> \<qty\> \<price\> [have-left]  | 0 days   | 1        |
| STACK      | \<character\>                             | 0 days   | 1        |
| STOP       |                                           |          |          |
| STUDY      | \<skill\>                                 | 7 days   | 3        |
| SURRENDER  | \<character\>                             | 0 days   | 1        |
| TAKE       | \<who\> \<item\> \<qty\> [have-left]    | 0 days   | 1        |
| TERRORIZE  | \<who\> \<severity\>                      | 7 days   | 3        |
| TRAIN      | \<kind\> \<days\>                         | as given | 3        |
| UNGARRISON | \<garrison\>                              | 1 day    | 3        |
| UNLOAD     |                                           | 0 days   | 3        |
| UNSTACK    | \<who\>                                   | 0 days   | 1        |
| USE        | \<skill\> [arguments...]                | varies   | 3        |
| VIS_EMAIL  | \<new email address for the player list\> |          |          |
| WAIT       | conditions                                | varies   | 1        |

## Submitting orders

Orders for Olympia game 4 should be sent to <olympia@shadowlandgames.com>.

The 'Reply-To:' header on turn reports is set to this e-mail address, so using the reply feature of your e-mail client should send orders to the right place.

Your orders are automatically loaded into the game and queued for your units. The scanner will send a reply as soon as it processes your mail, showing whether or not there were any errors with the orders it received.

Note that the order scanner does not do an exhaustive check of your orders' syntax; it checks that the commands given exist, and checks the parameters of those commands that are executed at parse time - notably those involved in parsing (begin, unit, password, email and vis_email), in report formatting (format and notab) and those which have immediate secondary effects (resend and lore). Other commands' parameters are only checked when they are executed during the turn run.

Orders must be of the following form:

begin player-number password

email, lore, password commands

unit player-number

commands for player entity: name, format or quit

unit unit-number

commands for unit

unit unit-number

commands for unit

end

The \`Subject:' line on your message is ignored.

The begin keyword tells the order scanner what your player number is. If you have not set a password, you do not need to supply one.

The unit command replaces a set of orders for a unit. Any pending orders for the unit will be cleared, and the new orders sent in will queue up. Orders that are still executing for the unit will not be interrupted unless the first order queued is the stop order.

Do not match an end for every unit command! There should only be one end, at the end of all of the unit sections. The Olympia order parser will not read beyond the end.

For example, here is a set of orders for player Fate [812], who has two characters, Osswid [5499] and Candide [1269]:

begin 812

password sneaky

unit 5499

explore

move east

explore

study 600

unit 1269

move north

stack 5499

end

The parser tries to be as flexible as possible. It is case insensitive and is not strict about spaces on a line, so you may use indentation to make your orders more readable.

Orders may be commented with the \`#' character. Everything from a \`#' to the end of the line will be ignored by the parser:

move north \# Head to Drassa to meet up with Osswid

stack 5499 \# stack with him

No one will read the comments but you. Neither the GM nor the Olympia engine will try to interpret them for any reason.

Note that arguments must be enclosed in quotes if they are more than one word:

name "Osswid the Constructor"

The acknowledgement will show any errors that occurred while the orders were being parsed, and list the current pending commands for all of your units.

There is a limit of 250 orders which may be queued per unit. Additional orders will be dropped and will not appear in the unit's command queue.

## Interrupting orders

Suppose the turn report shows the following orders queued:

unit 5499

\# \> study 160 (executing for three more days)

recruit 10

explore

Sending in new orders for this unit will not disturb the still-running study command unless the first order is stop.

For example: If this were sent in:

unit 5499

move south

This would be the result:

unit 5499

\# \> study 160 (executing for three more days)

move south

To interrupt the study and get on with the move right away, instead send in:

unit 5499

stop

move south

This will show:

unit 5499

\# \> study 160 (executing for three more days)

stop

move south

Note that the stop queues like any other order; it does not actually interrupt the executing command until the turn runs. This means that the stop itself can be replaced by sending in another set of orders later.

## Units not controlled by you (yet)

Orders may be sent in for units which are not yet under control, such as characters that you intend to bribe or terrorize into switching to your faction.

As soon as the unit comes under your control, the orders queued for it will begin to execute.

Orders may also be sent in for new nobles which will be formed during the turn. First choose one of the possible unit numbers from the choices listed near the beginning of the turn report:

The next five nobles formed will be: 5717 3215 4902 4489 5628

Supply one of these numbers as the first parameter to the form order:

form 5628 "Feasel the Wicked"

Then queue some orders for Feasel to execute as soon as he appears:

unit 5628

unstack

study 160

move out

recruit

...

## Use the order template

An order template appears at the bottom of the turn report.  This template lists all of the units for a player and shows any pending orders for those units.

Order template

---------------------

begin 812 \# Master Bogomil's Family

unit 2508 \# Tudor

\# \> make 74 (still executing)

unit 2947 \# Milo

unit 4375 \# Beorn

\# \> move s (executing for one more day)

pillage

recruit

unit 4763 \# Sylvia

unit 5977 \# Drango

\# \> collect 87 0 0 (still executing)

sail e

sail s

fish

explore

unit 5418 \# Comte de le Sang

end

Note that the layout of the order template matches the syntax the order scanner expects. Many players find it convenient to edit this template to add or change commands for their units. Mail everything from the begin to end (inclusive) to the order scanner.  It is wise to save a copy of the orders you submit in case there are errors and they need to be resent.

------------------------------------------------------------------------

**Be careful**

Beware of sending in different sets of orders too quickly. Sometimes messages sent within a short time of each other will arrive out of order. This can wreak havoc on your turn if the wrong orders arrive last. Compose your orders offline and review them before mailing. A simple typographical error in your orders could ruin your whole turn!

Some players make clever use of the PASSWORD order to make sure that an order set lingering out in the email system on the network which arrives late won't replace a more recent order set sent in. For example, say you submit some orders on Friday, and they don't show up by Monday. Monday you send an updated set of orders to order scanner and get an instant reply. The Friday orders have not arrived yet, but you're worried that they're out there and will arrive sooner or later, replacing the newer set of orders you just sent in. Solution: issue a PASSWORD order in the newer set, so the Friday orders will fail when they do arrive.

## Failed orders

Commands that fail generally take zero time.

For instance, if study is issued for a skill which the location does not offer, it will fail immediately, and take zero time. The failed study order will not take a week, and it will not count toward the limited study time for that month.

Production commands fail immediately if none of their input resources are available. For instance, recruit in a location with no peasants will immediately fail, taking zero time. However, resources may sometimes become depleted while the command is being executed. In such cases, the command may fail even after it has spent some time executing. 

## More order examples

You can replace orders for some units, but leave pending orders for other units alone, by only including unit sections for the ones you want to change.

If you want to see what orders are queued, but not change anything send in:

begin 999 \# whatever your player number is ...

end

To change your email address, send in:

begin 999

email <new@address.com> \# give your new address

end

As a security measure, the confirmation will be sent to both the new and old addresses.

To change the name of your faction, issue the name order for the faction's player entity:

begin 999

unit 999

name "Seekers of Fame and Power" ...

end

Important: Don't forget the unit command for the player entity.

## Player Entity

Each Olympia player faction has a number. This number is represented by an entity in the game. However, unlike a character, this entity is mostly used as a place holder for the faction. No one can see the faction entity, and it can issue very few orders.

For example, suppose player Fate [501] has one character Osswid [5499]. Fate does not exist in any location, so it does not receive a location report, and no one can see it. However, Osswid, being the player character for the faction, is sworn to Fate [501]. Fate may execute only a limited set of administrative orders.

Characters, not factions, issue most orders. Do not try to form or recruit with the player entity. For most turns, the player entity will have no orders queued for it.

The orders a player entity may issue are:

- accept
- admit
- default
- defend
- format
- hostile
- message
- name
- neutral
- notab
- press
- realname
- rumor
- times
- quit

## Quitting

To drop out of the game, issue the quit order for your player entity.

For instance, player 501 would quit by sending in the following orders:

begin 812 password

unit 812 \# don't forget unit for the player unit!

quit

end

No turn report will be sent for the turn in which a player quits.

## Order of characters in a location

When a character enters a location, it is added to the end of the list of characters already there. The unit that has been in a location the longest will appear at the top of the list.

If a character leaves a location and later returns, it will be put at the end of the list again.

For example:

Seen here:

Candide [1269]

Osswid [5499]

Feasel the Wicked [1109]

Candide has been here longest, followed by Osswid, then Feasel. If Candide were to leave and return, he would appear at the end of the list.

If a character unstacks from beneath another unit, the character will appear just after the unit, rather than at the end of the list.

For example:

Seen here:

Candide [1269], accompanied by:

Osswid [5499]

Feasel the Wicked [1109]

Osswid is stacked beneath Candide. If Osswid unstacks, he will appear after Candide, not after Feasel:

Seen here:

Candide [1269]

Osswid [5499]

Feasel the Wicked [1109]

Note: new player characters are added to the top of the list of characters in the safe haven in which they join, not the bottom. Additonal characters formed by them will appear at the bottom of the list as usual.

## Command priority

All orders have a priority of 0-4.

- Permission commands (admit, hostile, etc.) are priority 0.
- Zero-time commands and wait are priority 1.
- move and fly are priority 2.
- The sail command is priority 4.
- All other commands are priority 3.

The order scheduler will first try to start all priority 0 orders. Only when no more priority 0 orders are ready to start will a priority 1 order be started.

In other words, the order scheduler will not start an order at a higher priority when an order may be started at a lower priority.

Orders at the same priority are resolved in location order. If two units in a location are both waiting to start a move order, the first unit in the location will go first.

The above description of order priorities may seem complicated, but the intent is to let players ignore same-day synchronization issues in most cases. Rather that needing wait to guarantee that give happens before move, the lower priority of give makes this happen naturally.

For example, consider three units stacked together, top, mid and bot:

top:

move ec69

yew

mid:

unstack

recruit

bot:

recruit

These should be executed in the following order:

mid: unstack \# unstack is prio-1

top: move ec69 \# move is prio-2

mid: recruit \# recruit is prio-3 [top and bot arrive at ec69]

top: yew \# yew is prio-3

bot: recruit \# recruit is prio-3

The unstack happened first since it's a priority 1 command. The move went second. When top and bot finished moving, there were only priority three commands left, so they ran in location order.

------------------------------------------------------------------------

**Command precedence within a location**

Seen here:

Candide [1269]

Osswid [5499]

Feasel the Wicked [1109]

Order precedence within a location is an advantage for commands or skill uses which obtain resources from the location. For instance, if Candide and Osswid both attempted to harvest all of the lumber available in their location, Candide would have precedence, since his harvest order would finish before Osswid's, if they were started on the same day.

------------------------------------------------------------------------

**At the same time ...**

No two things ever happen at exactly the same instant in Olympia. Someone always goes first.

Suppose two characters were outside of a building (which nobody is inside), and both wanted to enter, to claim it:

Seen here:

Osswid [5499]

Candide [1269]

Inner locations:

Hooting Own Inn [ep76], inn

Both Osswid and Candid issue \`move ep76' as their first order on day 1 of the month. What happens?

Osswid's command begins before Candide's, since Osswid appears before Candid in the location list. Therefore, Osswid will enter the inn first.

    Osswid             ?            Candide
    ------- 8 days -------- 8 days --------
    city A          city B          city C

If Osswid and Candide both leave for city B on the same day, we cannot predict who will get their first.

## Olympian geography

Olympia's map is a large grid of locations called provinces. Groups of provinces form continents, islands and oceans. These collections are called regions, and are usually named.

A province's description will include a list of the directions in which a character may travel:

Plain [ae48], plain, in region Tollus

Routes leaving Plain [ae48]:

North, to Plain [ad48], 7 days

East, to Plain [ae49], 7 days

South, to Ocean [af48], Tymaerian Sea, 1 day

West, to Ocean [ae47], Tymaerian Sea, 1 day

This is a non-descript province in the Tollus region.

From this province, a character may travel north or east on foot or by horse, or may sail by ship to the south or west.

move north -or- move n -or- move ad48

move east -or- move e -or- move ae49

sail south -or- sail e -or- sail af48

sail west -or- sail w -or- sail ae47

Land movement will automatically use the fastest available mode. For example, if a character has enough horses for all of the members in the party to ride, then the travelers will go on horseback.

Ocean movement requires that the character be in a ship.

Route distances are rated for the number of days it normally takes to traverse them. Land distances are rated for a lightly loaded character walking, and ocean distances are given for an ordinary ship traveling in normal weather.

Actual travel times may differ from times given in the route listing. Land distances depend on the surrounding terrain and the modes of transport available. For example, horses often speed up movement, but over especially rough or treacherous terrain, they may actually slow travel because they must be led and managed. A stiff wind may speed ocean vessels, while lack of wind may slow their progress.

### Inner Locations

A province may contain sub-locations within its borders. Sub-locations may usually only be entered from the surrounding province. They will be listed separately in the location description:

Inner locations:

Carim [em28], city, 1 day

The city Carim may be entered with the \`move em28' order. Travel into a city requires one day.

*Note: move in may be used to enter a sub-location, although this order may be ambiguous if the location contains more than one sub-location. In such a case, the first sub-location in the Inner locations list will be entered. Using move in is not recommended if the entity number of the sub-location is known.*

Characters in a sub-location will receive a report for the surrounding province. However, characters in the outer province will not normally be able to see into an inner location without entering it.

### Inside a City

Carim [em28], city, in province Plain [ae48]

Routes leaving Carim [em28]:

Out, to Plain [ae48], 1 day

Inner locations:

Hooting Own Inn [ep76], inn

Characters in the city Carim may move out (or \`move ae48') to the surrounding province. They may also attempt to enter the inn, which is a sub-location of the city. Notice that no travel time rating is listed for the inn; entering it takes no time (zero days).

move out -or- move ae48

move ep76

A character in Carim will receive a location report both for the city as well as the surrounding province Plain [ae48]. Characters in the Hooting Own Inn will receive a location report for the inn and one for Carim, but will not get a report for Plain [ae48]. A character in the city may not see inside the inn without entering.

Characters in a sub-location receive a report for the immediate surrounding location.

Characters are not able to see into inner locations without going into them.

------------------------------------------------------------------------

**City default garrisons**

Every non-safe-haven city in the regular world (except safe havens) has an initial default garrison with 25-150 pikemen. Each garrison is set to *admit all* and *defend all*. Each noble stacked with the garrison will earn 2gp/day for aiding the city's defenses.

### Who else is here?

Characters spotted will be listed in the location report:

Seen here:

Fighters of Pelenth [2019], "carrying a gold banner"

Osswid the Constructor [5499]

All characters listed as being seen in the location may interact without requiring any travel. Thus, the Fighters of Pelenth and Osswid are considered to be in essentially the same place.

This is true whether the characters are in a province, a city, a ship, an inn, or some other sub-location.

However, a character in a sub-location may not interact with characters in the surrounding area. A noble in the city must first enter the inn before he may interact with those inside.

### More about geography

Olympian provinces are arranged in a square grid. Travel is possible in the four main compass points. Thus, to move diagonally, two move orders are required. To move northwest, for instance, one would first need to \`move n', then \`move w'.

The map coordinates for a province may be read from the province's entity number. The row is represented by the leading two letters, the column by the two digits. The northwest corner is [aa00], with rows increasing to the south, and columns increasing to the east.

        00   01   02  ... 79
      +----------------------
    aa|
    ab|          ab02
    ac|    ac01  ac02    ac99
    ad|          ad02
    af|
    ...  
    ...
    dz|          dz02

Row sequence is: "abcdfghjkmnpqrstvwxz"

Entity numbers for sub-locations do not correspond to any coordinate system.

The edges of the map are not passable, so for example it is not possible to travel either north or west from aa00.

### Holes in the Map

The map may have some holes, representing impassable provinces. Routes into some provinces may also be hidden.

Plain [cd21], plain, in region Tollus Routes leaving Plain [cd21]:

North, to Plain [cc21], 7 days

East, to Plain [cd22], 7 days

West, to Plain [cd20], 7 days

Notice the lack of a southern exit. This means that there is no known southern route from Plain [cd21], into what should be Plain [ce21]. Exploration may find a southern route, but it is possible that none may ever be found, and the terrain to the south is completely impassable.

### Hidden routes

If exploration finds a hidden route, any noble in the player's faction will be able to use it.

\> explore

A hidden route has been found!

South, to Plain [ce21], 7 days

The location description for this place will now include the hidden route:

Plain [cd21], plain, in region Tollus

Routes leaving Plain [cd21]:

North, to Plain [cc21], 7 days

South, to Plain [ce21], 7 days, hidden

East, to Plain [cd22], 7 days

West, to Plain [cd20], 7 days

However, units from other factions, even if they know that the hidden route's entity number is [ce21], will not be able to travel across it.

All factions with units in a stack traveling across a hidden route, with the exception of units being held prisoner, will learn of its existence. Nobles from factions wanting to learn how to use the hidden route can stack with a noble about to move across the route.

### Ocean ports

A ship in an ocean province may sail into an adjoining land province.

Ocean [cw12], ocean, in South Sea

Routes leaving Ocean [cw12]:

North, to Ocean [cv12], Atnos Sea, 4 days

East, to Mountain [cw13], West Camaris, impassable

South, to Plain [cx12], West Camaris, 1 day

West, to Ocean [cw11], 4 days

Inner locations:

Island [eb97], island, 1 day

A ship sailing in this ocean province may dock by sailing to Plain [cx12] or Island [eb97].

Ships may not dock in mountain provinces, as the rocky cliffs are too dangerous to approach. Routes between ocean and mountain provinces are marked \`impassable'.

### Port Cities

A city in a province adjoining an ocean will have been founded on the best spot for an ocean port. The ocean will only be accessible through the port city in this case, and not through the surrounding region.

Plain [ae48], plain, in region Tollus

Routes leaving Plain [ae48]:

...

West, to Ocean [ae47], Tymaerian Sea, impassable

Inner locations:

Carim [em28], port city, 1 day

Note that from the province surrounding the port city, access to the ocean is not possible.

Carim [em28], port city, in province Plain [ae48]

Routes leaving Carim [em28]:

West, to Ocean [ae47], Tymaerian Sea, 1 day

Out, to Plain [ae48], 1 day

However, ships may sail into and out of the port city itself. From the Tymaerian Sea, this looks like:

Ocean [af48], ocean, in Tymaerian

Sea Routes leaving Ocean [af48]:

North, city, to Carim [em28], Tollus, 1 day

South, to Ocean [ag48], 3 days

### An example city description

Drassa [ew66], port city, in province Forest [cu26], safe haven

Routes leaving Drassa [ew66]:

East, to Ocean [cu27], Atnos Sea, 1 day

South, to Ocean [cv26], Atnos Sea, 1 day

Out, to Forest [cu26], 1 day

Skills taught here:

Shipcraft [600]

Combat [610]

Construction [680]

Seen here:

Kosar the Indefectible [2022], with six peasants, one archer, two soldiers, accompanied by:

Dr. Pangloss [3682]

Law Netexus [2020], prisoner

Alion Krysaka [2785], prisoner

Ships docked at port:

HMS Pinafore [ib18], galley, owner:

Captain McCook [2019], with five workers

Market report:

No goods offered for trade.

### Wilderness and civilization

Every province has a civilization level. Provinces with no civilization (a level of zero) are considered wilderness. Civilization levels for provinces are shown in the turn report:

Mountain [cq24], mountain, in Lesser Atnos, civ-1

Forest [ac35], forest, in Torba Bacor, wilderness

The civilization level of a province is determined by the presence of cities and buildings, or half of the maximum civilization level of its surrounding provinces, whichever is higher.

There is no fixed civ level cap. However, only the first building of each type counts towards the civ level in a location. 

| Feature    | Contribution                |
|------------|-----------------------------|
| Safe Haven | 2                           |
| Castle     | 1.5 + improvement level / 4 |
| City       | 1                           |
| Tower      | 1                           |
| Temple     | 1                           |
| Inn        | 1                           |
| Mine       | 1                           |

Any fractional remainder is dropped after the contributions are summed.

Only the first building or feature of each type counts toward the civilization level. For example, if two inns in a province, only the first would add a civ point to the total.

### Dangerous places

Players should take care when exploring the Olympian world. There are many dangers, both from non-player characters (NPC's) as well as from other players. While the threat of death to nobles is always present, the following dangerous areas warrant extra caution.

- Hades

Hades, also known as The Land of the Dead, is a subterranean world populated with demons, ghouls and spirits thirsty for the blood of the living. Only the bravest warriors should consider walking these dark paths.

- Faery

The Faery world lies in a nearby, but separate reality. Occasionally a Faery hill will exist simultaneously in both Faery and the outside world. During this time, mortals may cross between the two worlds. Faery is protected by the Faery Hunt, a tough band of elves armed with magical bows. Each hunt group consists of 10-50 elves, each with a combat rating of (50,50,100). Rumors speak of a magical talisman, the elfstone, which allows mortals to pass unharmed in Faery, and to summon Faery hills to the mortal world.

- The Cloudlands

The Cloudlands is a small region which floats over Mt. Olympus and the Imperial City. It is generally only accessible by flight. The Cloudlands is home to three cities: Nimbus, Stratos and Aerovia. Weather magic is taught in these three cities.

### The Map

## Health

Nobles have a health rating of 1-100, which indicates how wounded they are. A noble with health 100 has no wounds; a noble with a health of zero dies. Nobles also have a flag which indicates whether they are suffering from an illness. A sick noble will lose some health each week, while a wounded noble who is otherwise free of illness will recover somewhat each week.

Health is shown in the turn report for each noble that the player controls.

This is a noble in perfect health:

Health: 100%

This is a rather sick noble:

Health: 38% (getting worse)

If the noble were cured of illness, this would instead show:

Health: 38% (getting better)

Medical technology is rather crude in the age of Olympia. Sanitation and hygiene are not the best. Even a minor wound runs a risk of developing into a serious, possibly life-threatening problem.

When a noble receives a new injury, their health is reduced by the amount of the injury, and a check is made to see whether they get sick. The chance that a character falls ill is (100 - health). Thus, the more seriously the noble is wounded, the greater the probability that infection will set in.

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

When a noble is killed in battle or dies, the body is moved into the province as an item which may be found with EXPLORE. An executed noble's body is given to the character who performed the execution.

A dead body exists for one and one half game years, at which point it fully decomposes, and the dead noble's spirit passes on.

Example: A dead body rots after twelve months have passed, so if a noble dies in turn 20, his body will decompose at the end of turn 32.

The bodies of nobles lost at sea will wash ashore somewhere.

Bodies decompose after one and one half years regardless of whether they remain in the province or are possesions of another.

Priests may learn a skill Lay to rest, hastening the passing of the dead noble's spirit. Some exceptionally skilled priests possess the ability to resurrect dead characters.

In general, NP's invested in characters are returned to the character's player if the character deserts or his body decomposes.

Characters which renounce service becaused of lapsed contract or fear loyalty do not immediately return the NPs to the previous owner. The previous owner of these characters will not receive NPs for them until they swear to a new player faction or die.

When a body decomposes, the number of NP's which were invested in the character are returned to the original owner.

## Markets

Characters issue buy and sell commands to indicate their desire to trade goods. Every trade must be between a buyer and a seller. Whenever a compatible buy and sell request are found, the trade will be executed.

Trades are only matched in cities, where merchants gather at the local bazaar to swap goods. (Note that characters are always free to use the give command to exchange items, regardless of where they are. But buy and sell orders are only matched in cities.)

A buyer indicates what he wants to buy, how much of it he wants, and the most he is willing to pay. A seller indicates what item is being sold, how much of it is to be sold, and the per-item price.

A trade will match if the seller's price does not exceed the maximum price specified by the buyer; if the buyer has enough gold to buy at least one of the item; if the seller possesses at least one of the item; and if the buyer and seller are both in the same city.

If the buyer is willing to pay more than the seller is asking, the trade takes place at the seller's price.

Example: A noble who wants to buy five iron [79] at no more than 10 gold each issues:

\> buy 79 5 10

Try to buy five iron [79] for 10 gold each.

If someone who had iron issues a sell order which matches this buyer's request, the trade would be executed:

\> sell 79 5 10

Sold five iron [79] for 50 gold.

The buyer would see:

Bought five iron [79] for 50 gold.

Either the buy or the sell order could have been issued first. If the order can't be matched with pending trades from other units in the city, it will become a standing order and remain in effect until executed or canceled.

\> buy 79 0

Cleared pending buy for iron [79].

Note that the buyer and seller can't specify what character they will deal with. They will trade with any unit that has a matching buy or sell order.

### Market report

The location report for each city includes a market report listing pending trades.

| trade | who  | price | qty | item               |
|-------|------|-------|-----|--------------------|
| buy   | 2019 | 100   | 1   | peasant [10]     |
| buy   | 4274 | 74    | 3   | iron [79]        |
| sell  | 3682 | 12    | 11  | sailor [19]      |
| sell  | 2019 | 50    | 1   | elite guard [15] |

A trade will not be listed unless it could be executed. For instance, a unit might issue an order to sell iron [79], even though the unit doesn't possess any. (Perhaps the character plans to get some iron later in the month, and wants the buy order to be in place when the iron arrives). This order will not be shown in the market report, because the seller doesn't have any iron to sell.

Similarly, the buyer must have enough gold to buy at least one of the item. No buy order would be listed for a penniless unit that wanted to buy five iron at 10 gold each. If the unit later obtained 10 gold, enough to buy one of the five desired units of iron, the order would be listed in the market report for one iron, not five.

Some traders may work through middlemen to hide their identity. In such cases, the who field of the market report will not show their unit number, and their identity will not be revealed when the trade is executed.

Even sneakier traders may have their pending trades omitted from the market report entirely.

### City purchasing

Some cities will issue buy or sell orders on their own behalf for certain goods. Cities trade orders are identity to those submitted by players when participating in markets.

Enterprising characters may turn a profit by using skills under the Trade [730] category to find new tradegoods for sale in city markets, and buyers for those tradegoods in other city markets.

### Resolution of trades

If a buyer has a choice between two or more sellers, the one offering the lowest price will be chosen. If several sellers offer at the same price, the one nearest the top of the location's character listing will win the trade. A seller with a choice among two or more buyers will pick the highest unit as it appears in the location's character listing.

Since characters are added to the end of the list when they enter a location, the units who have been in a place the longest tend to appear toward the top of the list. Characters who have been in a place the longest have an advantage when one of several possible trades may be matched.

However, consumption or production by the city itself will always have lowest priority. Cities defer to characters when multiple matching trades are possible.

The location market report is ordered according to how multiple trades will resolve.

| trade | who  | price | qty | item        |
|-------|------|-------|-----|-------------|
| sell  | 3682 | 10    | 1   | iron [79] |
| sell  | 2019 | 12    | 1   | iron [79] |

A buyer would get 3682's iron, since the price is lower.

| trade | who  | price | qty | item        |
|-------|------|-------|-----|-------------|
| sell  | 2019 | 10    | 1   | iron [79] |
| sell  | 3682 | 10    | 1   | iron [79] |

 A buyer would get 2019's iron. 2019 must appear before 3682 in the location character list.

| trade | who  | price | qty | item        |
|-------|------|-------|-----|-------------|
| buy   | 4846 | 10    | 1   | iron [79] |
| buy   | 1783 | 12    | 1   | iron [79] |

4846 would win the buy from a lone seller. 1783 must come after 4846 in the location character list.

| trade | who  | price | qty | item        |
|-------|------|-------|-----|-------------|
| buy   | 4846 | 10    | 1   | iron [79] |
| buy   | 1783 | 11    | 1   | iron [79] |

sell 79 1 10 will match 4846. sell 79 1 11 will match 1783.

### Partial trades

Partial trades are executed, if possible. If a buyer wants 100 iron, but only 25 iron are available, he will buy 25. The pending buy order will be reduced so that 75 more will be bought whenever possible.

### More market examples

A buyer who desires three iron at no more than five gold each orders:

\> buy 79 3 5

Try to buy three iron [79] for 5 gold each.

Suppose the only seller of iron in this city had ordered:

\> sell 79 5 6

Try to sell five iron [79] for 6 gold each.

The trade will not take place, because the seller's price exceeds the buyer's. Later, the buyer travels to a different city, where another seller has previously issued:

\> sell 79 10 4

Try to sell ten iron [79] for 4 gold each.

As soon as the buyer enters the city, the trade will be matched:

Bought three iron [79] for 12 gold.

Note that the trade takes place at the seller's price, and that the seller will still be offering seven iron at four gold each.

If the seller had only one unit of iron for sale, the trade would have executed, but the buyer would still be looking for two more iron to buy, at five gold each or better.

## Loyalty Bonds

Nobles are bound to their lords by one of three kinds of loyalty:

*contract*

The noble is being paid for his service

 

*oath*

The noble has taken an oath to serve his master

 

*fear*

The noble serves out of fear

 

The loyalty bond is rated. For example, a character's loyalty could be oath-1, contract-500, or fear-50.

A player's initial character has oath-2 loyalty.

Newly hired nobles have loyalty contract-500. Nobles are paid for their service with the `honor` command. A noble who issues \``honor`` 50`' will spend 50 gold to raise his own contract loyalty rating by 50 points.

A noble may take an oath of loyalty, pledging one or two Noble Points to secure it. This would yield loyalty oath-1 or oath-2. The `oath` order secures an oath of loyalty for a noble.

Nobles may be terrorized by their masters. The severity of their treatment accumulates in the loyalty rating: fear-10, for instance. Fear loyalty is maintained with the `terrorize` order.

Only one kind of loyalty may be active at a time.

Contract and fear loyalty decay over time. Contract loyalty loses the greater of 50 points or 10% of the current rating each month. Fear loyalty loses 1-2 rating points each month. Oath loyalty does not decay.

Units which fall to contract-0 or fear-0 have a 50% chance of deserting each month.

Nobles serving through contract or fear are susceptible to bribes, which may induce them to renounce loyalty to their lord, and pledge their service to the bribing faction. For details on bribing characters, see the `bribe` order.

Oath-1 nobles ignore all bribes. There is a persuasion skill which may cause an oath-1 noble to defect, although its use is difficult and rarely succeeds. Oath-1 nobles may reveal their factional affiliation if tortured.

Oath-2 nobles will not renounce loyalty to their lord under any circumstances, nor can they be forced to reveal any information about themselves.

Summary:

*contract*

Amount of gold invested in noble

Decays by max(50, 10% of current rating) each month

*fear*

Severity of `terrorize` used on noble

Decays 1-2 points each month

 

*oath*

1 or 2 NPs may be invested in an oath bond

Does not decay

Commands dealing with loyalty bonds:

- `bribe`
- `honor`
- `oath`
- `terrorize`

## Stacking

One unit may `stack` under another unit. Two or more units grouped in this way are referred to as a *stack*.

Stacks move together and fight together. Here is a stack of four units:

    Law Netexus [2020], accompanied by:
      Feasel the Wicked [1109]
      Drakkar the Trader [1752]
      Alion Krysaka [2785]

 

Law Netexus is the stack leader, the top-most unit in the stack.

Only one level of stack depth is shown, so all that can be determined from the location report is that Feasel, Drakkar and Alion are stacked somewhere beneath Law Netexus. The exact arrangement of stacking bonds is not shown.

Feasel might be stacked under Law Netexus, with Drakkar and Alion under Feasel. Or Feasel, Drakkar, and Alion may all be stacked directly beneath Law Netexus.

Generally, such internal arrangements are only important when the stack breaks up. If Drakkar is stacked beneath Feasel, he will stick with Feasel if Feasel drops out of the stack. But if Drakkar were stacked beneath Law Netexus, he would not follow Feasel if Feasel unstacked.

If Law Netexus issues a `move` order, the entire stack will move. If, however, Feasel issues a `move` order, he will first drop out of the stack before moving.

Similarly, if Law Netexus engages combat with the `attack` order, all characters in the stack will fight together. If one member of the stack is attacked, the entire stack will respond in defense.

Multiple levels of internal stacking can be useful if one wants several stacks to join together for a while, but then split apart later into their old arrangements.

Ocean ships may not be stacked together. There is no way to cluster ships into a fleet.

## Carrying capacity

Carrying capacity: Men and items are rated for how much they weigh, and how much they can carry, walking, riding or flying.

|              |        |         |        |        |
|--------------|--------|---------|--------|--------|
| Name         | Weight | Walking | Riding | Flying |
| Man          | 100    | 100     | \-     | \-     |
| Riding Horse | 1000   | 150     | 150    | \-     |
| Wild Horse   | 300    | self    | self   | \-     |
| Warmount     | 300    | 150     | 150    | \-     |
| Knight       | 400    | 100     | 100    | \-     |
| Elite Guard  | 400    | 100     | 100    | \-     |
| Ox           | 1000   | 1500    | self   | \-     |
| Winged Horse | 300    | 150     | 150    | 150    |

Carrying capacitiy {border="1" cellpadding="1" cellspacing="1"}

\``man`' includes all of the varieties of men, including peasants, sailors, workers, etc. as well as nobles. A knight includes both the man and the horse, hence the 400 weight.

In order to ride, the total riding capacity must cover the weights of all the units that may not ride themselves.

Examples:

*man + riding horse*

    ride capacity is 150 - 100 = 50
    walk capacity is 250

 

*man + riding horse + wild horse*

    ride capacity is 150 - 100 = 50
    walk capacity is 250

 

The wild horse can walk or ride on its own, but will not carry anything.

    ride capacity is 150 - 100 = 50
    walk capacity is 1750

 

The ox may be driven alongside the horse, but will not carry anything when moving so quickly.

*man + riding horse + ox*

A stack will ride if there is enough riding capacity to carry all of the non-riders. Otherwise, the stack will walk.

In order to fly, the total flying capacity for the stack must cover the weights of all units that can not fly themselves.

Stacks which are overloaded beyond their walking capacity will travel slower than normal. A stack loaded at 150% of its walking capacity will take 50% longer to traverse a route. Stacks overloaded to over 200% of walking capacity may not travel at all.

Weights and capacities are always considered for the stack as a whole. One unit may have all the men, and another unit may have all the horses. If they are stacked together, the distribution of items across units is irrelevant.

## Training men

Nobles may be accompanied by different kinds of men. Men may be peasants obtained with the `recruit` order, or peasants who have been trained into other kinds of men such as sailors, soldiers or workers.

The different kinds of men:

- peasant [10]
- worker [11]
- soldier [12]
- archer [13]
- knight [14]
- elite guard [15]
- pikeman [16]
- blessed soldier [17]
- sailor [19]
- swordsman [20]
- crossbowman [21]
- elite archer [22]
- pirate [24]

A peasant may be trained into a soldier, sailor, worker, or crossbowman. Soldiers may be given further training to become pikemen, swordsmen, or archers. More advanced fighters may be trained from swordsmen and archers.

         +--- sailor ------- pirate
        +---- worker
       +----- crossbowman
      /
    peasant ------ soldier ------ swordsman --- knight ----- elite guard
                      \
                       +---- archer ------ elite archer
                        +--- pikeman
                         +-- blessed soldier

 

Training a man may require that the noble have a certain skill, or possess some item. For instance, training soldiers into swordsmen requires a longsword [74] for each swordsman produced. The training character must also know Combat [610].

Some men may only be trained in certain kinds of locations. Elite guard and elite archers, for instance, may only be trained in castles. Blessed soldiers may only be trained in temples.

Training takes one day per man. Training five archers into five elite archers would take five days, for example. Training ten peasants into ten crossbowmen would take 10 days.

    num  kind            skill  input man       input item        
    where
    ---  -----------     -----  --------------  ----------------  ------
     11  worker           none  peasant [10]
     12  soldier           610  peasant [10]
     13  archer            615  soldier [12]    longbow [72]
     14  knight            616  swordsman [20]  warmount [53]
     15  elite guard       616  knight [14]     plate armor [73]  castle
     16  pikeman           610  soldier [12]    pike [75]
     17  blessed soldier   750  soldier [12]                      temple
     19  sailor            601  peasant [10]
     20  swordsman         616  soldier [12]    longsword [74]
     21  crossbowman       610  peasant [10]    crossbow [85]
     22  elite archer      615  archer [13]                       castle
     24  pirate            616  sailor [19]     longsword [74]    ship

 

- A character needs no skills to train a worker.
- To train a sailor requires Sailing [601], a subskill of Shipcraft [600].
- Training archers and elite archers requires Archery [615], a subskill of Combat [610].
- Training swordsmen, knights and elite guard requires Swordplay [616], a subskill of Combat [610].

For more information and examples, see the `train` order.

## Maintenance cost

Men such as soldiers, workers, archers, etc. must be paid in gold monthly or they will leave the service of their noble. Peasants do not willingly leave a noble's service, but will starve if they are not paid. This cost is charged to the noble holding them at the end of each month.

If the noble does not have enough gold to pay his men, he will ask other nobles in his stack (provided they belong to the same player) for gold. Thus, only one member of a stack need carry gold for maintenance costs for the entire stack. Nobles will not share gold with units from other players.

If the noble can only afford to pay some of his men, one-third of those not paid will leave service at the end of the month. The computer chooses which men remain and which leave or starve.

    num   kind            cost
    ---   -----------     ----
     10   peasant          1
     11   worker           2
     19   sailor           2
     21   crossbowman      2
     12   soldier          2
     13   archer           3
     16   pikeman          3
     17   blessed soldier  3
     20   swordsman        3
     24   pirate           3
     14   knight           4
     22   elite archer     4
     15   elite guard      5

 

Note that nobles may `drop` men to release them from service deliberately.

## Making weapons and armor

Weapons and armor are required for the training of some kinds of fighters. Archers require longbows, for instance, elite guard require plate armor, etc.

Weapons and armor are made with the `make` command. The Weaponsmithing [617] subskill of Combat [610] is required to make weapons and armor.

    num   item           material
    ---   ----           --------
     72   longbow        yew [68]
     73   plate armor    iron [79]
     74   longsword      iron [79]
     75   pike           wood [77]
     85   crossbow       wood [77]

 

One unit of the input material may turned into one of the desired items each day. For example, `make` 72 2 would spend two days turning two yew [68] into two longbows [72].

## Skills

Skills represent knowledge that Olympian characters may know. Shipbuilding, thievery, kidnaping, training soldiers, castle building, mixing potions, and forging magical artifacts are just a few of the possible actions which skills allow a character to perform.

Skills are divided into category skills and sub-skills within those categories. The skill categories are:

    num  name                      time to learn
    ---  ----                      -------------
    600  Shipcraft                 three weeks
    610  Combat                    three weeks
    630  Stealth                   four weeks
    650  Beastmastery              four weeks, 1 NP req'd
    670  Persuasion                four weeks
    680  Construction              three weeks
    690  Alchemy                   four weeks
    700  Forestry                  three weeks
    720  Mining                    three weeks
    730  Trade                     three weeks

 

There are also six schools of magic. For more details on learning and casting magical spells, see the *Magical Arts* section.

The category skill must be learned before any of the sub-skills within the category may be known.

With each skill learned, the player will receive a lore sheet describing background information about the skill and how it may be used. Most skills are invoked with the \``use`` ``skill`' order. The skill lore sheets will give specific information about arguments to `use` and and requirements or limitations for using the skill.

The lore sheets for the skill categories list some of the skills available for study within the category.

For instance, a noble wishing to undertake the study of Shipcraft would first learn the category skill with the `study` command:

    study 600

 

The Shipcraft lore sheet lists some of the sub-skills available for building and sailing ships. One of these is Sailing [601], the skill required to control a ship on the ocean. The aspiring captain could then order:

    study 601

 

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

      Skills taught here:
         Alchemy [690]

   

- The skill was discovered through `research`, and is listed as \``Partially known`' in the turn report.\

      Partially known skills:
         Improve ship rigging [9999], 0/7

   

- The player has a book which teaches the skill. Items which teach skills are shown in the inventory listing:\

      Inventory:
                   qty  name
                   ---  ----
                     1  Old book [6001]

      Old book [6001] permits study of the following skills:
         Alchemy [690]

   

### Fee to begin study

A fee of 100 gold is charged when study is first issued for a skill. The payment is used to acquire various materials needed for the study and eventual practice of the skill.

### Advanced skills

Some advanced skills require Noble Points to begin study. Each player begins the game with a number of NPs, and receives an additional one each year. Noble Points may be spent to acquire new nobles, or to learn advanced skills.

Most skills do not require Noble Points to learn. NPs are required for some heroic combat skills and for advanced magical spells.

### Study limit

Characters may STUDY up to 14 days per turn. Fast study days dont count towards this limit.

### Fast study

All new players are given 200+ "fast study" points. Each fast study point may be applied to a skill being studied in lieu of actually spending a day studying.

For example, the order `study 600 7`` would apply 7 fast study days to learning Shipcraft. This study order would take 0 days to execute. `

### Skill Experience

Experience is counted for each turn that a skill use is successfully completed. If a skill is used more than once per turn, only the first success will count towards experience.

Projects which take multiple turns to complete, such as shipbuilding or castle construction, only count towards experience when the project is finished.

    use             level
    -----           -----
     0-4            apprentice
     5-11           journeyman
    12-20           adept
    21-34           master
     35+            grand master

 Experience will speed work with some skills. For example, a master shipbuilder will be able to construct a galley somewhat faster than an apprentice. Some skill uses benefit more from experience than others.

### Skills not rated for experience

A few skills are not rated for experience. These may be skills which are not directly used, or ones for which experience has no meaning. If a skill is not rated for experience, the skill level will not be shown in the skill listing.

For instance:

    Skills known:
         Shipcraft [600]
               Sailing [601], apprentice
               Shipbuilding [602], apprentice
               Fishing [603], apprentice
         Combat [610]
               Survive fatal wound [611]
               Fight to the death [612]

 

Since experience is not applicable to Survive fatal wound [611] and Fight to the death [612], its levels are not shown.

### Summary of studying

The category skill must be learned before a sub-skill within the category may be studied.

Beginning study of any skill costs 100 gold.

Payment of Noble Points may be required to begin learning some advanced skills.

A source of instruction must be available the first time `study` is issued:

- The skill is taught by the location.
- Characters may `study` up to 14 days per month.
- The skill is commonly known, so it may be studied anywhere once its category skill is known.
- The skill was discovered through `research`, so it may now be studied.
- The skill is taught by a book or a scroll.

Once a skill is known, further `study` of that skill has no effect.

### Teaching

Direct character-to-character teaching is not possible in Olympia. However, it is said that some magicians and alchemists possess the ability to record knowledge on scrolls, which other characters may study from.

### An example of study

This is a rough, heavily edited example of how some study commands might look in a turn report. All of the other details that a real turn report would have were omitted to focus on the study orders.

Turn one:

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
       ...

 

Further skills may be found through research.

Turn two:

    1: > study 601
     1: Paid 100 gold to begin study.
     1: Will study Sailing for seven days.
     7: Learned Sailing [601].

 

(It doesn't matter where 601 is studied, since 600 offers it directly.)

Turn three:

    1: > study 690
     1: Instruction in [690] is not available here.

 

(Need to find a location or book that offers instruction in Alchemy)

Turn four:

    1: enter xxxx
     1: Arrival at City of Alchemists [xxxx]

    Skills taught here:
       Alchemy [690]

     1: > study 690
     1: Paid 100 gold to begin study.
     1: Will study Alchemy for seven days.

    Partially known skills:
       690  Alchemy, 7/14

 

(Alchemy requires 14 days of study to learn, seven of which we have completed.)

Turn five:

    1: > study 690
     1: Continue studying Alchemy.
     7: Learned Alchemy [690].

    Skills known:
       600  Shipcraft
            601  Sailing, apprentice
       690  Alchemy

 

## Research

Research attempts to discover sub-skills which are not commonly known or made available when the category skill is learned.

Research is mostly used to discover new magical spells, as few spells are granted when a magic school is learned. However, even common skills such as Shipcraft, Combat and Construction may have hidden sub-skills which can be found through research.

Research for all skill categories except Religion [750] must be performed in a tower, by the tower's owner (the first character inside the tower). Towers make good laboratories for scholarly investigations, and minimize distractions. Other occupants of the tower may not use `research`.

Research into Religion [750] must be performed in a temple, by the temple's owner.

Research by mages of 6th black circle level and above (maximum aura of 31 or higher) must be done in provinces with a civilization level of 1 or less.

When a category skill is learned, its lore sheet appears in the player's turn report, listing information about the skill as well as sub-skills which may be studied directly based on the parent skill.

An example:

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

 

The last line (\``Further skills ...`') indicates that there are sub-skills of Shipcraft which are not mentioned in the lore sheet. These hidden skills may represent rare or hard-to-learn knowledge, or perhaps technology which has not yet been discovered.

Study of these skills is not possible based simply on knowledge of the parent skill. Sailing [601] can be learned by a character, no matter where he is, once Shipcraft is known. Hidden Shipcraft skills, however, must be learned in other ways, even if the character has learned their entity numbers from other players.

There are two possible ways such hidden skills may be learned:

1.  Through a rare book or scroll which offers instruction in the skill.
2.  Through research.

Research is the more difficult choice. However, to learn rare sub-skills there may be no alternative. Perhaps there are no players who already know the rare sub-skill, and so cannot record scrolls to instruct others. Or if there are, they choose to keep their knowledge secret.

Research incurs a fee of 25 gold to pay for miscellaneous materials and costs.

Each week of research yields a 25% chance that a new skill will be discovered.

If research is successful, the new sub-skill will be added to the character's partially known skill list. It then must be studied in order to become fully known and usable.

For example:

    1: > research 600
    1: Will research Shipcraft for seven days.
    7: Research uncovers a new skill:  Improve ship rigging [9999].
    7: To begin learning this skill, order 'study 9999'.

    Partially known skills:
       9999  Improve ship rigging, 0/7

 

The \``0/7`' qualifying the new sub-skill in the partially known skill list indicates that seven full days of study are required to learn the skill, and none of them have yet been completed.

    1: > study 9999
    1: Paid 100 gold to begin study.
    1: Will study Improve ship rigging for seven days.
    7: Learned Improve ship rigging [9999].

 

The `research` order is no longer used on the new skill once it becomes partially known. Research may continue to be used on Shipcraft [600], however, to seek out more hidden sub-skills.

Category skills may not be learned through research.

## Magic

Magic is a dead cat in an oil-stained burlap bag.\
Magic is a smelly old man, despised but feared by his neighbors.\
Magic is what the king turns to, when his soldiers fail.\
        --- a long-dead wise man of Areth Pirn

Magic is the dark art by which events are influenced outside of the normal boundaries of cause-and-effect. Rather than the glamorous ideal of shining wizards casting powerful fireballs at wicked foes, the reality of magic instead tends to be base, tedious work which earns few friends.

Hated and feared, the magician pursues his craft out of the sight of men. Like wisps of smoke rising from an ember cast into dry straw, so the mage's spells slowly take hold, woven with secret knowledge and foul ingredients.

The casting of a magical spell is accomplished with three ingredients: Knowledge of the spell, possession of any items necessary to fuel the spell, and a sufficient level of magical aura to perform the ritual or ceremony.

### Aura

Aura is a mystical force necessary to cast spells. Powerful mages will have a high aura rating, while apprentice sorcerers may only command a few points worth of aura.

Characters are rated for their current aura level and their maximum aura level. With each magical spell learned, a character will gain one point of maximum and current aura. Current aura is depleted by casting spells, and is naturally replenished at a rate of two points per turn. Current aura will increase until it reaches the maximum aura level. Other ways of gaining current aura may be found as the mage researches his craft.

Spells are rated on the amount of aura which the casting mage must possess. Minor spells may demand only one point of aura to cast. Powerful spells may require a current aura level of ten or higher.

### Magician status

Characters receive a rating in the turn reports based on their magical abiliity:

    maximum aura            label
    ------------            -----
       6-10                 conjurer
      11-15                 mage
      16-20                 wizard
      21-30                 sorcerer
      30+                   ??

 

For example:

    Osswid the Brave [5639], wizard, with three workers

 

The Basic magic [800] spell Appear common [803] allows magicians to prevent this label from displaying.

### Required Items

Many spells will require the magician to possess a rare or obscure item in order for the cast to succeed. Many of these items exist, of interest chiefly to sorcerers. Roots from plants found only in dense forests, bat's wings, and a dark blue powder which produces a brilliant cobalt flame when burned are only a few. Usually the required item is consumed by the attempt to cast the spell.

### Study of Spells

Magic is divided into six schools of study:

    num  name                       time to learn
    ---  ----                       -------------
    800  Magic                      four weeks, 1 NP req'd
    820  Weather magic              five weeks, 1 NP req'd
    840  Scrying                    five weeks, 1 NP req'd
    860  Gatecraft                  five weeks, 1 NP req'd
    880  Artifact construction      six weeks, 1 NP req'd
    900  Necromancy                 six weeks, 2 NP req'd

 

Magical spells are simply sub-skills of one of the magical skill categories.

An aspiring mage will issue the `study` order to learn the basics of a particular school of magic, known as the category skill. For example, one wishing to pursue knowledge of Magic [800] would order:

    study 800

 

Once the category skill has been learned, the mage will receive a lore sheet listing some of the known spells of that school. The magician may then attempt to learn these spells through study.

A character must know the category skill for a school of magic before a spell in that school can be known.

Only some of the spells in the each school are commonly known. The more rare, obscure or powerful spells will need to be discovered via `research` or by finding magic scrolls describing them.

Knowledge of individual spells in a school of magic is not possible without having learned the category skill.

### Schools of Magic

*Magic [800]*

The most common and well-known of the magical schools, Magic nonetheless has many useful and powerful spells.

Since most cities offer instruction in Magic, and several useful spells may be learned quickly, apprentice mages often begin their studies here.

*Weather magic [820]*

The study of spells to control the elements. Advanced weather magicians are said to research forgotten elf-lore in the hopes of finding tools to battle evil.

Weather magic is taught in Nimbus, Stratos and Aerovia, the three cities of the Cloudlands and randomly in non safe haven cities.

*Scrying [840]*

Scrying is the study of magical far-seeing, the ability to view images or learn information at a distance.

Scrying is taught in the Faery Cities, and randomly in non safe haven cities.

*Gatecraft [860]*

The study of ancient portals of teleportation which long ago connected distant regions of the land.

Knowledge of this art is taught in all safe haven cities, and in random non safe haven cities.

*Artifact construction [880]*

The realm of very advanced sorcerers, artifact construction concerns the making of physical items to focus, amplify or otherwise enhance a mage's power. Most spells in this school require the magician to possess high levels of aura.

Knowledge of this art is taught in random non safe haven cities and in some cities in Hades.

*Necromancy [900]*

The darkest of the Dark Arts, necromancy involves trafficking with undead or demon spirits to gain knowledge. Hated by civilized men, feared by other magicians, the necromancer seeks power and domination over the physical world.  

 

Necromancy is taught in the City of the Dead, in Hades and randomly in non safe haven cities.

 

Monsters may sometimes be found guarding ancient books which teach the rare magical skill categories.

## Religion

Learning the skill category Religion [750] labels the character as a priest. Religion [750] requires 1 Noble Point and five weeks to learn. Religion [750] may be studied in any temple. Religion is not taught by cities. Skills in the Religion category are known as prayers.

Temples yield 100 gold in offerings per month to their owner, if the owner is a priest. Temples may be built anywhere, except inside another building.

Research into Religion [750] must be performed in a temple rather than a tower.

## Battles

Combat in Olympia occurs when one stack attacks another stack.

A unit will be defended by any other characters it is stacked with. Thus, a unit which is part of a stack cannot be attacked alone. No matter which unit is specified in the `attack` order, the entire stack containing the unit will respond.

Only the leader of a stack may initiate combat. If another unit in a stack wants to issue `attack`, it must first drop out of the stack.

Combat involves nobles, who possess strong heroic fighting abilities, and fighters who may accompany them. Fighters include soldiers, pikemen, swordsmen, knights, elite guard, crossbowmen, archers, and elite archers.

### Resolution of battle

Each kind of fighter has an attack and defense rating:

                     attack    defense    missile
                    +----------------------------
    peasant         |   1         1
    worker          |   1         1
    sailor          |   1         1
    soldier         |   5         5
    pikeman         |   5        30
    swordsman       |  15        15
    pirate          |   5 (15)    5 (15)
    knight          |  45 (20)   45 (20)
    elite guard     |  90 (65)   90 (65)
    crossbowman     |   1         1         25
    archer          |   5         5         50
    elite archer    |  10        10         75

 

- Numbers in parenthesis are for shipboard combat.
- Nobles have strong innate heroic abilities, and so even an unarmed, untrained noble is a formidable opponent on the battlefield.
- A noble is rated (attack=80, defense=80, missile=0).
- A blessed soldier fights with the same attack and defense values as a regular soldier, but has a 50% chance of surviving a hit.
- A pirate fights (15,15,0) on a ship, but only (5,5,0) on land.
- Knights and elite guard receive a -25 penalty to both their attack and defense ratings when fighting on a ship or in a swamp province.
- Peasants, workers and sailors only fight when attacked, i.e. when they are members of a party which is the target of an attack.

Combat is resolved as follows:

1.  A random man is chosen from the attacking side to hit a random target from the defender.

2.  A random defender is similarly chosen to hit a random target from the attacking side.

3.  Repeat, alternating sides until the smaller side has had as many chances to hit as it has attackers.

4.  The larger side then gets N hits in a row, where N is the difference between the number of attackers in the larger side and the smaller side.

5.  Repeat until a side breaks. A side breaks when its total offensive plus defensive value falls by 50%.

    For instance, a noble plus two pikemen has an offensive value of (80+80) + (5+30) + (5+30) = 230 with a break point of 115.

    A noble plus two knights has an offensive value of (80+80) + (45+45) + (45+45) = 340 with a break point of 170.

    Therefore, a noble with two pikemen will continue to fight if the pikemen are killed. However, a noble with two knights will be declared the loser if both knights are killed.

    The offensive value used is either the attack or the effective missile rating, whichever is higher.

The chance that the attacker will score a hit against the target is:

    A = attacker's attack rating
    B = target's defense rating

    A / (A + B)

 

For example:

    A(attack=90) vs B(defense=45)     A has a 2/3 chance of 
    killing B
    A(attack=90) vs B(defense=90)     A has a 1/2 chance of killing B

 

If the attacker scores a hit against a noble, the noble will receive a random wound of 1-100 health points. Note that there is a 1% chance that a perfectly healthy noble will be instantly killed, and a greater chance that a previously wounded noble will die.

Wounded nobles do not continue fighting, even if their wounds are minor.

If a hit is scored against a fighter (soldiers, pikemen, archers, etc.), the fighter is killed. However, a blessed soldier has a 50% chance of surviving a hit.

A man successfully attacking a building or a ship will cause one point of damage to the structure. A siege engine attacking a building will cause 5-10 points of damage.

The winner in combat will attempt to take prisoners. The chance that a given defeated unit will be taken prisoner is proportional to the sizes of the remaining forces.

(Taking prisoners in battle and claiming loot requires many soldiers to run after the fleeing enemy. Thus, the chance of success is based on numerical advantage rather than combat skill.)

    1:1     25%
    2:1     50%
    3:1     75%

 

If the winning side outnumbers the defeated force by 2:1, there is a 50% chance that a given defeated unit will be captured. Defeated units which are not captured retreat from battle. If they are occupying a building or are located in a city, they may flee into the outlying province.

The victor always has at least a 25% chance of taking a unit prisoner, but no better than a 75% chance.

The victor will claim the defender's position in the location list if it is better, or will move into the defender's structure, ejecting the losing force. (The attacker may specify a flag to the `attack` order to inhibit this behavior.)

Prisoners are stripped of their belongings by the victor, including any men accompanying the prisoners, such as workers or peasants. The stack leader of the winning force receives all of the loot from the battle.

When a noble is taken prisoner (including via SURRENDER), a portion of the prisoner's items will always be lost. If a unique item is lost, it will have to have found by exploration of the province. <span id="Front and rear"></span>

### Front and rear

Units may issue the `behind` command to declare whether they will line the front or the rear in battle. Rear units do not become targets for the enemy until all of the units in the rows in front of them have been killed. Only missile fighters, such as archers, may attack from the rear.

The leader of each stack (the top-most unit in the stack) will be the last unit to receive hits, regardless of its `behind` status.

### Missile Attacks

Units in the rear may attack with their missile rating. Units in front attack with either their missile rating, or their attack rating, whichever is higher.

Thus, a noble with rating (attack=80, defense=80, missile=40) will do an attack of 40 when in the rear, and an attack of 80 when in front.

Weather effects on battle are as follows:

- Rain or wind cut the missile rating of archers or elite archers in half, and the missile rating of crossbowmen to 1/4 normal.
- Fog cuts the missile rating of all figures to 1/4 normal.

If a fighter's missile rating is zero, it can not attack from the rear. In this case, the fighter will use its attack rating like a regular soldier.

### Fortifications

Structures which may aid fighters in battle are rated for their defensive bonus.

    Castle Imperius [gx56], castle, defense 25

 

During battle, the fortification rating is added to the defense number for the men who fit inside the structure.

    structure               men protected
    ---------               -------------
    Castle                  first 500
    Tower                   first 100
    Galley or roundship     first 50
    Other structures        first 50

 

Attacking fighters may randomly select the structure instead of an enemy fighter. The attack is resolved in the same way as for two fighters, using the attacker's attack rating and the structure's defense rating. If the attacker is successful the structure's defense rating will be lowered by one point.

Once the defense rating reaches zero, further hits will cause the building to become damaged. A fully damaged building (100% damage) will collapse, ejecting its occupants.

Siege engines always select the structure as a target.

    engine          attack  defense   missile
    ------          ------  -------   -------
    catapult          25      200       25
    battering ram     30      250
    siege tower       30      250

 

Siege engines do 5-10 points of damage to the structure per hit.

Siege engines are not used in combat at sea.

### Item bonuses

Nobles may possess items which grant attack, defense or missile bonuses in combat. Only one item may be wielded for each category. If the noble possesses multiple items with bonuses in the same area, the item with the largest bonus will be chosen.

For example, suppose Osswid had the following items:

    Shield of Achilles [fx78]       +25 defense
    Sword of Death [gl23]           +10 attack
    Mithril Axe [wt29]              +15 attack
    Magic javelin [ht02]            +5 missile

 

Osswid would wield the javelin and axe, and wear the shield. He would not use the Sword of Death.

Nobles automatically use items with combat bonuses in battle. No special orders are needed to wield them.

## Prisoners

Characters may become prisoners by losing to an enemy in battle, or by using the surrender order (see orders section for description).

Since prisoners are unable to report where they are and what they are seeing, they do not contribute to the turn report of their faction. The player's turn report will show that that a unit is being held prisoner, but little else.

Prisoners will not execute any orders while they are in captivity. Queued orders will remain pending, but none will be processed.

Prisoners when spotted appear as stacked units, marked with the \``prisoner`' string:

    Seen here:
       Kosar the Indefectible [2022], with six peasants, one archer,
       two soldiers, accompanied by:
          Alion Krysaka [2785], prisoner

 

Unstacking a prisoner sets them free. Kosar could free Alion by ordering \``unstack`` 2785`'.

Prisoners may be transferred between units with the `give` command. Kosar could transfer Alion to Osswid [501] by ordering:

    give 501 2785

 

------------------------------------------------------------------------

**Prisoner escapes**

Prisoners are always on the lookout for ways to escape. Units holding prisoners can reduce their chances by remaining inside a structure, not transferring prisoners with give, and not traveling with prisoners. 

Each week (four times each game turn), a prisoner being held by a unit which is outside of a building has a 2% chance of escaping. A prisoner being held by a unit which is inside a structure, such as a castle, tower, inn or ship, has a 1% chance of escaping.Each time a prisoner is transferred with the give order, there is a 2% chance of a escape. Also, each time a unit holding a prisoner engages in travel which takes longer than one day, there is a 2% chance that the prisoner will be able to get free. Thus, short movement, such as entering or exiting a building, will not give the prisoners additional opportunities for escape, but traveling between provinces with prisoners will.Prisoners inside building or sub-locations will flee out into the surrounding location upon gaining their freedom. Escaped prisoners on ships will leap over the side and swim to a nearby shore.

## Permissions and declared attitudes

Commands dealing with permissions:

- `admit`
- `hostile`
- `neutral`
- `defend`
- `default`

Attitudes can be declared by or for either specific units, or an entire faction. For instance, player [613] could declare a permission or attitude for player [555], or a specific attitude for individual units within player 555's faction.

Declaring a permission for a player works so long as the player's units are not concealing their faction identity with Conceal faction [635], a subskill of Stealth [630].

### Allowing entrance and stacking

By default, a unit may not stack with a character belonging to another faction. A unit is also denied entry to a building or ship controlled by another player.

Players may allow units from other factions to stack with them or enter buildings or ships they control with the `admit` order.

### Combat attitudes

A unit may have one of four combat attitudes to another unit:

`hostile`

Attack on sight.

 

`defend`

Defend other unit if attacked.

 

`neutral`

Do nothing if other unit is attacked.

 

`default`

Neutral to units in other factions; Defend units in the same faction unless either one is concealing its lord.

 

Every character, and player faction entity, keeps three lists of units or other factions towards which it has declared attitudes. A unit is either on the `hostile`, `defend`, or `neutral` list. If a unit does not appear on any of the three lists, it has attitude `default`.

Example:

    player          778                     816
                      hostile 816

    units           4205                    6499
                    4600                    6530, concealing lord
                                            6599

 

Player 778 has declared player 816 hostile. One of 816's characters is concealing its lord.

If 4205 or 4600 run into unit 6499, they will attack it on sight. However, since 6530 is hiding its affiliation with 816, it will not be attacked on sight.

If 6499 is attacked and both 6530 and 6599 are present, 6599 will aid in the defense, but 6530 will not, because that might give away its affiliation.

If player 816 wanted 6530 to defend the faction's units anyway, either 816 or 6530 should issue the order \``defend`` 816`'. This would override the default attitude of units in the faction to one another.

Attitude toward units is considered before attitude toward the unit's faction. Thus, one may declare a faction hostile, but exclude certain units within the faction by specifically declaring them neutral.

A unit must be the top-most character in its stack to aid in defense. If a top-most unit joins a combat because of `defend`, it will bring its entire stack along, even if the other members of the stack have not declared a `defend` attitude.

Defenders only help when units are attacked, not when they initiate attacks. For example, if `A` has declared \``defend`` ``B`', and `B` attacks `C`, `A` will not help `B`, even if `B` loses the battle.

Characters declared `defend` to units which are guarding a province against pillaging will aid the guards if they are attacked, either explicitly with `attack`, or implicitly via \``pillage`` 1`'.

Units which joined a combat because of a `defend` declaration are shown with the qualification \``ally`' in the combat report.

## Buildings

The following structures may be built by characters with the Construction [680] skill:

    kind        effort    material   skill         where
    ----        ------    --------   -----         -----
    inn            300     75 wood    680          province or city
    mine           500     25 wood    680 or 720   mountain or rocky hill
    temple       1,000    50 stone    680          anywhere
    tower        2,000   100 stone    680          anywhere
    castle      10,000   500 stone    680          province or city

 

The term \``anywhere`' means a province, city, or other sublocation. Buildings may not be built inside other buildings, with the exception of towers. Up to six towers may be built inside a castle.

Effort is in worker-days.

The builder must know the required skill, have at least three workers, and possess at least one-fifth of the necessary construction materials in his inventory.

To start building, unstack from beneath other characters and issue one of the following `build` orders:

- `build inn` "`name of inn`"
- `build mine` "`name of mine`"
- `build temple` "`name of temple`"
- `build tower` "`name of tower`"
- `build castle` "`name of castle`"

For example, a character who wanted to build a tower would need to know Construction [680], have at least three workers [11], and possess at least 20 stone [78].

The construction materials will immediately be deducted from the builder's inventory and put to use. The builder and his workers will be placed inside the new structure.

The remaining construction materials are paid as work on the structure progresses. A second fifth is required as the building becomes 20% complete, the third fifth at 40%, etc. Construction halts if the builder runs out of materials.

The building is completed when the required number of worker-days is invested in construction.

To resume construction of a partially completed building, first enter the structure, then issue the appropriate build order (such as `build inn`, `build mine`, etc.)

Mines may only be built in mountain provinces or on rocky hills. Only one mine per location is allowed. Mines may also be built by characters with the Mining [720] skill.

Only one castle may be built in each province. The castle may be built either in the province or a city, if the province has one. Castles may not be built in sub-locations other than cities.

## Inns

Inns must be built in provinces or cities. A good site for an inn is just outside a city, as the inn will benefit from the patronage of many travelers.

The restive environment inns provide help wounded nobles to heal. Nobles located in an inn have an 10% increased chance of fighting off infection each week.

Inns generate income each month from the visitors who stop for a meal and a pint of ale, or spend in the night in one of the inn's rooms. (Nobles who enter are not directly charged. Other patrons are anonymous; the only indication of their presence is the income the inn generates.)

Inns generate between 50 and 75 gold per month. Looting and pillaging scares away customers, and so lowers the inn's income. If more than one inn is built in a province, the profits are split between them.

Income is paid to the owner of the inn at the end of each month.

## Control of buildings and ships

The owner of a building or a ship is the first character shown inside. This will be the first character to have entered the building or ship, unless another character attacked the previous owner and taken position at the head of the list.

For example, assume the following characters are on a ship, the HMS Pinafore [4000]:

    Candide the Captain [1269], with ten sailors
    Osswid [5499], accompanied by:
       Feasel the Wicked [1109]

 

Since Candide is first in the ship's character list, he is the owner of the ship. Only Candide may issue sail orders, or change the ship's name. If Candide were to leave, Osswid would become the new captain of the ship.

The owner of a building or ship may determine who may enter with the `admit` order. The default is to refuse entry to units from other factions.

If the first character inside a building leaves, he is no longer the owner.

A building is a castle, tower, inn, temple or mine.

## Mining

A mine is a deep shaft or tunnel which allows workers to extract valuable resources from the earth, such as iron and gold. At most one mine may be built in each mountain province or rocky hill.

A new mine has an initial depth of one. The mine shaft becomes deeper as characters use it to obtain natural resources. The shaft will become one level deeper for every three uses of a mining extraction skill.

As the depth of the shaft increases, the mix of resources obtainable changes. Iron is usually found nearest the surface. As one proceeds deeper, gold may be found in higher quantities. Other rare elements may be found by going deeper still.

The deeper a mine becomes, the more frequently cave-ins or other accidents will occur. With each accident, the mine's damage percentage will rise. If not attended with `repair`, the mine will eventually collapse. Once a mine collapses, it remains in the province for one game year (eight game months). Characters may not enter or use a collapsed mine. After the year has passed, the collapsed mine will vanish, and a new mine may be built in the location.

## Opium

Opium is produced in swamp regions, and consumed by markets in desert, plain, forest, and mountain provinces. All markets (except those in swamp regions) have some level of opium demand. However, this demand will not be visible in the market report at low levels.

Satisfying opium demand in a market will cause the next month's demand to be higher. As peasants become addicted to opium, the increased demand will be shown in the location's market report. If no opium is sold to a market, the demand will fall.

Opium adversely affects the city's tax base. The more opium the market buys, the more tax revenues will be reduced.

## Pillaging

Gangs of ten or more fighters may use the `pillage` command to seize loot from a province or a city. Pillaging siezes the tax base for a location, leaving none available for taxation.

Pillaging has a harmful effect on the future tax revenue of the location. The more a location has been pillaged, the lower its tax base. Provinces and cities take four months to recover from each pillaging. For example, a city which was pillaged for five months in a row would take twenty months from the first pillaging to return to its normal tax base.

Pillagers must first defeat any units guarding the province, including any province garrisons.

## Province control and ownership

Each province and city generates an amount of gold each month. This gold is known as its *tax base*. The size of the tax base is determined by the civilization level of the province. This gold may be pillaged, taxed by a castle, or taxed by a garrison. It does not accumulate if left uncollected at the end of a turn.

A castle automatically collects all of the gold in its province.

The \``garrison`` ``castle`' order installs a group of at least ten soldiers in a province to claim it and guard against pillaging. Garrisons must be bound to a castle.

A garrison pays maintenance for its members from the province tax base, then forwards 1/2 of the remaining gold to its castle.

Garrisons with fewer than 10 fighting men will pay maintenance for themselves, but will not be able to forward tax to a castle, guard against pillaging or obey decree orders.

The castle owner gains status from the number of provinces under control:

    provinces       rank
    ---------       ----
       1-5          lord
       6-12         knight
      13-25         baron
      25-37         count
      38-50         earl
      51-63         marquess
       64+          duke
      region        king     (region must have at least 15 provinces)

 

A noble may `pledge` to another noble, granting status and control of owned provinces. The status of a noble who pledges is the smaller of the original status or one below the rank of the pledge target:

    new status = min(original status, one below rank of pledge 
    target)

 

Control of a province allows one to change its name or the name of any of its sublocations, take items from the garrison, and issue decrees to watch for certain units, or to attack specified units on sight.

The castle continues to receive the income from garrisoned provinces, even if the castle's owner is pledged to another noble.

Every noble in the pledge chain shares control of the garrisoned provinces. In other words, a castle owner may pledge to a noble, who in turn may pledge to a third noble, etc. Thus a province may have any number of rulers.

Visitors to a province are informed of the castle to which the garrison is bound, and the top-most ruler in the pledge chain (which may simply be the owner of the castle):

    Province controlled by Amber Keep [0909], castle, in Forest 
    [cj12]
    Ruled by Erekosse [5210], baron
    ...
    Seen here:
      Garrison [780], garrison, on guard, with ten soldiers

 

### Tax base

Each province generates a tax base each month. The amount of gold fed into the tax base is determined by the civilization level of the province:

    civ level       tax gold
    ---------       --------
    wilderness         50
    civ-1             100
    civ-2             150
    civ-3             200
    civ-4             250
    civ-5             300
    civ-6             350
    civ-7             400

 

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

Garrisons are groups of men who are stationed in provinces to protect them, and collect taxes in the name of a castle. Garrisons must be created with the `garrison` order, and must be bound to a castle located in the same region.

For example, suppose that the region Lesser Atnos had 20 provinces. One of these provinces contains Amber Keep [0909]. A garrison bound to Amber Keep could be stationed in each of the 20 provinces (including the province containing the castle itself).

Continuing the example, garrison units not in the Lesser Atnos region could not be bound to Amber Keep. The castle a garrison is bound to must be in the same region.

Garrisons can be bound to any castle in the region. If Lesser Atnos had two castles, some of the garrisons could be bound to one, and the rest to the second castle.

### More about garrisons

A garrison in a province containing a castle must be bound to that castle.

A garrison may only be installed in a province adjoining a province which already contains a garrison bound to the same castle, or the province the castle is in.

Garrisons are established with the \``garrison`` ``castle`' order. Ten soldiers are required to create a garrison. The `garrison` order must be issue at the outer level of a province; one can't establish a garrison while inside a city, building or other sublocation.

The garrison pays the maintenance cost of its men directly from the tax base of the province. Half of the remaining tax base is forwarded to the castle the garrison is bound to.

For example, a garrison of ten soldiers would require 20 gold per month to support. This would leave 280 gold remaining in a typical province. 50% of this, or 140 gold, would be forwarded to the owner of the garrison's castle.

Example:

    > garrison cy09
    Installed Garrison [780], garrison, on guard, with ten soldiers

 

Visitors to this province would see:

    Province controlled by Amber Keep [0909], castle, in Forest 
    [cj12]
    Ruled by Erekosse [5210], baron

 

Note that Erekosse may be located inside the castle, or the castle's owner may have pledged service to him, in which case Erekosse could be anywhere.

### Garrison reports

Garrisons do not provide full location reports to their owners. They do notice any resource depletion activity, such as timber cutting or mining, as well as any large or unusual parties which enter their province. This includes any stack of five units or more, any party of 20 or more men, and most monsters or wild beasts.

Garrisons do not monitor activity in hidden locations, even if the players who rule over the garrisons have discovered the hidden locations.

The \``decree watch`` ``who`' order may be given by a ruler to instruct all garrisons to watch for a particular unit. This is useful for locating individuals who would otherwise go unnoticed by the garrisons.

### Referring to garrisons

Since a province may only have one garrison, garrisons may be referred to without knowing their entity number. The keyword `garrison` will match the province's garrison, if there is one.

Examples:

    give garrison 12 5
    attack garrison

 

### Status

The number of provinces a noble controls determines his status or rank:

    provinces       rank
    ---------       ----
       1-5          lord
       6-12         knight
      13-25         baron
      25-37         count
      38-50         earl
      51-63         marquess
       64+          duke
      region        king     (region must have at least 15 provinces)

 

In addition, if a character has control over every province in a region, and the region contains at least 15 provinces, then the character is given the rank of king.

Provinces may be directly owned, if the noble is the owner of a castle, or indirectly, through other pledged nobles.

### Pledging land

A noble may `pledge` his lands to another noble. This grants the pledge target status by increasing the number provinces he may rule over.

For example, suppose there are two castle owners, Osswid and Feasel. Osswid has garrisoned six provinces, and Feasel has three. Osswid is therefore a baron, and Feasel is a lord.

If Feasel and Osswid both pledge to Candide, Candide would attain the rank of Count. Osswid and Feasel would remain at the same rank in this example.

Candide would receive garrison reports for all provinces which Osswid and Feasel control. He would have the same privileges in the controlled provinces: he could take items from the garrisons, alter the names of the provinces or their sublocations, and issue watch and hostile decrees.

However, the income generated by the provinces would continue to be forwarded to the castles. No extra income goes to the pledge target.

### Status after pledging

The status of a noble `A` who is pledged to another noble `B` will be either `A`'s original status, as determined by how many provinces he controls, or one rank below `B`, whichever is lower.

For example, a noble with 5 provinces who pledges to a king will remain a baron. However, if pledged to another baron, the noble's rank would fall to lord.

## Relics

Relics are unique item artifacts which are introduced into the game via quests with monsters. All relics except the Throne return to the netherworld after use or some delay to be given out to a new adventurer via QUEST.

- Imperial Throne [401]

  Long ago, the emperor of Olympia sat on his throne in the emperor's palace, high amidst Mt. Olympus. It is said that whoever rebuilds the famed castle on Mt. Olympus, and sits upon the throne, will be titled Emperor of Olympia.

- Crown of Prosperity [402]

  The Crown of Prosperity was once worn by the most prosperous mortal to ever rule in Olympia, King Damar. Damar now wears his crown in the underworld and reflects on his past adventures of long ago, before the first Great Ending swept away all that he knew, and carried him from his beloved city of Kircarth to the land of the dead.

  Sometimes nobles are able to acquire the Crown and hold it for a time. The crown infuses whatever province it ends each turn in with a measure of prosperity and economic health, equivalent to a +2 increase in the province's civilization level.

  However, this prosperity does not last forever, for King Damar's ghostly hand invariably will reach out from the underworld to reclaim his relic. The Crown can be expected to return to its rightful owner 12-24 turns after its appearance in the mortal world.

- Skull of Bastrestric [403]

  The most feared of the ancients was a wizard of terrible power known as Bastrestric ther Archymonaged. Bastrestric routinely incinerated his foes (or anyone who offended him) with bursts of raw aura energy directed from his black tower in the castle built on Mt. Olympus.

  Though he was the most powerful living mage in the known world, and by any measure a fearsome, unholy force, Bastrestric yearned for ever greater abilities. He felt increasingly constrained by the limits of his mortal body, and thus in time resolved to abandon it. BtA's spirit jumped free of his body, plunging directly into the aura rivers which bind together the deepest structures of the very world itself. Only a scorched, empty body was left behind in his tower.

  BtA's spirit has passed on, but the remants of his mortal body continue to radiate intense power.

  Use of BtA's skull (USE 403) causes an intense aura burst which a mage will attempt to absorb. If successful, the mage will gain a 50-75 boost to current aura (within the limit of 5 times the mage's maximum aura).

  There is a 25% chance that the mage will be killed by use of the skull. Non-mages who use the skull will be instantly killed.

  BtA's skull will vanish with use, or within 10-20 turns after appearing in the mortal world.

## Ships

Players may build two kinds of ships in Olympia: galleys and roundships.

A galley, also known as a warship, is a slender rowed vessel. Galleys require 14 sailors as oarsmen for travel, and may carry up to 5,000 units of cargo.

The roundship, also known as the merchantman, is a deep, wide sailing ship, usually with one or two masts and steered with great, oar-like paddles. The large cargo space makes it well suited for trade and extended ocean travel.

With favorable winds, roundships will make better time than galleys, even when fully loaded. Roundships require a crew of eight sailors travel, and may carry up to 25,000 units of cargo.

The Shipcraft [600] skill category allows nobles to train sailors, and has sub-skills for building and sailing ships.

    ship            capacity        sailors
    ----            --------        -------
    galley             5,000           14
    roundship         25,000            8

 

Ships may be damaged while sailing by storms, or by submerged rocks in coastal waters. Damaged ships may be repaired with the `repair` command. Repairing a damaged ship requires one unit of pitch [261].

The capacity of a damaged ship is a reduced in proportion to the amount of damage. A galley with 10% damage may only carry 4,500 units of cargo.

### How to sail

The Sailing [601] subskill of Shipcraft [600] is required to pilot a ship. Shipcraft may be learned in any port city. For information about piloting galleys and roundships, see the sail order.

### Making ships

Construction of ships requires a character to have the Shipbuilding [602] skill, a subskill of Shipcraft [600].

    ship               effort               material
    ----               ------               --------
    galley          250 worker-days          50 wood
    roundship       500 worker-days         100 wood

 

To begin construction of a ship, the shipbuilder should unstack from beneath other characters and issue one of the following `build` orders:

- `build galley` `name of galley`
- `build roundship` `name of roundship`

One-fifth of the lumber will immediately be deducted from the shipbuilder's inventory and put to use building the new ship. The shipbuilder and his workers will be placed inside the new ship.

At least three workers are needed to begin construction of a ship.

Until the ship is completed, it will be shown as \``in progress`':

    Ships docked at port:
       HMS Pinafore [1111], galley-in-progress, 28% completed, owner:
          Osswid the Constructor [5499], with five workers

 

The remaining construction materials are deducted from the builder's inventory as work on the ship progresses. A second fifth of the lumber is required when the ship becomes 20% complete, a third fifth when the ship becomes 40% complete, etc. Construction halts if the builder runs out of materials

For example, a noble who wanted to build a galley would need to know Shipbuilding [602], have at least three workers, and start with at least 10 wood [77]. (40 more wood is required to bring the galley to completion).

As soon as the required number of worker-days has been invested in construction, the ship will be christened and declared seaworthy.

To resume construction of a partially completed ship, first enter the ship, then issue the either `build galley` or `build roundship`.

Ships may only be built in port cities.

### Operating a ferry

There are several commands useful for operating a commercial ferry:

`fee`

Sets the fee passengers will be charged

 

`board`

Passengers use this to pay their fee and board a ferry

 

`unload`

Unload passengers once the destination is reached

 

`ferry`

Signals passengers waiting in port that they may board

 

The captain of a ferry must issue \``fee`` ``gold`' command to set a fee which will be charged to passengers wishing to `board` his ship. The fee is expressed as how many gold pieces per 100 weight of the passenger's stack will be charged.

For instance, if the captain wanted to charge 1/2 gold per unit weight of the passengers stack, he would issue \``fee`` 50`'. (a 50 gold fee for every 100 weight).

The fee is a property of the captain, not the ship.

A character may clear the fee with the \``fee`` 0`' order. If no fee is set, then the ship is not considered to be operating as a ferry, and characters are may not use the `board` order to enter the ship.

Passengers issue \``board`` ``ship`' to board a ferry. The order will fail if the ship is not present, or if it is not operating as a ferry (the captain of the ship has no fee set). `board` will cause the character to pay the captain the required boarding fee, then move the character's stack onto the ship.

The captain shouldn't issue an `admit` order to let characters on board who will be paying to take the ferry. Otherwise they could board the ship with `move` instead of `board`, bypassing the ferry fee.

### Ferry synchronization

Suppose now that the captain has a ship, has set a fee, announced his service in the *Olympia Times* and with `post`, and is now ready to ferry passengers. What should he tell them? How will the synchronization work?

Passengers should travel to the port the ferry will be arriving at and issue \``wait ferry`` ``ship`'.

When the ship arrives at port, the captain should order `unload` to eject his current load of passengers, then `ferry` to signal any passengers waiting in port that they may now board.

Note that passengers should not use `wait ship` to wait for the ferry. Otherwise, they will attempt to `board` as soon as the ship reaches the port. The captain's `unload` order may not have executed at this point. In this case, `board` orders might fail because there won't be any room to enter the ship until the existing load of passengers has disembarked.

Example:

    Posted by Captain McCook [3402]:
    "Captain McCook's ferry to Drassa departs each
     Sunsear on the 15th.  The fee is 1 gold/wt.
     Issue WAIT FERRY 1234 then BOARD 1234 for
     ferry service.   Arrr!"

 

Captain McCook's orders:

    > fee 100                 # 1 gold/wt. is our fee
    > sail ...                # arrival at port
    > unload                  # unload current passengers
    > wait day 15             # wait until stated time of departure
    > ferry                   # sound our horn
    > sail ...                # on to the next port

 

Passengers wishing to travel on McCook's ferry:

    > wait ferry 1234             # [1234] is McCook's ship
    > board 1234                  # pay our gold and embark
    > wait loc destination        # Captain McCook will 
    unload us at the
                                  # end of the journey, so wait until we
                                  # find ourselves in the destination city.

 

## Summary of important tables

### Item combat values

                         name    swamp man-like    beast  
    fighter
                         ----    ----- --------    -----  -------
                 peasant [10]     no       yes      no    (1,1,0)
                  worker [11]     no       yes      no    (1,1,0)
                 soldier [12]     no       yes      no    (5,5,0)
                  archer [13]     no       yes      no    (5,5,50)
                  knight [14]     no       yes      no    (45,45,0)
             elite guard [15]     no       yes      no    (90,90,0)
                 pikeman [16]     no       yes      no    (5,30,0)
         blessed soldier [17]     no       yes      no    (5,5,0)
           ghost warrior [18]     no       yes      no    (0,0,0)
                  sailor [19]     no       yes      no    (1,1,0)
               swordsman [20]     no       yes      no    (15,15,0)
             crossbowman [21]     no       yes      no    (1,1,25)
            elite archer [22]     no       yes      no    (10,10,75)
           angry peasant [23]     no       yes      no    (2,1,0)
                  pirate [24]     no       yes      no    (5,5,0)
                     elf [25]     no       yes      no    (50,50,100)
                  spirit [26]     no       yes      no    (50,50,0)
                  undead [31]     no       yes      no    (10,10,0)
                  savage [32]     no       yes      no    (1,1,0)
                skeleton [33]     no       yes      no    (6,6,0)
               barbarian [34]     no       yes      no    (2,1,0)
              wild horse [51]     yes      no       no     -
            riding horse [52]     yes      no       no     -
                warmount [53]     yes      no       no     -
            winged horse [54]     yes      no       no     -
                  nazgul [55]     yes      no       yes   (80,80,0)
           battering ram [60]     no       no       no    (30,250,0)
                catapult [61]     no       no       no    (25,200,25)
             siege tower [62]     no       no       no    (30,250,0)
                      ox [76]     yes      no       no     -
               ratspider [81]     no       no       yes   (5,5,0)
                centaur [271]     yes      no       yes   (30,30,0)
               minotaur [272]     yes      no       yes   (30,5,0)
           giant spider [278]     yes      no       yes   (150,100,0)
                    rat [279]     yes      no       yes   (3,3,0)
                   lion [280]     yes      no       yes   (100,100,0)
             giant bird [281]     yes      no       yes   (200,150,0)
           giant lizard [282]     yes      no       yes   (45,45,0)
                 bandit [283]     no       yes      no    (3,3,0)
                chimera [284]     yes      no       yes   (130,130,0)
                 harpie [285]     yes      no       yes   (80,120,0)
                 dragon [286]     yes      no       yes   (500,500,250)
                    orc [287]     no       yes      yes   (20,15,0)
                 gorgon [288]     yes      no       yes   (10,20,0)
                   wolf [289]     yes      no       yes   (5,5,0)
                cyclops [291]     no       yes      yes   (25,75,0)
                  giant [292]     no       yes      yes   (75,25,0)
                  faery [293]     no       yes      no    (9,9,9)
                  hound [295]     yes      no       no    (1,1,0)

Items with a "yes" in the "swamp" column will cause their stack to move more slowly through swamps, as they need to be lead.

Items with a "yes" in the "beast" column are counted as beasts for the various Beastmastery skills.

### Weights

A -1 in a carrying capacity field means that the item will carry its own weight, but nothing additional. Capacity does not include the weight of the item itself.

    item name                     weight land ride  fly
    ---- ----                     ------ ---- ----  ---
       1 gold                          0    0    0    0
      10 peasant                     100  100    0    0
      11 worker                      100  100    0    0
      12 soldier                     100  100    0    0
      13 archer                      100  100    0    0
      14 knight                      400  100  100    0
      15 elite guard                 400  100  100    0
      16 pikeman                     100  100    0    0
      17 blessed soldier             100  100    0    0
      18 ghost warrior                 0    0    0    0
      19 sailor                      100  100    0    0
      20 swordsman                   100  100    0    0
      21 crossbowman                 100  100    0    0
      22 elite archer                100  100    0    0
      23 angry peasant               100  100    0    0
      24 pirate                      100  100    0    0
      25 elf                         100  100    0    0
      26 spirit                      100  100    0    0
      31 undead                      100  100    0    0
      32 savage                      100  100    0    0
      33 skeleton                    100  100    0    0
      34 barbarian                   100  100    0    0
      51 wild horse                1,000   -1   -1    0
      52 riding horse              1,000  150  150    0
      53 warmount                  1,000  150  150    0
      54 winged horse              1,000  150  150  150
      55 nazgul                    1,500  150  150  150
      59 flotsam                      30    0    0    0
      60 battering ram               300    0    0    0
      61 catapult                    300    0    0    0
      62 siege tower                 300    0    0    0
      63 ratspider venom               0    0    0    0
      64 lana bark                     1    0    0    0
      65 avinia leaf                   1    0    0    0
      66 spiny root                    1    0    0    0
      67 farrenstone                   1    0    0    0
      68 yew                           3    0    0    0
      69 elfstone                      1    0    0    0
      70 mallorn wood                 30    0    0    0
      71 pretus bones                  1    0    0    0
      72 longbow                       3    0    0    0
      73 plate armor                   5    0    0    0
      74 longsword                     5    0    0    0
      75 pike                          5    0    0    0
      76 ox                        2,000 1,500   -1    0
      77 wood                         30    0    0    0
      78 stone                       100    0    0    0
      79 iron                         10    0    0    0
      80 leather                      30    0    0    0
      81 ratspider                     1    0    0    0
      82 mithril                       5    0    0    0
      83 gate crystal                  1    0    0    0
      84 blank scroll                  2    0    0    0
      85 crossbow                     10    0    0    0
      87 fish                          2    0    0    0
      93 opium                         1    0    0    0
      94 woven basket                  1    0    0    0
      95 clay pot                      5    0    0    0
      98 drum                          2    0    0    0
      99 hide                         50    0    0    0
     102 lead                         10    0    0    0
     261 pitch                       100    0    0    0
     271 centaur                   1,000  150  150    0
     272 minotaur                    800   -1   -1    0
     278 giant spider                500   -1   -1    0
     279 rat                          50   -1   -1    0
     280 lion                        750   -1   -1    0
     281 giant bird                1,500   -1   -1   -1
     282 giant lizard              1,000   -1   -1    0
     283 bandit                      100  100    0    0
     284 chimera                     800   -1   -1    0
     285 harpie                      500   -1   -1   -1
     286 dragon                    2,500   -1   -1   -1
     287 orc                         100  100   -1    0
     288 gorgon                      200   -1   -1    0
     289 wolf                        100   -1   -1    0
     290 crystal orb                   1    0    0    0
     291 cyclops                     200  200    0    0
     292 giant                       200  200    0    0
     293 faery                        10   10    0    0
     295 hound                        65   -1    0    0
     401 Imperial Throne               0    0    0    0

### Training tables


         +--- sailor ------- pirate
        +---- worker
       +----- crossbowman
      /
    peasant ------ soldier ------ swordsman --- knight ----- elite guard
                      \
                       +---- archer ------ elite archer
                        +--- pikeman
                         +-- blessed soldier

 

    num  kind            skill  input man       input item        
    where
    ---  -----------     -----  --------------  ----------------  ------
     11  worker           none  peasant [10]
     12  soldier           610  peasant [10]
     13  archer            615  soldier [12]    longbow [72]
     14  knight            616  swordsman [20]  warmount [53]
     15  elite guard       616  knight [14]     plate armor [73]  castle
     16  pikeman           610  soldier [12]    pike [75]
     17  blessed soldier   750  soldier [12]                      temple
     19  sailor            601  peasant [10]
     20  swordsman         616  soldier [12]    longsword [74]
     21  crossbowman       610  peasant [10]    crossbow [85]
     22  elite archer      615  archer [13]                       castle
     24  pirate            616  sailor [19]     longsword [74]    ship

 

### Combat ratings of fighters

    attack    defense    missile
                    +----------------------------
    peasant         |   1         1
    worker          |   1         1
    sailor          |   1         1
    soldier         |   5         5
    pikeman         |   5        30
    swordsman       |  15        15
    pirate          |   5 (15)    5 (15)
    knight          |  45 (20)   45 (20)
    elite guard     |  90 (65)   90 (65)
    crossbowman     |   1         1         25
    archer          |   5         5         50
    elite archer    |  10        10         75

 

- A noble is rated (attack=80, defense=80, missile=0).
- Blessed soldiers fight with the same attack and defense values as regular soldiers, but has a 50% chance of surviving a hit.
- Pirates fight (15,15,0) on a ship, but only (5,5,0) on land.
- Knights and elite guard receive a -25 penalty to both their attack and defense ratings when fighting on a ship or in a swamp province.

### Maintenance costs

    num   kind            cost
    ---   -----------     ----
     10   peasant          1
     11   worker           2
     19   sailor           2
     21   crossbowman      2
     12   soldier          2
     13   archer           3
     16   pikeman          3
     17   blessed soldier  3
     20   swordsman        3
     24   pirate           3
     14   knight           4
     22   elite archer     4
     15   elite guard      5

### Castle improvement

    level   stone   worker-days
    -----   -----   -----------
      1       50       1000
      2       60       1250
      3       70       1500
      4       80       1750
      5       90       2000
      6      100       2500

IMPROVE [days]

Runs for the specified number of days, or until the next improvement level is reached, whichever comes first.

### Castle garrisons

A castle's improvement level determines how many provinces may be garrisoned to it:

    castle
    improvement   provinces      rank
    -----------   ---------      ----
         0           1-5         lord
         1           6-12        knight
         2          13-24        baron
         3          25-37        count
         4          38-50        earl
         5          51-63        marquess
         6           64+         duke
                    region       king

(A region must have at least 15 provinces to have a king)

### Protection

    structure               men 
    protected
    ---------               -------------
    Castle                  first 500
    Tower                   first 100
    Galley or roundship     first 50
    Other structures        first 50

### Civilization

| Feature    | Contribution                |
|------------|-----------------------------|
| Safe Haven | 2                           |
| Castle     | 1.5 + improvement level / 4 |
| City       | 1                           |
| Tower      | 1                           |
| Temple     | 1                           |
| Inn        | 1                           |
| Mine       | 1                           |

### Beastmastery

Each use of BREED incurs a 20% chance of an accident which may kill one parent.

The time necessary for the BREED command to function is dependent on the type of animal produced:

    beast           days
        -----           ----
        dragon [286]             45
        giant bird [281]         28
        giant spider [278]       21
        chimera [284]            21
        nazgul [55]              14
            centaur [271]            14
        lion [280]               14
        harpie [285]             14
        all other beasts          7

### Quest

    where        what           combat      how many
    -----        ----           -------     --------
    islands      pirate [24]        (5,5,0)     5-30
             cyclops [291]      (25,75,0)   1-5
     
    caves        rat [279]      (3,3,0)     10-50
             wolf [289]     (5,5,0)     3-10
             ratspider [81]     (5,5,0)     5-20
             gorgon [288]       (10,20,0)   3-5
             orc [287]      (20,15,0)   5-20

    ruins        bandit [283]       (3,3,0)     2-10
             cyclops [291]      (25,75,0)   2-5
             minotaur [272]     (30,5,0)    3-10
             centaur [271]      (30,30,0)   3-10
             giant lizard [282] (45,45,0)   1-3

    battlefields     skeleton [33]      (6,6,0)     10-100
             spirit [26]        (50,50,0)   5-50
             giant [292]        (75,25,0)   3-10
             nazgul [55]        (80,80,0)   5-20

    graveyards   undead [31]        (10,10,0)   10-100
             harpie [285]       (80,120,0)  3-10
             giant spider [278] (150,100,0) 3-10
             giant bird [281]   (200,150,0) 1-3

    lairs
             lion [280]     (100,100,0) 3-8
             chimera [284]      (130,130,0) 2-10
             dragon [286]       (500,500,250)   1

    enchanted
    forests,
    faery hills  faery [293]        (9,9,9)     5-20
             elf [25]       (50,50,100) 5-20

    pits & bogs  rat [279]      (3,3,0)     5-25
             gorgon [288]       (10,20,0)   3-7
             cyclops [291]      (25,75,0)   2-3
             minotaur [272]     (30,5,0)    1-5

    sand pits    giant lizard [282] (45,45,0)   15-30

    circles of
    stone        skeleton [33]      (6,6,0)     3-15
             gorgon [288]       (10,20,0)   3-15
             orc [287]      (20,15,0)   3-15
             cyclops [291]      (25,75,0)   3-15
             minotaur [272]     (30,5,0)    3-15
             centaur [271]      (30,30,0)   3-15
             spirit [26]        (50,50,0)   3-15
             giant [292]        (75,25,0)   3-15
             nazgul [55]        (80,80,0)   3-15
             harpie [285]       (80,120,0)  3-15
             chimera [284]      (130,130,0) 3-15

    tunnel chamber,  rat [279]      (3,3,0)     10-50
    levels 1-6   ratspider [81]     (5,5,0)     5-20
             gorgon [288]       (10,20,0)   3-6
             orc [287]      (20,15,0)   5-20

    tunnel chamber,  cyclops [291]      (25,75,0)   5-10
    levels 5-6   minotaur [272]     (30,5,0)    5-15
             giant lizard [282] (45,45,0)   5-15
             giant spider [278] (150,100,0) 5-15

### Experience

    use             level
    -----           -----
     0-4            apprentice
     5-11           journeyman
    12-20           adept
    21-34           master
     35+            grand master

### Skill listing and learning times

    Skill 
    schools:

    600    Shipcraft                          three weeks
    610    Combat                             three weeks
    630    Stealth                            four weeks
    650    Beastmastery                       four weeks, 1 NP req'd
    670    Persuasion                         four weeks
    680    Construction                       three weeks
    690    Alchemy                            four weeks
    700    Forestry                           three weeks
    720    Mining                             three weeks
    730    Trade                              three weeks
    750    Religion                           five weeks, 1 NP req'd
    800    Magic                              four weeks, 1 NP req'd
    820    Weather magic                      five weeks, 1 NP req'd
    840    Scrying                            five weeks, 1 NP req'd
    860    Gatecraft                          five weeks, 1 NP req'd
    880    Artifact construction              six weeks, 1 NP req'd
    900    Necromancy                         six weeks, 2 NP req'd
    920    Advanced sorcery                   six weeks

    Shipcraft [600]
       601    Sailing                            two weeks
       602    Shipbuilding                       two weeks
       603    Fishing                            two weeks

    Combat [610]
       611    Survive fatal wound                four weeks
       612    Fight to the death                 four weeks
       613    Construct catapult                 two weeks
       614    Defense                            two weeks
       615    Archery                            three weeks
       616    Swordplay                          two weeks
       617    Weaponsmithing                     two weeks

    Stealth [630]
       631    Petty thievery                     two weeks
       632    Determine inventory of character   two weeks
       633    Determine skills of character      two weeks
       634    Determine character's lord         two weeks
       635    Conceal faction                    four weeks
       636    Learn of richest nearby noble      two weeks
       637    Torture prisoner                   three weeks
       638    Conceal self                       four weeks
       639    Sneak into structure               three weeks

    Beastmastery [650]
       651    Bird spy                           three weeks
       652    Capture beasts in battle           four weeks
       653    Use beasts in battle               four weeks
       654    Breed beasts                       four weeks
       655    Catch wild horses                  two weeks
       656    Train wild horse to riding horse   two weeks
       657    Train wild horse to warmount       three weeks
       658    Summon wild men                    three weeks
       659    Persuade wild men to remain        three weeks
       661    Breed hound                        three weeks

    Persuasion [670]
       671    Bribe noble                        two weeks
       672    Persuade oathbound noble           three weeks
       673    Raise peasant mob                  two weeks
       674    Rally peasant mob                  three weeks
       675    Incite mob violence                three weeks
       676    Train angry peasants               two weeks

    Construction [680]
       681    Construct siege tower              two weeks
       682    Stone quarrying                    two weeks

    Alchemy [690]
       691    Brew healing potion                three weeks
       692    Record skill on scroll             three weeks
       693    Extract venom from ratspider       three weeks
       694    Make potion of slavery             three weeks
       695    Collect rare elements              two weeks
       696    Brew potion of death               three weeks
       697    Turn lead into gold                four weeks

    Forestry [700]
       701    Construct battering ram            two weeks
       702    Harvest lumber                     two weeks
       703    Harvest yew                        two weeks
       704    Collect rare foliage               two weeks
       705    Harvest mallorn wood               two weeks
       706    Harvest opium                      two weeks
       707    Improve opium production           two weeks

    Mining [720]
       721    Mine iron                          two weeks
       722    Mine gold                          two weeks
       723    Mine mithril                       two weeks

    Trade [730]
       731    Conceal identity of trader         two weeks
       732    Find tradegood for sale            two weeks
       733    Find market for tradegood          two weeks

    Religion [750]
       751    Receive vision                     four weeks
       752    Lay to rest                        four weeks
       753    Preparatory ritual                 three weeks
       754    Resurrect dead noble               four weeks
       755    Remove blessing from soldiers      four weeks
       756    Immunity from Vision               four weeks

    Magic [800]
       801    Meditate                           two weeks
       802    Perform common tasks for gold      two weeks
       803    Appear common                      two weeks
       804    View current aura level of others  three weeks
       805    Heal                               three weeks
       806    Modern magic script                three weeks
       807    Reveal abilities of another mage   two weeks
       808    Tap health for aura                three weeks
       809    Shroud abilities from scry         three weeks
       811    Detect ability scry                three weeks
       812    Dispel ability shroud              three weeks
       813    Advanced meditation                three weeks
       814    Hinder meditation                  three weeks

    Weather magic [820]
       821    Fierce wind                        three weeks
       822    Bind storm to ship                 three weeks
       823    Scribe weather symbols             three weeks
       824    Summon wind                        three weeks
       825    Summon rain                        three weeks
       826    Summon fog                         three weeks
       827    Direct storm                       three weeks
       828    Dissipate storm                    three weeks
       829    Renew storm strength               three weeks
       831    Lightning bolt                     four weeks
       832    Seize control of storm             three weeks
       833    Fog of death                       three weeks

    Scrying [840]
       841    Scry location                      three weeks
       842    Ciphered writing of Areth-Pirn     three weeks
       843    Shroud location from magical scry  two weeks
       844    Dispel location shroud             three weeks
       845    Create magical barrier             three weeks
       846    Remove magical barrier             three weeks
       847    Locate character                   three weeks
       848    Detect location scry               three weeks
       849    Farcasting                         three weeks
       851    Save farcast state                 three weeks
       852    Banish undead                      three weeks

    Gatecraft [860]
       861    Detect gates                       two weeks
       862    Jump through gate                  two weeks
       863    Language of the Ancients           three weeks
       864    Seal gate                          two weeks
       865    Unseal gate with key               two weeks
       866    Notify if gate unsealed            two weeks
       867    Forcefully unseal gate             three weeks
       868    Reveal gate key                    three weeks
       869    Notify of gate jumps               two weeks
       871    Teleport                           two weeks
       872    Reverse jump through gate          three weeks

    Artifact construction [880]
       881    Forge auraculum                    two weeks
       882    Arcane symbols                     three weeks
       883    Forge magical weapon               three weeks
       884    Forge magical armor                three weeks
       885    Forge magical bow                  three weeks
       886    Curse noncreator loyalty           two weeks
       887    Reveal creator of artifact         two weeks
       888    Reveal where artifact was created  two weeks
       889    Destroy artifact                   three weeks
       891    Cloak creator of artifact          four weeks
       892    Cloak region of artifact creation  two weeks
       893    Dispel cloaking from artifact      three weeks
       894    Forge palantir                     three weeks

    Necromancy [900]
       901    Raise undead                       three weeks
       902    Summon ghost warriors              three weeks
       903    Runes of Evil                      three weeks
       904    Summon demon lord                  three weeks
       905    Renew demon bond                   three weeks
       906    Banish demon lord                  three weeks
       907    Eating of the dead                 four weeks
       908    Aura blast                         three weeks
       909    Absorb aura blast                  three weeks
       911    Transcend death                    four weeks, 1 NP req'd

    Advanced sorcery [920]
       921    Trance                             four weeks
       922    Teleport items                     four weeks
