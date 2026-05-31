---
title: Combat
weight: 8
toc: true
---

### Resolution of battle

Each kind of fighter has an attack and defense rating:

|              | attack | attack (shipboard) | defense | defense (shipboard) | missile |
| ------------ | ------ | ------------------ | ------- | ------------------- | ------- |
| peasant      | 1      |                    | 1       |                     |         |
| worker       | 1      |                    | 1       |                     |         |
| sailor       | 1      |                    | 1       |                     |         |
| soldier      | 5      |                    | 5       |                     |         |
| pikeman      | 5      |                    | 30      |                     |         |
| swordsman    | 15     |                    | 15      |                     |         |
| pirate       | 5      | (15)               | 5       | (15)                |         |
| knight       | 45     | (20)               | 45      | (20)                |         |
| elite guard  | 90     | (65)               | 90      | (65)                |         |
| crossbowman  | 1      |                    | 1       |                     | 25      |
| archer       | 5      |                    | 5       |                     | 50      |
| elite archer | 10     |                    | 10      |                     | 75      |

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

The chance that an attacker with attack rating `Ar` will score a hit against a defender with defense rating `Dr` is:

```
Ar / (Ar + Br)
```

For example:

| Attack Rating | Defense Rating | Chance to Hit                  |
| ------------- | -------------- | ------------------------------ |
| Ar = 90       | Dr = 45        | 2/3 chance of hitting defender |
| Ar = 90       | Dr = 90        | 1/2 chance of hitting defender |

If the attacker scores a hit against a noble, the noble will receive a random wound of 1-100 health points. Note that there is a 1% chance that a perfectly healthy noble will be instantly killed, and a greater chance that a previously wounded noble will die.

Wounded nobles do not continue fighting, even if their wounds are minor.

If a hit is scored against a fighter (soldiers, pikemen, archers, etc.), the fighter is killed. However, a blessed soldier has a 50% chance of surviving a hit.

A man successfully attacking a building or a ship will cause one point of damage to the structure. A siege engine attacking a building will cause 5-10 points of damage.

The winner in combat will attempt to take prisoners. The chance that a given defeated unit will be taken prisoner is proportional to the sizes of the remaining forces.

(Taking prisoners in battle and claiming loot requires many soldiers to run after the fleeing enemy. Thus, the chance of success is based on numerical advantage rather than combat skill.)

| Advantage | Chance of success |
| --------- | ----------------- |
| 1:1       | 25%               |
| 2:1       | 50%               |
| 3:1       | 75%               |

If the winning side outnumbers the defeated force by 2:1, there is a 50% chance that a given defeated unit will be captured. Defeated units which are not captured retreat from battle. If they are occupying a building or are located in a city, they may flee into the outlying province.

The victor always has at least a 25% chance of taking a unit prisoner, but no better than a 75% chance.

The victor will claim the defender's position in the location list if it is better, or will move into the defender's structure, ejecting the losing force. (The attacker may specify a flag to the `ATTACK` order to inhibit this behavior.)

Prisoners are stripped of their belongings by the victor, including any men accompanying the prisoners, such as workers or peasants. The stack leader of the winning force receives all of the loot from the battle.

When a noble is taken prisoner (including via `SURRENDER`), a portion of the prisoner's items will always be lost. If a unique item is lost, it will have to have found by exploration of the province.

### Front and rear

Units may issue the `BEHIND` command to declare whether they will line the front or the rear in battle. Rear units do not become targets for the enemy until all of the units in the rows in front of them have been killed. Only missile fighters, such as archers, may attack from the rear.

The leader of each stack (the top-most unit in the stack) will be the last unit to receive hits, regardless of its `BEHIND` status.

### Missile Attacks

Units in the rear may attack with their missile rating. Units in front attack with either their missile rating, or their attack rating, whichever is higher.

Thus, a noble with rating (attack=80, defense=80, missile=40) will do an attack of 40 when in the rear, and an attack of 80 when in front.

Weather effects on battle are as follows:

- Rain or wind cut the missile rating of archers or elite archers in half, and the missile rating of crossbowmen to 1/4 normal.
- Fog cuts the missile rating of all figures to 1/4 normal.

If a fighter's missile rating is zero, it can not attack from the rear. In this case, the fighter will use its attack rating like a regular soldier.

### Fortifications

Structures which may aid fighters in battle are rated for their defensive bonus.

```report
Castle Imperius [gx56], castle, defense 25
```

During battle, the fortification rating is added to the defense number for the men who fit inside the structure.

| structure           | men protected |
| ------------------- | ------------- |
| Castle              | first 500     |
| Tower               | first 100     |
| Galley or roundship | first 50      |
| Other structures    | first 50      |

Attacking fighters may randomly select the structure instead of an enemy fighter. The attack is resolved in the same way as for two fighters, using the attacker's attack rating and the structure's defense rating. If the attacker is successful the structure's defense rating will be lowered by one point.

Once the defense rating reaches zero, further hits will cause the building to become damaged. A fully damaged building (100% damage) will collapse, ejecting its occupants.

Siege engines always select the structure as a target.

| engine        | attack | defense | missile |
| ------------- | ------ | ------- | ------- |
| catapult      | 25     | 200     | 25      |
| battering ram | 30     | 250     |         |
| siege tower   | 30     | 250     |         |

Siege engines do 5-10 points of damage to the structure per hit.

Siege engines are not used in combat at sea.

### Item bonuses

Nobles may possess items which grant attack, defense or missile bonuses in combat. Only one item may be wielded for each category. If the noble possesses multiple items with bonuses in the same area, the item with the largest bonus will be chosen.

For example, suppose Osswid had the following items:

| item                      | bonus       |
| ------------------------- | ----------- |
| Shield of Achilles [fx78] | +25 defense |
| Sword of Death [gl23]     | +10 attack  |
| Mithril Axe [wt29]        | +15 attack  |
| Magic javelin [ht02]      | +5 missile  |

Osswid would wield the javelin and axe, and wear the shield. He would not use the Sword of Death.

Nobles automatically use items with combat bonuses in battle. No special orders are needed to wield them.

## Prisoners

Characters may become prisoners by losing to an enemy in battle, or by using the `SURRENDER` order (see orders section for description).

Since prisoners are unable to report where they are and what they are seeing, they do not contribute to the turn report of their faction. The player's turn report will show that that a unit is being held prisoner, but little else.

Prisoners will not execute any orders while they are in captivity. Queued orders will remain pending, but none will be processed.

Prisoners when spotted appear as stacked units, marked with the `prisoner` string:

```report
Seen here:
  Kosar the Indefectible [2022], with six peasants, one archer,
  two soldiers, accompanied by:
    Alion Krysaka [2785], prisoner
```

Unstacking a prisoner sets them free. Kosar could free Alion by ordering `unstack 2785`.

Prisoners may be transferred between units with the `GIVE` command. Kosar could transfer Alion to Osswid [501] by ordering:

```orders
give 501 2785
```

### Prisoner escapes

Prisoners are always on the lookout for ways to escape. Units holding prisoners can reduce their chances by remaining inside a structure, not transferring prisoners with `GIVE`, and not traveling with prisoners.

Each week (four times each game turn), a prisoner being held by a unit which is outside of a building has a 2% chance of escaping. A prisoner being held by a unit which is inside a structure, such as a castle, tower, inn or ship, has a 1% chance of escaping.Each time a prisoner is transferred with the give order, there is a 2% chance of a escape. Also, each time a unit holding a prisoner engages in travel which takes longer than one day, there is a 2% chance that the prisoner will be able to get free. Thus, short movement, such as entering or exiting a building, will not give the prisoners additional opportunities for escape, but traveling between provinces with prisoners will. Prisoners inside building or sub-locations will flee out into the surrounding location upon gaining their freedom. Escaped prisoners on ships will leap over the side and swim to a nearby shore.

## Permissions and declared attitudes

Commands dealing with permissions:

- `ADMIT`
- `HOSTILE`
- `NEUTRAL`
- `DEFEND`
- `DEFAULT`

Attitudes can be declared by or for either specific units, or an entire faction. For instance, player [613] could declare a permission or attitude for player [555], or a specific attitude for individual units within player 555's faction.

Declaring a permission for a player works so long as the player's units are not concealing their faction identity with Conceal faction [635], a subskill of Stealth [630].

### Allowing entrance and stacking

By default, a unit may not stack with a character belonging to another faction. A unit is also denied entry to a building or ship controlled by another player.

Players may allow units from other factions to stack with them or enter buildings or ships they control with the `ADMIT` order.

### Combat attitudes

A unit may have one of four combat attitudes to another unit:

- `HOSTILE` - Attack on sight.

- `DEFEND` - Defend other unit if attacked.

- `NEUTRAL` - Do nothing if other unit is attacked.

- `DEFAULT` - Neutral to units in other factions; Defend units in the same faction unless either one is concealing its lord.

Every character, and player faction entity, keeps three lists of units or other factions towards which it has declared attitudes. A unit is either on the `HOSTILE`, `DEFEND`, or `NEUTRAL` list. If a unit does not appear on any of the three lists, it has attitude `DEFAULT`.

Example:

```
player 778                     816
       hostile 816
units  4205                    6499
       4600                    6530, concealing lord
                               6599
```

Player 778 has declared player 816 hostile. One of 816's characters is concealing its lord.

If 4205 or 4600 run into unit 6499, they will attack it on sight. However, since 6530 is hiding its affiliation with 816, it will not be attacked on sight.

If 6499 is attacked and both 6530 and 6599 are present, 6599 will aid in the defense, but 6530 will not, because that might give away its affiliation.

If player 816 wanted 6530 to defend the faction's units anyway, either 816 or 6530 should issue the order `defend 816`. This would override the default attitude of units in the faction to one another.

Attitude toward units is considered before attitude toward the unit's faction. Thus, one may declare a faction hostile, but exclude certain units within the faction by specifically declaring them neutral.

A unit must be the top-most character in its stack to aid in defense. If a top-most unit joins a combat because of `DEFEND`, it will bring its entire stack along, even if the other members of the stack have not declared a `DEFEND` attitude.

Defenders only help when units are attacked, not when they initiate attacks. For example, if `A` has declared `defend B`, and `B` attacks `C`, `A` will not help `B`, even if `B` loses the battle.

Characters declared `DEFEND` to units which are guarding a province against pillaging will aid the guards if they are attacked, either explicitly with `ATTACK`, or implicitly via `pillage 1`.

Units which joined a combat because of a `DEFEND` declaration are shown with the qualification `ally` in the combat report.

