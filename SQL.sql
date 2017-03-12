Select All Stuff:

```

select id, timestamp, price, volume, total_transaction_cost from temptable;


```

Select transaction with largest total_transaction_cost:
(Does not mean this was largest transaction in terms of volume!!!)
```

select max(total_transaction_cost), id, timestamp from temptable;

```

Select transaction with largest volume:

```

select max(volume), id, timestamp from temptable;


```
