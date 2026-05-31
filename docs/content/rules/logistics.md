---
title: Loyalty, Stacking & Upkeep
weight: 6
toc: true
---

## Loyalty Bonds

Nobles are bound to their lords by one of three kinds of loyalty:

- _contract_ - The noble is being paid for his service

- _oath_ - The noble has taken an oath to serve his master

- _fear_ - The noble serves out of fear

The loyalty bond is rated. For example, a character's loyalty could be oath-1, contract-500, or fear-50.

A player's initial character has oath-2 loyalty.

Newly hired nobles have loyalty contract-500. Nobles are paid for their service with the **HONOR** command. A noble who issues `honor 50` will spend 50 gold to raise his own contract loyalty rating by 50 points.

A noble may take an oath of loyalty, pledging one or two Noble Points to secure it. This would yield loyalty oath-1 or oath-2. The **OATH** order secures an oath of loyalty for a noble.

Nobles may be terrorized by their masters. The severity of their treatment accumulates in the loyalty rating: fear-10, for instance. Fear loyalty is maintained with the **TERRORIZE** order.

Only one kind of loyalty may be active at a time.

Contract and fear loyalty decay over time. Contract loyalty loses the greater of 50 points or 10% of the current rating each month. Fear loyalty loses 1-2 rating points each month. Oath loyalty does not decay.

Units which fall to contract-0 or fear-0 have a 50% chance of deserting each month.

Nobles serving through contract or fear are susceptible to bribes, which may induce them to renounce loyalty to their lord, and pledge their service to the bribing faction. For details on bribing characters, see the `BRIBE` order.

Oath-1 nobles ignore all bribes. There is a persuasion skill which may cause an oath-1 noble to defect, although its use is difficult and rarely succeeds. Oath-1 nobles may reveal their factional affiliation if tortured.

Oath-2 nobles will not renounce loyalty to their lord under any circumstances, nor can they be forced to reveal any information about themselves.

Summary:

- _contract_ - Amount of gold invested in noble.
  Decays by max(50, 10% of current rating) each month.

- _fear_ - Severity of `terrorize` used on noble.
  Decays 1-2 points each month.

- _oath_ - 1 or 2 NPs may be invested in an oath bond.
  Does not decay.

Commands dealing with loyalty bonds:

- `BRIBE`
- `HONOR`
- `OATH`
- `TERRORIZE`

## Stacking

One unit may `STACK` under another unit. Two or more units grouped in this way are referred to as a _stack_.

Stacks move together and fight together. Here is a stack of four units:

```report
Law Netexus [2020], accompanied by:
  Feasel the Wicked [1109]
  Drakkar the Trader [1752]
  Alion Krysaka [2785]
```

Law Netexus is the stack leader, the top-most unit in the stack.

Only one level of stack depth is shown, so all that can be determined from the location report is that Feasel, Drakkar and Alion are stacked somewhere beneath Law Netexus. The exact arrangement of stacking bonds is not shown.

Feasel might be stacked under Law Netexus, with Drakkar and Alion under Feasel. Or Feasel, Drakkar, and Alion may all be stacked directly beneath Law Netexus.

Generally, such internal arrangements are only important when the stack breaks up. If Drakkar is stacked beneath Feasel, he will stick with Feasel if Feasel drops out of the stack. But if Drakkar were stacked beneath Law Netexus, he would not follow Feasel if Feasel unstacked.

If Law Netexus issues a `MOVE` order, the entire stack will move. If, however, Feasel issues a `MOVE` order, he will first drop out of the stack before moving.

Similarly, if Law Netexus engages combat with the `ATTACK` order, all characters in the stack will fight together. If one member of the stack is attacked, the entire stack will respond in defense.

Multiple levels of internal stacking can be useful if one wants several stacks to join together for a while, but then split apart later into their old arrangements.

Ocean ships may not be stacked together. There is no way to cluster ships into a fleet.

## Carrying capacity

Carrying capacity: Men and items are rated for how much they weigh, and how much they can carry, walking, riding or flying.

| Name         | Weight | Walking | Riding | Flying |
| ------------ | ------ | ------- | ------ | ------ |
| Man          | 100    | 100     | -      | -      |
| Riding Horse | 1000   | 150     | 150    | -      |
| Wild Horse   | 300    | self    | self   | -      |
| Warmount     | 300    | 150     | 150    | -      |
| Knight       | 400    | 100     | 100    | -      |
| Elite Guard  | 400    | 100     | 100    | -      |
| Ox           | 1000   | 1500    | self   | -      |
| Winged Horse | 300    | 150     | 150    | 150    |

_Man_ includes all of the varieties of men, including peasants, sailors, workers, etc. as well as nobles. A _knight_ includes both the man and the horse, hence the 400 weight.

In order to ride, the total riding capacity must cover the weights of all the units that may not ride themselves.

Examples:

- _man + riding horse_ - ride capacity is 150 - 100 = 50; walk capacity is 250

- _man + riding horse + wild horse_ - ride capacity is 150 - 100 = 50; walk capacity is 250

- _wild horse_ - ride capacity is 150 - 100 = 50; walk capacity is 1750; can walk or ride on its own, but will not carry anything

- _man + riding horse + ox_ - the ox may be driven alongside the horse, but will not carry anything when moving so quickly

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

```text
     +--- sailor ------- pirate
    +---- worker
   +----- crossbowman
  /
peasant ------ soldier ------ swordsman --- knight ----- elite guard
                 \
                  +---- archer ------ elite archer
                   +--- pikeman
                    +-- blessed soldier
```

Training a man may require that the noble have a certain skill, or possess some item. For instance, training soldiers into swordsmen requires a longsword [74] for each swordsman produced. The training character must also know Combat [610].

Some men may only be trained in certain kinds of locations. Elite guard and elite archers, for instance, may only be trained in castles. Blessed soldiers may only be trained in temples.

Training takes one day per man. Training five archers into five elite archers would take five days, for example. Training ten peasants into ten crossbowmen would take 10 days.

| num | kind            | skill | input man      | input item       | where  |
| --- | --------------- | ----- | -------------- | ---------------- | ------ |
| 11  | worker          | none  | peasant [10]   |                  |        |
| 12  | soldier         | 610   | peasant [10]   |                  |        |
| 13  | archer          | 615   | soldier [12]   | longbow [72]     |        |
| 14  | knight          | 616   | swordsman [20] | warmount [53]    |        |
| 15  | elite guard     | 616   | knight [14]    | plate armor [73] | castle |
| 16  | pikeman         | 610   | soldier [12]   | pike [75]        |        |
| 17  | blessed soldier | 750   | soldier [12]   |                  | temple |
| 19  | sailor          | 601   | peasant [10]   |                  |        |
| 20  | swordsman       | 616   | soldier [12]   | longsword [74]   |        |
| 21  | crossbowman     | 610   | peasant [10]   | crossbow [85]    |        |
| 22  | elite archer    | 615   | archer [13]    |                  | castle |
| 24  | pirate          | 616   | sailor [19]    | longsword [74]   | ship.  |

- A character needs no skills to train a worker.
- To train a sailor requires Sailing [601], a subskill of Shipcraft [600].
- Training archers and elite archers requires Archery [615], a subskill of Combat [610].
- Training swordsmen, knights and elite guard requires Swordplay [616], a subskill of Combat [610].

For more information and examples, see the `TRAIN` order.

## Maintenance cost

Men such as soldiers, workers, archers, etc. must be paid in gold monthly or they will leave the service of their noble. Peasants do not willingly leave a noble's service, but will starve if they are not paid. This cost is charged to the noble holding them at the end of each month.

If the noble does not have enough gold to pay his men, he will ask other nobles in his stack (provided they belong to the same player) for gold. Thus, only one member of a stack need carry gold for maintenance costs for the entire stack. Nobles will not share gold with units from other players.

If the noble can only afford to pay some of his men, one-third of those not paid will leave service at the end of the month. The computer chooses which men remain and which leave or starve.

| num | kind            | cost |
| --- | --------------- | ---- |
| 10  | peasant         | 1    |
| 11  | worker          | 2    |
| 19  | sailor          | 2    |
| 21  | crossbowman     | 2    |
| 12  | soldier         | 2    |
| 13  | archer          | 3    |
| 16  | pikeman         | 3    |
| 17  | blessed soldier | 3    |
| 20  | swordsman       | 3    |
| 24  | pirate          | 3    |
| 14  | knight          | 4    |
| 22  | elite archer    | 4    |
| 15  | elite guard     | 5    |

Note that nobles may `DROP` men to release them from service deliberately.

## Making weapons and armor

Weapons and armor are required for the training of some kinds of fighters. Archers require longbows, for instance, elite guard require plate armor, etc.

Weapons and armor are made with the `MAKE` command. The Weaponsmithing [617] subskill of Combat [610] is required to make weapons and armor.

| num | item        | material  |
| --- | ----------- | --------- |
| 72  | longbow     | yew [68]  |
| 73  | plate armor | iron [79] |
| 74  | longsword   | iron [79] |
| 75  | pike        | wood [77] |
| 85  | crossbow    | wood [77] |

One unit of the input material may turned into one of the desired items each day. For example, `MAKE` 72 2 would spend two days turning two yew [68] into two longbows [72].

