bourbon-finder
===============

A tool for finding bourbon (or any product) from Virginia ABC stores. This allows you to
build a wishlist of hard to find bourbons and search the stores you are close to.

Requires Go 1.10 or later.


## Building

    git clone https://github.com/austinkeeley/bourbon-finder.git
    export GOPATH=$(pwd)
    make

## Usage

Create a json file that contains the stores and products you want to search. 
See the example below.

For stores, the store `name` is just for readability pourposes. The `storeID` 
is the actual ID to search on.

For the wishlist, the `productCode` is the actual ID to search for and the name
is for display purposes. Note that different product sizes have different 
`productCode` values.

Example (using `data/test_config.json`):

```json
{
    "stores": [
        {
            "name": "Kempsville Road",
            "storeID": 437
        },
        {
            "name": "South Battlefield Blvd.",
            "storeID": 349

        }
    ],

    "wishlist": [
        {
            "name": "Eagle Rare 10 Year Bourbon 375 ml",
            "productCode": "017764"
        },
        {
            "name": "Blanton's Single Barrel Bourbon 375 ml",
            "productCode": "016845"
        }
    ]
}
```

Pass this file as an argument with `-c` 

```
cmd/bourbon-finder/bourbon-finder -c data/test_config.json
```

There is also a web server version that can be started with `-w`. This will render the
results as HTML and serve the content on port 5001.

```
cmd/bourbon-finder/bourbon-finder -c data/test_config.json -w
```

