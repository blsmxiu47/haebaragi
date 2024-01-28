# haebaragi API

## Description

Haebaragi is a flashcards-like app, focused on language learning.

[API Spec](https://github.com/blsmxiu47/haebaragi/blob/main/openapi/openapi.json)

## Test Examples

Show welcome message and help:

```bash
curl http://localhost:8080
```

GET all words:

```bash
curl http://localhost:8080/word
```

Make a POST request to add a new word:

```bash
curl -X POST -H "Content-type: application/json" -d @test_data/POST_word.json http://localhost:8080/word
```
