---
title: Buildings & Economy
weight: 9
toc: true
---

## Buildings

The following structures may be built by characters with the Construction [680] skill:

| kind   | effort | material  | skill      | where                  |
| ------ | ------ | --------- | ---------- | ---------------------- |
| inn    | 300    | 75 wood   | 680        | province or city       |
| mine   | 500    | 25 wood   | 680 or 720 | mountain or rocky hill |
| temple | 1,000  | 50 stone  | 680        | anywhere               |
| tower  | 2,000  | 100 stone | 680        | anywhere               |
| castle | 10,000 | 500 stone | 680        | province or city       |

The term `anywhere` means a province, city, or other sublocation. Buildings may not be built inside other buildings, with the exception of towers. Up to six towers may be built inside a castle.

Effort is in worker-days.

The builder must know the required skill, have at least three workers, and possess at least one-fifth of the necessary construction materials in his inventory.

To start building, unstack from beneath other characters and issue one of the following `BUILD` orders:

- `build inn "name of inn"`
- `build mine "name of mine"`
- `build temple "name of temple"`
- `build tower "name of tower"`
- `build castle "name of castle"`

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

```report
Candide the Captain [1269], with ten sailors
Osswid [5499], accompanied by:
  Feasel the Wicked [1109]
```

Since Candide is first in the ship's character list, he is the owner of the ship. Only Candide may issue sail orders, or change the ship's name. If Candide were to leave, Osswid would become the new captain of the ship.

The owner of a building or ship may determine who may enter with the `ADMIT` order. The default is to refuse entry to units from other factions.

If the first character inside a building leaves, he is no longer the owner.

A building is a castle, tower, inn, temple or mine.

## Mining

A mine is a deep shaft or tunnel which allows workers to extract valuable resources from the earth, such as iron and gold. At most one mine may be built in each mountain province or rocky hill.

A new mine has an initial depth of one. The mine shaft becomes deeper as characters use it to obtain natural resources. The shaft will become one level deeper for every three uses of a mining extraction skill.

As the depth of the shaft increases, the mix of resources obtainable changes. Iron is usually found nearest the surface. As one proceeds deeper, gold may be found in higher quantities. Other rare elements may be found by going deeper still.

The deeper a mine becomes, the more frequently cave-ins or other accidents will occur. With each accident, the mine's damage percentage will rise. If not attended with `REPAIR`, the mine will eventually collapse. Once a mine collapses, it remains in the province for one game year (eight game months). Characters may not enter or use a collapsed mine. After the year has passed, the collapsed mine will vanish, and a new mine may be built in the location.

## Opium

Opium is produced in swamp regions, and consumed by markets in desert, plain, forest, and mountain provinces. All markets (except those in swamp regions) have some level of opium demand. However, this demand will not be visible in the market report at low levels.

Satisfying opium demand in a market will cause the next month's demand to be higher. As peasants become addicted to opium, the increased demand will be shown in the location's market report. If no opium is sold to a market, the demand will fall.

Opium adversely affects the city's tax base. The more opium the market buys, the more tax revenues will be reduced.

## Pillaging

Gangs of ten or more fighters may use the `PILLAGE` command to seize loot from a province or a city. Pillaging siezes the tax base for a location, leaving none available for taxation.

Pillaging has a harmful effect on the future tax revenue of the location. The more a location has been pillaged, the lower its tax base. Provinces and cities take four months to recover from each pillaging. For example, a city which was pillaged for five months in a row would take twenty months from the first pillaging to return to its normal tax base.

Pillagers must first defeat any units guarding the province, including any province garrisons.

