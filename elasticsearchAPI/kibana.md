# Kibanaを利用してElasticsearchを操作する

## Management: データを挿入する

~~~
POST namespace/data/1
{
    "id":1,
    "code":200,
    "date":"2021-01-01T12:34:56Z",
    "message":"XXX_SOFT:INFO:00000:Process completed",
}

POST namespace/data/2
{
    "id":2,
    "code":403,
    "date":"2021-01-02T12:34:56Z",
    "message":"XXX_SOFT:ERROR:10000:Trouble occurred",
}
~~~

## Discover: 検索式を作成する

* 下記にアクセスする

~~~
[Management] > [Index Patterns]
~~~

* `[Index name or pattern]` に Index名(上記の例ではnamespace)を入力する
* `[Time-field name]` に自動検出された時刻フィールドが表示されることを確認する
* `[Create]` を押下する
* Index Patternが作成されたことを確認する
* `[Discover]` を押下する


## Visualize: 表示方法を決めてグラフ化する

## Dashboard: データを素早く見る
