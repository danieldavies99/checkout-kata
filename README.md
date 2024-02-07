# Daniel Davies Checkout Kata
This repo contains the code for my solution to the Think Money Checkout Kata
Assignment url: https://github.com/thinkmoney/checkout-kata

# Brief Explanation
I decided to store the pricing data in a .json file and load it in rather than hard code it in GoLang because I figured that in a real life application this data would probably come from an API request or a local DB. I wrote this following TDD, I added the happy path tests first and then decided to add some error handling as well.

If you look at the commit history you'll see that I decided to refactor my pricing module after I already had a working solution. In my first pass I stored all the pricing information in a slice like this:

```
[
    {
        Sku: "A",
        UnitPrice: 10,
        MultiBuyPrice: 25,
        MultiBuyCount: 3,
    },
    ...
]
```

I then iterated through the priceList and matched scanned items in the ```getTotalPrice()``` function. This was problematic because as the priceList increases in size, the for loop in ```getTotalPrice()``` would have to iterate through more items even if they weren't scanned. So I refactored the pricelist to the be stored as a key-value pair map instead. This means that we don't have to worry about having a really big pricelist as we can address the prices we need directly (I appreciate that in this example the largest price list is 26 because there's only 26 possible SKUs, but my solution could work with any string sku, E.g. 'AAAAA' which I figured was a more realistic use-case)

Steps to run:
1. Clone repo and cd into project
2. Run ```go run .```

To run tests:
1. Clone repo and cd into project
2. Run ```go test ./... -v```

OR

Use a test runner like VsCode's test runner (this is what I normally do)
![VsCode Test Pass](https://i.imgur.com/kvdAXSb.png)