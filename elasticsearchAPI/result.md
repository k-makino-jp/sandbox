# ERROR Log

~~~
> curl --request GET --url http://localhost:9200/library/_analyze --header 'content-type: application/json' --header 'user-agent: vscode-restclient' --data '{"tokenizer": "letter","text": "2021/01/01T12:34:56Z:SOFTWARE:ERROR:00000:resident"}' | jq
{
  "tokens": [
    {
      "token": "T",
      "start_offset": 10,
      "end_offset": 11,
      "type": "word",
      "position": 0
    },
    {
      "token": "Z",
      "start_offset": 19,
      "end_offset": 20,
      "type": "word",
      "position": 1
    },
    {
      "token": "SOFTWARE",
      "start_offset": 21,
      "end_offset": 29,
      "type": "word",
      "position": 2
    },
    {
      "token": "ERROR",
      "start_offset": 30,
      "end_offset": 35,
      "type": "word",
      "position": 3
    },
    {
      "token": "resident",
      "start_offset": 42,
      "end_offset": 50,
      "type": "word",
      "position": 4
    }
  ]
}
~~~