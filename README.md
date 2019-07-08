- [gurume ETL](#gurume-ETL)
  - [ETL overview](#ETL-overview)
  - [nogada](#nogada)
    - [transform `gurume list` to json format](#transform-gurume-list-to-json-format)
    - [`gurume.json` on elasticsearch](#gurumejson-on-elasticsearch)
  - [elastic cloud](#elastic-cloud)
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

## WARN!! - ZERO WIDTH SPACE, must handle (U+200B)

## generate processed txt
## 1. build images (you can skip it when using local go env)
docker-compose build gurume 

## 2. process gurume.txt -> gurume.processed.1.txt
### local env case
go run main.go formatData --file gurume.txt

### docker-compose env case
docker-compose run --rm gurume formatData --file gurume.txt
```

### transform `gurume list` to json format

```s
### local env case
go run main.go formatJSON

### docker-compose env case
docker-compose run --rm gurume formatJSON

## 3. check file
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
docker-compose build elasticsearch2 elasticsearch3 elasticsearch

## cluster up
docker-compose up -d elasticsearch2 elasticsearch3 elasticsearch

## mapping check
curl localhost:9200/gurume_index/_mapping | jq

## search test
curl \
 -H 'Content-Type: application/json'\
 -X POST 'localhost:9200/gurume_index/gurume/_search'\
 --data '{ "from": 0, "size": 30, "query" : { "match" : { "category" : "닭곰탕" } }}' | jq '.hits.hits[]._source.category'

curl \
 -H 'Content-Type: application/json'\
 -X POST 'localhost:9200/gurume_index/gurume/_search'\
 --data '{ "from": 0, "size": 30, "query" : { "match" : { "station" : "을지로 4가역" } }}' | jq '.hits.hits[]._source.station'

```

## elastic cloud
- https://cloud.elastic.co

## Backend - elasticsearch client (golang)
- TBU

## Frontend - (vue)
- TBU