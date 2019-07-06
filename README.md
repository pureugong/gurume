- [gurume ETL](#gurume-ETL)
  - [ETL overview](#ETL-overview)
  - [nogada](#nogada)
    - [extract `gurume category`... !](#extract-gurume-category)
    - [extract `gurume list` from gurume.txt](#extract-gurume-list-from-gurumetxt)
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

```json
// gurume.json example ?
[
    {
        "name": "이레김밥",            // es mapping type : text ?
        "category": ["김밥"],         // es mapping type : text ?
        "station": "낙성대역",         // es mapping type : keyword
        "town": "인헌동",             // es mapping type : keyword
        "instagram": "https://xxxx", // es mapping type : keyword
        "google_map": "https://xxxx" // es mapping type : keyword
    },
    //...
]
```

## nogada
### extract `gurume category`... !

```s
## 식당 카테고리 만 grep (숫자로 시작하고 + 쩜이 있고 + 스페이스 + 텍스트)
grep "^\d*\.\ *" gurume.txt

## 식당에도 숫자가...있네.. - 있는 라인 제거
grep "^\d*\.\ *" gurume.txt | grep -v '-'

## 카테고리 + 밑에 데이터 있는지 확인.. 그런데 56개..만.. 57-마지막까지는 데이터가 없다..
## 다시 오리지널 데이터 입수
https://monk4.tistory.com/2

## 93개 카테고리 추출
grep "^\d*\.\ *" gurume.txt | grep -v '-' | head -n 93

## `68. 복어`가 카테고리에서 누락.. 직접 타이핑해서 추가

## gurume_category.txt 파일생성
grep "^\d*\.\ *" gurume.txt | grep -v '-' | head -n 93 > gurume-category.txt 

## data verification - head
head gurume-category.txt 
1. 평양냉면
2. 메밀국수 (소바)
3. 막국수
4. 콩국수
5. 국밥, 해장국
6. 설렁탕
7. 감자탕
8. 순대
9. 닭볶음탕
10. 추어탕

## data verification - tail
tail gurume-category.txt 
84. 참치
85. 랍스터, 킹크랩, 대게
86. 남도 음식 (민어, 홍어, 병어, 낙지, 보리굴비, 남도한정식)
87. 횟집 (세꼬시, 과메기, 물회, 막회, 해산물, 꽃새우)
88. 방송맛집
89. 서울 노포 식당
90. 떡집
91. 디저트
92. 빵집
93. 커피
```

### extract `gurume list` from gurume.txt
- TBU

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
