---
title: Ships
weight: 11
toc: true
---

## Ships

Players may build two kinds of ships in Olympia: _galleys_ and _roundships_.

A galley, also known as a warship, is a slender rowed vessel. Galleys require 14 sailors as oarsmen for travel, and may carry up to 5,000 units of cargo.

The roundship, also known as the merchantman, is a deep, wide sailing ship, usually with one or two masts and steered with great, oar-like paddles. The large cargo space makes it well suited for trade and extended ocean travel.

With favorable winds, roundships will make better time than galleys, even when fully loaded. Roundships require a crew of eight sailors travel, and may carry up to 25,000 units of cargo.

The Shipcraft [600] skill category allows nobles to train sailors, and has sub-skills for building and sailing ships.

| ship      | capacity | sailors |
| --------- | -------- | ------- |
| galley    | 5,000    | 14      |
| roundship | 25,000   | 8       |

Ships may be damaged while sailing by storms, or by submerged rocks in coastal waters. Damaged ships may be repaired with the `REPAIR` command. Repairing a damaged ship requires one unit of pitch [261].

The capacity of a damaged ship is a reduced in proportion to the amount of damage. A galley with 10% damage may only carry 4,500 units of cargo.

### How to sail

The Sailing [601] subskill of Shipcraft [600] is required to pilot a ship. Shipcraft may be learned in any port city. For information about piloting galleys and roundships, see the sail order.

### Making ships

Construction of ships requires a character to have the Shipbuilding [602] skill, a subskill of Shipcraft [600].

| ship      | effort          | material |
| --------- | --------------- | -------- |
| galley    | 250 worker-days | 50 wood  |
| roundship | 500 worker-days | 100 wood |

To begin construction of a ship, the shipbuilder should unstack from beneath other characters and issue one of the following `BUILD` orders:

- `build galley "name of galley"`
- `build roundship "name of roundship"`

One-fifth of the lumber will immediately be deducted from the shipbuilder's inventory and put to use building the new ship. The shipbuilder and his workers will be placed inside the new ship.

At least three workers are needed to begin construction of a ship.

Until the ship is completed, it will be shown as `in progress`:

```report
Ships docked at port:
  HMS Pinafore [1111], galley-in-progress, 28% completed, owner:
    Osswid the Constructor [5499], with five workers
```

The remaining construction materials are deducted from the builder's inventory as work on the ship progresses. A second fifth of the lumber is required when the ship becomes 20% complete, a third fifth when the ship becomes 40% complete, etc. Construction halts if the builder runs out of materials

For example, a noble who wanted to build a galley would need to know Shipbuilding [602], have at least three workers, and start with at least 10 wood [77]. (40 more wood is required to bring the galley to completion).

As soon as the required number of worker-days has been invested in construction, the ship will be christened and declared seaworthy.

To resume construction of a partially completed ship, first enter the ship, then issue the either `BUILD GALLEY` or `BUILD ROUNDSHIP`.

Ships may only be built in port cities.

### Operating a ferry

There are several commands useful for operating a commercial ferry:

- `FEE` - Sets the fee passengers will be charged

- `BOARD` - Passengers use this to pay their fee and board a ferry

- `UNLOAD` - Unload passengers once the destination is reached

- `FERRY` - Signals passengers waiting in port that they may board

The captain of a ferry must issue `FEE GOLD` command to set a fee which will be charged to passengers wishing to `BOARD` his ship. The fee is expressed as how many gold pieces per 100 weight of the passenger's stack will be charged.

For instance, if the captain wanted to charge 1/2 gold per unit weight of the passengers stack, he would issue `fee 50`. (a 50 gold fee for every 100 weight).

The fee is a property of the captain, not the ship.

A character may clear the fee with the `fee 0` order. If no fee is set, then the ship is not considered to be operating as a ferry, and characters are may not use the `BOARD` order to enter the ship.

Passengers issue `board ship` to board a ferry. The order will fail if the ship is not present, or if it is not operating as a ferry (the captain of the ship has no fee set). `BOARD` will cause the character to pay the captain the required boarding fee, then move the character's stack onto the ship.

The captain shouldn't issue an `ADMIT` order to let characters on board who will be paying to take the ferry. Otherwise they could board the ship with `move` instead of `BOARD`, bypassing the ferry fee.

### Ferry synchronization

Suppose now that the captain has a ship, has set a fee, announced his service in the _Olympia Times_ and with `POST`, and is now ready to ferry passengers. What should he tell them? How will the synchronization work?

Passengers should travel to the port the ferry will be arriving at and issue `WAIT FERRY SHIP`.

When the ship arrives at port, the captain should order `UNLOAD` to eject his current load of passengers, then `FERRY` to signal any passengers waiting in port that they may now board.

Note that passengers should not use `WAIT SHIP` to wait for the ferry. Otherwise, they will attempt to `BOARD` as soon as the ship reaches the port. The captain's `UNLOAD` order may not have executed at this point. In this case, `BOARD` orders might fail because there won't be any room to enter the ship until the existing load of passengers has disembarked.

Example:

```report
Posted by Captain McCook [3402]:
  "Captain McCook's ferry to Drassa departs each
  Sunsear on the 15th.  The fee is 1 gold/wt.
  Issue WAIT FERRY 1234 then BOARD 1234 for
  ferry service.   Arrr!"
```

Captain McCook's orders:

```orders
> fee 100                 # 1 gold/wt. is our fee
> sail ...                # arrival at port
> unload                  # unload current passengers
> wait day 15             # wait until stated time of departure
> ferry                   # sound our horn
> sail ...                # on to the next port
```

Passengers wishing to travel on McCook's ferry:

```orders
> wait ferry 1234         # [1234] is McCook's ship
> board 1234              # pay our gold and embark
> wait loc destination    # Captain McCook will unload us at the
                          # end of the journey, so wait until we
                          # find ourselves in the destination city.
```

