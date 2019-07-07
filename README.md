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

## WARN!! - ZERO WIDTH SPACE, must handle (U+200B)

## ... continue

```

### transform `gurume list` to json format

```s
## generate go cmd
go run main.go formatData --file gurume.txt

## check file
head -n2 data/gurume.processed.1.json | jq
{
  "category": "평양냉면",
  "station": "을지로 3가역",
  "town": "입정동",
  "name": "을지면옥"
}
{
  "category": "평양냉면",
  "station": "압구정역, 학동역",
  "town": "논현동",
  "name": "논현동 평양면옥"
}
```

### `gurume.json` on elasticsearch
```sh
## build es
docker-compose build

## cluster up
docker-compose up -d elasticsearch2 elasticsearch3 elasticsearch

## mapping check
curl localhost:9200/gurume_index/_mapping | jq

## search test
curl \
 -H 'Content-Type: application/json'\
 -X POST 'localhost:9200/gurume_index/gurume/_search'\
 --data '{ "from": 0, "size": 30, "query" : { "match" : { "category" : "곰탕" } }}' | jq '.hits.hits[]._source.category'

```

## AWS elasticsearch
- TBU

## Backend - elasticsearch client (golang)
- TBU

## Frontend - (vue)
- TBU