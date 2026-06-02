---
title: Orders
weight: 2
toc: true
---

### Command summary

| Command    | Arguments                               | Time     | Priority |
| ---------- | --------------------------------------- | -------- | -------- |
| ACCEPT     | <from-who> <item> [qty]                 | 0 days   | 0        |
| ADMIT      | <who or what> [ALL] [units]             | 0 days   | 0        |
| ATTACK     | <target> [flag]                         | 1 day    | 3        |
| BANNER     | [unit] "message"                        | 0 days   | 1        |
| BEHIND     | <number>                                | 0 days   | 1        |
| BOARD      | <ship> [maximum fee]                    | 0 days   | 2        |
| BRIBE      | <who> <amount> [flag]                   | 7 days   | 3        |
| BUILD      | <structure> "Name" [max days] [id]      | varies   | 3        |
| BUY        | <item> <qty> <price> [have-left]        | 0 days   | 1        |
| CATCH      | [number of horses] [days]               | as given | 3        |
| CLAIM      | <item> [number]                         | 0 days   | 1        |
| COLLECT    | <item> [number] [days]                  | as given | 3        |
| CONTACT    | <who>                                   | 0 days   | 0        |
| DECREE     | <decree>                                | 0 days   | 0        |
| DEFAULT    | <who>                                   | 0 days   | 0        |
| DEFEND     | <who>                                   | 0 days   | 0        |
| DIE        |                                         | 0 days   | 1        |
| DROP       | <item> <qty> [have-left]                | 0 days   | 1        |
| EMAIL      | <new email address>                     |          |          |
| EXECUTE    | [prisoner]                              | 0 days   | 1        |
| EXHUME     | [body]                                  | 7 days   | 3        |
| EXPLORE    |                                         | 7 days   | 3        |
| FEE        | [gold per 100 wt]                       | 0 days   | 1        |
| FERRY      |                                         | 0 days   | 1        |
| FISH       | [number of fish] [days]                 | as given | 3        |
| FLAG       | string                                  | 0 days   | 1        |
| FLY        | <direction or destination> [...]        | varies   | 2        |
| FORGET     | <skill>                                 | 0 days   | 1        |
| FORM       | <unit> "Name of new character"          | 7 days   | 3        |
| FORMAT     | <number>                                |          |          |
| GARRISON   | <castle>                                | 1 day    | 3        |
| GET        | <who> <item> [qty] [have-left]          | 0 days   | 1        |
| GIVE       | <to-who> <item> [qty] [have-left]       | 0 days   | 1        |
| GUARD      | <flag>                                  | 0 days   | 1        |
| HONOR      | <amount>                                | 1 day    | 3        |
| HOSTILE    | <who>                                   | 0 days   | 0        |
| IMPROVE    | [days]                                  | varies   | 3        |
| LORE       | <lore sheet>                            |          |          |
| MAKE       | <item> [qty]                            | as given | 3        |
| MESSAGE    | <# of lines of text> <to-who>           | 1 day    | 3        |
| MOVE       | <direction or destination> [...]        | varies   | 2        |
| NAME       | [unit] "new name for unit"              | 0 days   | 1        |
| NEUTRAL    | <who>                                   | 0 days   | 0        |
| NOTAB      | <number>                                |          |          |
| OATH       | <level>                                 | 1 day    | 3        |
| PASSWORD   | ["password"]                            |          |          |
| PAY        | <to-who> [amount] [have-left]           | 0 days   | 1        |
| PILLAGE    | [flag]                                  | 7 days   | 3        |
| PLAYERS    |                                         |          |          |
| PLEDGE     | <who>                                   | 0 days   | 1        |
| POST       | <# of lines of following text>          | 1 day    | 3        |
| PRESS      | <# of lines of text>                    | 0 days   | 1        |
| PROMOTE    | <who>                                   | 0 days   | 1        |
| PUBLIC     |                                         |          |          |
| QUARRY     | [number of stones] [days]               | as given | 3        |
| QUEST      |                                         | 7 days   | 3        |
| QUIT       |                                         |          |          |
| RAZE       | [building]                              | varies   | 3        |
| RECRUIT    | [days]                                  | as given | 3        |
| REPAIR     | [days]                                  | as given | 3        |
| RESEARCH   | <skill>                                 | 7 days   | 3        |
| RESEND     | [turn]                                  |          |          |
| RUMOR      | <# lines of text>                       | 0        | 1        |
| SAIL       | <direction or destination> [...]        | varies   | 4        |
| SEEK       | [who]                                   | 7 days   | 3        |
| SELL       | <item> <qty> <price> [have-left]        | 0 days   | 1        |
| STACK      | <character>                             | 0 days   | 1        |
| STOP       |                                         |          |          |
| STUDY      | <skill>                                 | 7 days   | 3        |
| SURRENDER  | <character>                             | 0 days   | 1        |
| TAKE       | <who> <item> <qty> [have-left]          | 0 days   | 1        |
| TERRORIZE  | <who> <severity>                        | 7 days   | 3        |
| TRAIN      | <kind> <days>                           | as given | 3        |
| UNGARRISON | <garrison>                              | 1 day    | 3        |
| UNLOAD     |                                         | 0 days   | 3        |
| UNSTACK    | <who>                                   | 0 days   | 1        |
| USE        | <skill> [arguments...]                  | varies   | 3        |
| VIS_EMAIL  | <new email address for the player list> |          |          |
| WAIT       | conditions                              | varies   | 1        |

## Submitting orders

Orders for Olympia game 4 should be sent to _olympia@shadowlandgames.com_.

The **Reply-To:** header on turn reports is set to this e-mail address, so using the reply feature of your e-mail client should send orders to the right place.

Your orders are automatically loaded into the game and queued for your units. The scanner will send a reply as soon as it processes your mail, showing whether or not there were any errors with the orders it received.

Note that the order scanner does not do an exhaustive check of your orders' syntax; it checks that the commands given exist, and checks the parameters of those commands that are executed at parse time - notably those involved in parsing (begin, unit, password, email and vis_email), in report formatting (format and notab) and those which have immediate secondary effects (resend and lore). Other commands' parameters are only checked when they are executed during the turn run.

Orders must be of the following form:

```orders
begin player-number password

email, lore, password commands

unit player-number

  commands for player entity: name, format or quit

unit unit-number

  commands for unit

unit unit-number

  commands for unit

end
```

The **Subject:** line on your message is ignored.

The **BEGIN** keyword tells the order scanner what your player number is. If you have not set a password, you do not need to supply one.

The **UNIT** command replaces a set of orders for a unit. Any pending orders for the unit will be cleared, and the new orders sent in will queue up. Orders that are still executing for the unit will not be interrupted unless the first order queued is the stop order.

Do not match an **END** for every unit command! There should only be one `end`, at the end of all of the unit sections. The Olympia order parser will not read beyond the `end`.

For example, here is a set of orders for player Fate [812], who has two characters, Osswid [5499] and Candide [1269]:

```orders
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
```

The parser tries to be as flexible as possible. It is case insensitive and is not strict about spaces on a line, so you may use indentation to make your orders more readable.

Orders may be commented with the `#` character. Everything from a `#` to the end of the line will be ignored by the parser:

```orders
move north # Head to Drassa to meet up with Osswid
stack 5499 # stack with him
```

No one will read the comments but you. Neither the GM nor the Olympia engine will try to interpret them for any reason.

Note that arguments must be enclosed in quotes if they are more than one word:

```orders
name "Osswid the Constructor"
```

The acknowledgement will show any errors that occurred while the orders were being parsed, and list the current pending commands for all of your units.

There is a limit of 250 orders which may be queued per unit. Additional orders will be dropped and will not appear in the unit's command queue.

## Interrupting orders

Suppose the turn report shows the following orders queued:

```orders
unit 5499
  # > study 160 (executing for three more days)
  recruit 10
  explore
```

Sending in new orders for this unit will not disturb the still-running `STUDY` command unless the first order is **STOP**.

For example: If this were sent in:

```orders
unit 5499
  move south
```

This would be the result:

```orders
unit 5499
  # > study 160 (executing for three more days)
  move south
```

To interrupt the `STUDY` and get on with the `MOVE` right away, instead send in:

```orders
unit 5499
  stop
  move south
```

This will show:

```orders
unit 5499
  # > study 160 (executing for three more days)
  stop
  move south
```

Note that the `STOP` queues like any other order; it does not actually interrupt the executing command until the turn runs. This means that the `STOP` itself can be replaced by sending in another set of orders later.

## Units not controlled by you (yet)

Orders may be sent in for units which are not yet under control, such as characters that you intend to bribe or terrorize into switching to your faction.

As soon as the unit comes under your control, the orders queued for it will begin to execute.

Orders may also be sent in for new nobles which will be formed during the turn. First choose one of the possible unit numbers from the choices listed near the beginning of the turn report:

```report
The next five nobles formed will be: 5717 3215 4902 4489 5628
```

Supply one of these numbers as the first parameter to the form order:

```orders
form 5628 "Feasel the Wicked"
```

Then queue some orders for Feasel to execute as soon as he appears:

```orders
unit 5628
  unstack
  study 160
  move out
  recruit
```

## Use the order template

An order template appears at the bottom of the turn report. This template lists all of the units for a player and shows any pending orders for those units.

Order template

```orders
begin 812 # Master Bogomil's Family

unit 2508 # Tudor
  # > make 74 (still executing)

unit 2947 # Milo

unit 4375 # Beorn
  # > move s (executing for one more day)
  pillage
  recruit

unit 4763 # Sylvia

unit 5977 # Drango
  # > collect 87 0 0 (still executing)
  sail e
  sail s
  fish
  explore

unit 5418 # Comte de le Sang

end
```

Note that the layout of the order template matches the syntax the order scanner expects. Many players find it convenient to edit this template to add or change commands for their units. Mail everything from the `BEGIN` to `END` (inclusive) to the order scanner. It is wise to save a copy of the orders you submit in case there are errors and they need to be resent.

---

**Be careful**

Beware of sending in different sets of orders too quickly. Sometimes messages sent within a short time of each other will arrive out of order. This can wreak havoc on your turn if the wrong orders arrive last. Compose your orders offline and review them before mailing. A simple typographical error in your orders could ruin your whole turn!

Some players make clever use of the **PASSWORD** order to make sure that an order set lingering out in the email system on the network which arrives late won't replace a more recent order set sent in. For example, say you submit some orders on Friday, and they don't show up by Monday. Monday you send an updated set of orders to order scanner and get an instant reply. The Friday orders have not arrived yet, but you're worried that they're out there and will arrive sooner or later, replacing the newer set of orders you just sent in. Solution: issue a `PASSWORD` order in the newer set, so the Friday orders will fail when they do arrive.

## Failed orders

Commands that fail generally take zero time.

For instance, if `STUDY` is issued for a skill which the location does not offer, it will fail immediately, and take zero time. The failed `STUDY` order will not take a week, and it will not count toward the limited study time for that month.

Production commands fail immediately if none of their input resources are available. For instance, `RECRUIT` in a location with no peasants will immediately fail, taking zero time. However, resources may sometimes become depleted while the command is being executed. In such cases, the command may fail even after it has spent some time executing.

## More order examples

You can replace orders for some units, but leave pending orders for other units alone, by only including unit sections for the ones you want to change.

If you want to see what orders are queued, but not change anything send in:

```orders
begin 999 # whatever your player number is ...
end
```

To change your email address, send in:

```orders
begin 999
email <new@address.com> # give your new address
end
```

As a security measure, the confirmation will be sent to both the new and old addresses.

To change the name of your faction, issue the **NAME** order for the faction's player entity:

```orders
begin 999
unit 999
  name "Seekers of Fame and Power" ...
end
```

Important: Don't forget the unit command for the player entity.

## Player Entity

Each Olympia player faction has a number. This number is represented by an entity in the game. However, unlike a character, this entity is mostly used as a place holder for the faction. No one can see the faction entity, and it can issue very few orders.

For example, suppose player Fate [501] has one character Osswid [5499]. Fate does not exist in any location, so it does not receive a location report, and no one can see it. However, Osswid, being the player character for the faction, is sworn to Fate [501]. Fate may execute only a limited set of administrative orders.

Characters, not factions, issue most orders. Do not try to form or recruit with the player entity. For most turns, the player entity will have no orders queued for it.

The orders a player entity may issue are:

- ACCEPT
- ADMIT
- DEFAULT
- DEFEND
- FORMAT
- HOSTILE
- MESSAGE
- NAME
- NEUTRAL
- NOTAB
- PRESS
- REALNAME
- RUMOR
- TIMES
- QUIT

## Quitting

To drop out of the game, issue the quit order for your player entity.

For instance, player 501 would quit by sending in the following orders:

```orders
begin 812 password
unit 812 # don't forget unit for the player unit!
  quit
end
```

No turn report will be sent for the turn in which a player quits.

## Order of characters in a location

When a character enters a location, it is added to the end of the list of characters already there. The unit that has been in a location the longest will appear at the top of the list.

If a character leaves a location and later returns, it will be put at the end of the list again.

For example:

```report
Seen here:
  Candide [1269]
  Osswid [5499]
  Feasel the Wicked [1109]
```

Candide has been here longest, followed by Osswid, then Feasel. If Candide were to leave and return, he would appear at the end of the list.

If a character unstacks from beneath another unit, the character will appear just after the unit, rather than at the end of the list.

For example:

```report
Seen here:
  Candide [1269], accompanied by:
    Osswid [5499]
    Feasel the Wicked [1109]
```

Osswid is stacked beneath Candide. If Osswid unstacks, he will appear after Candide, not after Feasel:

```report
Seen here:
  Candide [1269]
  Osswid [5499]
  Feasel the Wicked [1109]
```

Note: new player characters are added to the top of the list of characters in the safe haven in which they join, not the bottom. Additonal characters formed by them will appear at the bottom of the list as usual.

## Command priority

All orders have a priority of 0-4.

- Permission commands (`ADMIT`, `HOSTILE`, etc.) are priority 0.
- Zero-time commands and wait are priority 1.
- `MOVE` and `FLY` are priority 2.
- The `SAIL` command is priority 4.
- All other commands are priority 3.

The order scheduler will first try to start all priority 0 orders. Only when no more priority 0 orders are ready to start will a priority 1 order be started.

In other words, the order scheduler will not start an order at a higher priority when an order may be started at a lower priority.

Orders at the same priority are resolved in location order. If two units in a location are both waiting to start a `MOVE` order, the first unit in the location will go first.

The above description of order priorities may seem complicated, but the intent is to let players ignore same-day synchronization issues in most cases. Rather that needing `WAIT` to guarantee that give happens before move, the lower priority of `GIVE` makes this happen naturally.

For example, consider three units stacked together, top, mid and bot:

```
top:
  move ec69
  yew

mid:
  unstack
  recruit

bot:
  recruit
```

These should be executed in the following order:

1. _mid_: unstack # unstack is prio-1
2. _top_: move ec69 # move is prio-2
3. _mid_: recruit # recruit is prio-3 (_top_ and _bot_ arrive at ec69)
4. _top_: yew # yew is prio-3
5. _bot_: recruit # recruit is prio-3

The `UNSTACK` happened first since it's a priority 1 command. The `MOVE` went second. When _top_ and _bot_ finished moving, there were only priority three commands left, so they ran in location order.

### Command precedence within a location

```report
Seen here:
  Candide [1269]
  Osswid [5499]
  Feasel the Wicked [1109]
```

Order precedence within a location is an advantage for commands or skill uses which obtain resources from the location. For instance, if Candide and Osswid both attempted to `HARVEST` all of the lumber available in their location, Candide would have precedence, since his `HARVEST` order would finish before Osswid's, if they were started on the same day.

### At the same time

No two things ever happen at exactly the same instant in Olympia. Someone always goes first.

Suppose two characters were outside of a building (which nobody is inside), and both wanted to enter, to claim it:

```report
Seen here:
  Osswid [5499]
  Candide [1269]

Inner locations:
  Hooting Own Inn [3102], inn
```

Both Osswid and Candid issue `move 3102` as their first order on day 1 of the month. What happens?

Osswid's command begins before Candide's, since Osswid appears before Candid in the location list. Therefore, Osswid will enter the inn first.

```
Osswid             ?            Candide
------- 8 days -------- 8 days --------
city A          city B          city C
```

If Osswid and Candide both leave for city `B` on the same day, we cannot predict who will get their first.

