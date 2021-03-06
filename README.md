- [gurume ETL README](#gurume-ETL-README)
  - [ETL overview](#ETL-overview)
- [TODO](#TODO)
  - [nogada](#nogada)
    - [transform `gurume list` to json format](#transform-gurume-list-to-json-format)
    - [`gurume.json` on elasticsearch](#gurumejson-on-elasticsearch)
  - [elastic cloud (optional)](#elastic-cloud-optional)
  - [Backend - elasticsearch client (golang)](#Backend---elasticsearch-client-golang)
  - [Frontend - (vue)](#Frontend---vue)
    - [TODO](#TODO-1)

# gurume ETL README

## ETL overview
- extract `gurume category` from gurume.txt
- extract `gurume list` from gurume.txt
- transform `gurume list` to json format like below ?? a.k.a `gurume.json`
- load `gurume.json` on elasticsearch

```js
// gurume.json example (v0.0.1)
[
  {
    "category": [
        {"name": "소고기"},
        {"name": "숙성 고기집"}
    ],
    "town": "서초동",
    "station": [
        {"name": "강남역"}
    ],
    "name": "어사담",
    "note": "드라이에이징"
  },
]
```

# TODO
- [x] ETL pipeline
- [x] design ES mapping
- [x] ES cloudsetup
  - [x] add user dictionary (category, station, town)
  - [ ] create api client role
- [x] Backend ES client
- [ ] Frontend app
- [ ] AWS ECS setup
- [ ] jenkins build / deploy pipeline

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

## bulk request to ES
docker-compose run --rm gurume ingestES

## mapping check
curl localhost:9200/gurume_index/_mapping | jq

## search test
curl \
 -H 'Content-Type: application/json'\
 -X POST 'localhost:9200/gurume_index/gurume/_search'\
 --data '{ "from": 0, "size": 30, "query" : { "match" : { "category.name" : "닭곰탕" } }}' | jq '.hits.hits[]._source.category'

curl \
 -H 'Content-Type: application/json'\
 -X POST 'localhost:9200/gurume_index/gurume/_search'\
 --data '{ "from": 0, "size": 30, "query" : { "match" : { "station.name" : "을지로 4가역" } }}' | jq '.hits.hits[]._source.station'

```

## elastic cloud (optional)
- https://cloud.elastic.co
- create ES cluster, then update `.env` file accordingly
```sh
# .env example
GURUME_ENV=production
ES_CLUSTER_HOST=https://xxxxxxxxxxxxxx.ap-northeast-1.aws.found.io
ES_CLUSTER_PORT=9200
ES_CLUSTER_USER_ID=hoge
ES_CLUSTER_USER_PW=hoge
LOG_LEVEL=info
```

- ingest data to ES cluster
```sh
## bulk request to ES
docker-compose run --rm gurume ingestES
```

## Backend - elasticsearch client (golang)
- API check in local env
```sh
### local env case
go run main.go api

### docker-compose env case
docker-compose run --rm gurume api
```

- Docker image build and push (manual)
```sh
## FYI, it will be automated by build pipeline
docker build -t pureugong/gurume:latest  .
$(aws --profile pureugong-gurume  ecr get-login --no-include-email)
docker tag pureugong/gurume:latest {aws-ecr-host}/{ecr-repo-name}:{version}
docker push {aws-ecr-host}/{ecr-repo-name}:{version}
```

## Frontend - (vue)

### TODO
- [ ] S3 bucket
- [ ] routing
- [ ] build pipeline
- [ ] autocomplete tag