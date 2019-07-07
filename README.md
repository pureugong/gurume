- [gurume ETL](#gurume-ETL)
  - [ETL overview](#ETL-overview)
  - [nogada](#nogada)
    - [transform `gurume list` to json format](#transform-gurume-list-to-json-format)
    - [`gurume.json` on elasticsearch](#gurumejson-on-elasticsearch)
  - [AWS elasticsearch](#AWS-elasticsearch)
  - [Backend - elasticsearch client (golang)](#Backend---elasticsearch-client-golang)
  - [Frontend - (vue)](#Frontend---vue)

# gurume ETL

## ETL overview
- extract `gurume category` from gurume.txt
- extract `gurume list` from gurume.txt
- transform `gurume list` to json format like below ?? a.k.a `gurume.json`
- load `gurume.json` on elasticsearch

```js
// gurume.json example ?
[
    {
        "name": "이레김밥",            // es mapping type : text ?
        "category": ["김밥"],         // es mapping type : text ?
        "station": "낙성대역",         // es mapping type : keyword
        "town": "인헌동",             // es mapping type : keyword
        "instagram": "https://xxxx", // es mapping type : keyword
        "google_map": "https://xxxx" // es mapping type : keyword
        "remark": "xxxxxxx"          // es mapping type : text ?
    },
    //...
]
```

## nogada

```s
## check exception case
go run main.go gurume.txt | grep -v 'info\|review\|hotel' | grep exception

## expcetion case update gurume.txt

## add 노포식당 handling

## generate processed txt
go run main.go gurume.txt > gurume.processed.1.txt

```

### transform `gurume list` to json format
- TBU

### `gurume.json` on elasticsearch
- TBU

## AWS elasticsearch
- TBU

## Backend - elasticsearch client (golang)
- TBU

## Frontend - (vue)
- TBU



grep -E "\s\W{1,3}동\s\-" gurume-list.txt | sed 's/.*/\n&/'





sed 's/.*/\n&/g'