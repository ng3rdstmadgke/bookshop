- [gihyo-ms-dev-book/handson/sec2/2.2.2-bff/bff | GitLab](https://gitlab.com/gihyo-ms-dev-book/handson/sec2/2.2.2-bff/bff)

# プロジェクト作成

```bash
cd $PROJECT_DIR/application/bff
npm init --yes
npm install @apollo/server express graphql cors body-parser @grpc/grpc-js @grpc/proto-loader
```

# GraphQLを起動

```bash
(
  cd $PROJECT_DIR/application/bff
  npm start
)
```

http://localhost:4000/graphql にアクセス


# クエリを投げてみる

## idを指定してBookを取得

Operation

- `query`
  - `Book` : クエリ名 (任意)
  - `($bookId: Int)` : このクエリ内で利用される変数定義
    - `book` : resolver.jsで定義されたクエリフィールド
    - `(id: $bookId)` : `$bookId` を `id` パラメータに渡す
      - `id, title, author, price` : レスポンスとして取得する項目

```
query Book($bookId: Int) {
  book(id: $bookId) {
    id
    title
    author
    price
  }
}
```

Variables

```
{
  "bookId": 1
}
```

Response


```
{
  "data": {
    "book": {
      "id": 1,
      "title": "The Awakening",
      "author": "Kate Chopin",
      "price": 1000
    }
  }
}
```


## すべてのBookを取得


Operation

```
query Books {
  books {
    id
    title
    author
    price
  }
}
```

Response


```
{
  "data": {
    "books": [
      {
        "id": 1,
        "title": "The Awakening",
        "author": "Kate Chopin",
        "price": 1000
      },
      {
        "id": 2,
        "title": "hogehoge",
        "author": "abcdefg",
        "price": 2000
      }
    ]
  }
}
```