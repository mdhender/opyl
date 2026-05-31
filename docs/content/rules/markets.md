---
title: Markets
weight: 5
toc: true
---

## Markets

Characters issue BUY and SELL commands to indicate their desire to trade goods. Every trade must be between a buyer and a seller. Whenever a compatible BUY and SELL request are found, the trade will be executed.

Trades are only matched in cities, where merchants gather at the local bazaar to swap goods. (Note that characters are always free to use the give command to exchange items, regardless of where they are. But BUY and SELL orders are only matched in cities.)

A buyer indicates what he wants to buy, how much of it he wants, and the most he is willing to pay. A seller indicates what item is being sold, how much of it is to be sold, and the per-item price.

A trade will match if the seller's price does not exceed the maximum price specified by the buyer; if the buyer has enough gold to buy at least one of the item; if the seller possesses at least one of the item; and if the buyer and seller are both in the same city.

If the buyer is willing to pay more than the seller is asking, the trade takes place at the seller's price.

Example: A noble who wants to buy five iron [79] at no more than 10 gold each issues:

```report
> buy 79 5 10

Try to buy five iron [79] for 10 gold each.
```

If someone who had iron issues a sell order which matches this buyer's request, the trade would be executed:

```report
> sell 79 5 10

Sold five iron [79] for 50 gold.
```

The buyer would see:

```report
Bought five iron [79] for 50 gold.
```

Either the buy or the sell order could have been issued first. If the order can't be matched with pending trades from other units in the city, it will become a standing order and remain in effect until executed or canceled.

```report
> buy 79 0

Cleared pending buy for iron [79].
```

Note that the buyer and seller can't specify what character they will deal with. They will trade with any unit that has a matching buy or sell order.

### Market report

The location report for each city includes a market report listing pending trades.

```report
| trade | who  | price | qty | item             |
| ----- | ---- | ----- | --- | ---------------- |
| buy   | 2019 | 100   | 1   | peasant [10]     |
| buy   | 4274 | 74    | 3   | iron [79]        |
| sell  | 3682 | 12    | 11  | sailor [19]      |
| sell  | 2019 | 50    | 1   | elite guard [15] |
```

A trade will not be listed unless it could be executed. For instance, a unit might issue an order to SELL iron [79], even though the unit doesn't possess any. (Perhaps the character plans to get some iron later in the month, and wants the buy order to be in place when the iron arrives). This order will not be shown in the market report, because the seller doesn't have any iron to sell.

Similarly, the buyer must have enough gold to buy at least one of the item. No buy order would be listed for a penniless unit that wanted to buy five iron at 10 gold each. If the unit later obtained 10 gold, enough to buy one of the five desired units of iron, the order would be listed in the market report for one iron, not five.

Some traders may work through middlemen to hide their identity. In such cases, the who field of the market report will not show their unit number, and their identity will not be revealed when the trade is executed.

Even sneakier traders may have their pending trades omitted from the market report entirely.

### City purchasing

Some cities will issue BUY or SELL orders on their own behalf for certain goods. Cities trade orders are identity to those submitted by players when participating in markets.

Enterprising characters may turn a profit by using skills under the Trade [730] category to find new tradegoods for sale in city markets, and buyers for those tradegoods in other city markets.

### Resolution of trades

If a buyer has a choice between two or more sellers, the one offering the lowest price will be chosen. If several sellers offer at the same price, the one nearest the top of the location's character listing will win the trade. A seller with a choice among two or more buyers will pick the highest unit as it appears in the location's character listing.

Since characters are added to the end of the list when they enter a location, the units who have been in a place the longest tend to appear toward the top of the list. Characters who have been in a place the longest have an advantage when one of several possible trades may be matched.

However, consumption or production by the city itself will always have lowest priority. Cities defer to characters when multiple matching trades are possible.

The location market report is ordered according to how multiple trades will resolve.

```report
| trade | who  | price | qty | item      |
| ----- | ---- | ----- | --- | --------- |
| sell  | 3682 | 10    | 1   | iron [79] |
| sell  | 2019 | 12    | 1   | iron [79] |
```

A buyer would get 3682's iron, since the price is lower.

```report
| trade | who  | price | qty | item      |
| ----- | ---- | ----- | --- | --------- |
| sell  | 2019 | 10    | 1   | iron [79] |
| sell  | 3682 | 10    | 1   | iron [79] |
```

A buyer would get 2019's iron. 2019 must appear before 3682 in the location character list.

```report
| trade | who  | price | qty | item      |
| ----- | ---- | ----- | --- | --------- |
| buy   | 4846 | 10    | 1   | iron [79] |
| buy   | 1783 | 12    | 1   | iron [79] |
```

4846 would win the buy from a lone seller. 1783 must come after 4846 in the location character list.

```report
| trade | who  | price | qty | item      |
| ----- | ---- | ----- | --- | --------- |
| buy   | 4846 | 10    | 1   | iron [79] |
| buy   | 1783 | 11    | 1   | iron [79] |
```

`sell 79 1 10` will match 4846. `sell 79 1 11` will match 1783.

### Partial trades

Partial trades are executed, if possible. If a buyer wants 100 iron, but only 25 iron are available, he will buy 25. The pending buy order will be reduced so that 75 more will be bought whenever possible.

### More market examples

A buyer who desires three iron at no more than five gold each orders:

```report
> buy 79 3 5

Try to buy three iron [79] for 5 gold each.
```

Suppose the only seller of iron in this city had ordered:

```report
> sell 79 5 6

Try to sell five iron [79] for 6 gold each.
```

The trade will not take place, because the seller's price exceeds the buyer's. Later, the buyer travels to a different city, where another seller has previously issued:

```report
> sell 79 10 4

Try to sell ten iron [79] for 4 gold each.
```

As soon as the buyer enters the city, the trade will be matched:

```report
Bought three iron [79] for 12 gold.
```

Note that the trade takes place at the seller's price, and that the seller will still be offering seven iron at four gold each.

If the seller had only one unit of iron for sale, the trade would have executed, but the buyer would still be looking for two more iron to buy, at five gold each or better.

