---
title: Geography & Movement
weight: 3
toc: true
---

## Olympian geography

Olympia's map is a large grid of locations called provinces. Groups of provinces form continents, islands and oceans. These collections are called regions, and are usually named.

A province's description will include a list of the directions in which a character may travel:

```report
Plain [ae48], plain, in region Tollus
  Routes leaving Plain [ae48]:
    North, to Plain [ad48], 7 days
    East, to Plain [ae49], 7 days
    South, to Ocean [af48], Tymaerian Sea, 1 day
    West, to Ocean [ae47], Tymaerian Sea, 1 day
```

This is a non-descript province in the Tollus region.

From this province, a character may travel north or east on foot or by horse, or may sail by ship to the south or west.

```
move north # OR move n OR move ad48
move east  # OR move e OR move ae49

sail south # OR sail s OR sail af48
sail west  # OR sail w OR sail ae47
```

Land movement will automatically use the fastest available mode. For example, if a character has enough horses for all of the members in the party to ride, then the travelers will go on horseback.

Ocean movement requires that the character be in a ship.

Route distances are rated for the number of days it normally takes to traverse them. Land distances are rated for a lightly loaded character walking, and ocean distances are given for an ordinary ship traveling in normal weather.

Actual travel times may differ from times given in the route listing. Land distances depend on the surrounding terrain and the modes of transport available. For example, horses often speed up movement, but over especially rough or treacherous terrain, they may actually slow travel because they must be led and managed. A stiff wind may speed ocean vessels, while lack of wind may slow their progress.

### Inner Locations

A province may contain sub-locations within its borders. Sub-locations may usually only be entered from the surrounding province. They will be listed separately in the location description:

```report
Inner locations:
  Carim [em28], city, 1 day
```

The city Carim may be entered with the `MOVE em28` order. Travel into a city requires one day.

**Note**: _`MOVE IN` may be used to enter a sub-location, although this order may be ambiguous if the location contains more than one sub-location. In such a case, the first sub-location in the Inner locations list will be entered. Using `MOVE IN` is not recommended if the entity number of the sub-location is known._

Characters in a sub-location will receive a report for the surrounding province. However, characters in the outer province will not normally be able to see into an inner location without entering it.

### Inside a City

```report
Carim [em28], city, in province Plain [ae48]
  Routes leaving Carim [em28]:
    Out, to Plain [ae48], 1 day

Inner locations:
  Hooting Own Inn [ep76], inn
```

Characters in the city Carim may move out (or `MOVE ae48`) to the surrounding province. They may also attempt to enter the inn, which is a sub-location of the city. Notice that no travel time rating is listed for the inn; entering it takes no time (zero days).

```orders
move out # OR move ae48
move ep76
```

A character in Carim will receive a location report both for the city as well as the surrounding province Plain [ae48]. Characters in the Hooting Own Inn will receive a location report for the inn and one for Carim, but will not get a report for Plain [ae48]. A character in the city may not see inside the inn without entering.

Characters in a sub-location receive a report for the immediate surrounding location.

Characters are not able to see into inner locations without going into them.

#### City default garrisons

Every non-safe-haven city in the regular world (except safe havens) has an initial default garrison with 25-150 pikemen. Each garrison is set to _ADMIT all_ and _DEFEND all_. Each noble stacked with the garrison will earn 2gp/day for aiding the city's defenses.

### Who else is here?

Characters spotted will be listed in the location report:

```report
Seen here:
  Fighters of Pelenth [2019], "carrying a gold banner"
  Osswid the Constructor [5499]
```

All characters listed as being seen in the location may interact without requiring any travel. Thus, the Fighters of Pelenth and Osswid are considered to be in essentially the same place.

This is true whether the characters are in a province, a city, a ship, an inn, or some other sub-location.

However, a character in a sub-location may not interact with characters in the surrounding area. A noble in the city must first enter the inn before he may interact with those inside.

### More about geography

Olympian provinces are arranged in a square grid. Travel is possible in the four main compass points. Thus, to move diagonally, two move orders are required. To move northwest, for instance, one would first need to `MOVE n`, then `MOVE w`.

A province's map coordinate can be read from its bracketed code. The leading letters give the row; the trailing number gives the column. The northwest corner of the grid is [a1], with rows increasing to the south and columns increasing to the east. So `Plain [ad48]` lies directly north of `Plain [ae48]`, and `[ae49]` directly east of it.

```
      1    2    3   ...
    +---------------------
 a  |  a1   a2   a3
 b  |  b1   b2   b3
 c  |  c1   c2   c3
 .  |
 z  |  z1   z2   z3
 aa |  aa1  aa2  aa3
 ab |  ab1  ab2  ab3
```

Rows are lettered with a sequence that skips the letters most easily confused with digits — `i`, `j`, `l`, and `o` are left out:

```
a b c d e f g h k m n p q r s t u v w x y z
```

Columns are plain numbers. Neither part is padded with leading zeros, so the corner is `a1`, not `a01` or `aa01`. After the single letter `z` — the 22nd row — rows continue with two letters, `aa`, `ab`, and so on, so the map has no fixed size limit.

Entity numbers for sub-locations do not correspond to any coordinate system.

The edges of the map are not passable, so for example it is not possible to travel either north or west from [a1].

### Holes in the Map

The map may have some holes, representing impassable provinces. Routes into some provinces may also be hidden.

```report
Plain [cd21], plain, in region Tollus Routes leaving Plain [cd21]:
  North, to Plain [cc21], 7 days
  East, to Plain [cd22], 7 days
  West, to Plain [cd20], 7 days
```

Notice the lack of a southern exit. This means that there is no known southern route from Plain [cd21], into what should be Plain [ce21]. Exploration may find a southern route, but it is possible that none may ever be found, and the terrain to the south is completely impassable.

### Hidden routes

If exploration finds a hidden route, any noble in the player's faction will be able to use it.

```report
> explore
  A hidden route has been found!
  South, to Plain [ce21], 7 days
```

The location description for this place will now include the hidden route:

```report
Plain [cd21], plain, in region Tollus
  Routes leaving Plain [cd21]:
    North, to Plain [cc21], 7 days
    South, to Plain [ce21], 7 days, hidden
    East, to Plain [cd22], 7 days
    West, to Plain [cd20], 7 days
```

However, units from other factions, even if they know that the hidden route's entity number is [ce21], will not be able to travel across it.

All factions with units in a stack traveling across a hidden route, with the exception of units being held prisoner, will learn of its existence. Nobles from factions wanting to learn how to use the hidden route can stack with a noble about to move across the route.

### Ocean ports

A ship in an ocean province may sail into an adjoining land province.

```report
Ocean [cw12], ocean, in South Sea
  Routes leaving Ocean [cw12]:
    North, to Ocean [cv12], Atnos Sea, 4 days
    East, to Mountain [cw13], West Camaris, impassable
    South, to Plain [cx12], West Camaris, 1 day
    West, to Ocean [cw11], 4 days

Inner locations:
  Island [eb97], island, 1 day
```

A ship sailing in this ocean province may dock by sailing to Plain [cx12] or Island [eb97].

Ships may not dock in mountain provinces, as the rocky cliffs are too dangerous to approach. Routes between ocean and mountain provinces are marked `impassable`.

### Port Cities

A city in a province adjoining an ocean will have been founded on the best spot for an ocean port. The ocean will only be accessible through the port city in this case, and not through the surrounding region.

```report
Plain [ae48], plain, in region Tollus
  Routes leaving Plain [ae48]:
    West, to Ocean [ae47], Tymaerian Sea, impassable

Inner locations:
  Carim [em28], port city, 1 day
```

Note that from the province surrounding the port city, access to the ocean is not possible.

```report
Carim [em28], port city, in province Plain [ae48]
  Routes leaving Carim [em28]:
    West, to Ocean [ae47], Tymaerian Sea, 1 day
    Out, to Plain [ae48], 1 day
```

However, ships may sail into and out of the port city itself. From the Tymaerian Sea, this looks like:

```report
Ocean [af48], ocean, in Tymaerian
  Sea Routes leaving Ocean [af48]:
    North, city, to Carim [em28], Tollus, 1 day
    South, to Ocean [ag48], 3 days
```

#### An example city description

```report
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
```

### Wilderness and civilization

Every province has a civilization level. Provinces with no civilization (a level of zero) are considered wilderness. Civilization levels for provinces are shown in the turn report:

```report
Mountain [cq24], mountain, in Lesser Atnos, civ-1

Forest [ac35], forest, in Torba Bacor, wilderness
```

The civilization level of a province is determined by the presence of cities and buildings, or half of the maximum civilization level of its surrounding provinces, whichever is higher.

There is no fixed civ level cap. However, only the first building of each type counts towards the civ level in a location.

| Feature    | Contribution                |
| ---------- | --------------------------- |
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

#### Hades

Hades, also known as The Land of the Dead, is a subterranean world populated with demons, ghouls and spirits thirsty for the blood of the living. Only the bravest warriors should consider walking these dark paths.

#### Faery

The Faery world lies in a nearby, but separate reality. Occasionally a Faery hill will exist simultaneously in both Faery and the outside world. During this time, mortals may cross between the two worlds. Faery is protected by the Faery Hunt, a tough band of elves armed with magical bows. Each hunt group consists of 10-50 elves, each with a combat rating of (50,50,100). Rumors speak of a magical talisman, the elfstone, which allows mortals to pass unharmed in Faery, and to summon Faery hills to the mortal world.

#### The Cloudlands

The Cloudlands is a small region which floats over Mt. Olympus and the Imperial City. It is generally only accessible by flight. The Cloudlands is home to three cities: Nimbus, Stratos and Aerovia. Weather magic is taught in these three cities.

### The Map

    TODO: insert ASCII image of map here

