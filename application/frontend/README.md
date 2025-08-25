- sec2
  - [gihyo-ms-dev-book/handson/sec2/2.2.3-frontend/frontend | GitLab](https://gitlab.com/gihyo-ms-dev-book/handson/sec2/2.2.3-frontend/frontend)
  - [gihyo-ms-dev-book/handson/sec2/2.2.4-deploy/frontend | GitLab](https://gitlab.com/gihyo-ms-dev-book/handson/sec2/2.2.4-deploy/frontend)


# プロジェクト作成

```bash
cd $PROJECT_DIR/application
npx create-react-app frontend
cd frontend

npm install @apollo/client graphql
```

# アプリのビルド

```bash
npm run build
```

# アプリの起動

```bash
(
  cd $PROJECT_DIR/application/frontend &&
  npm start
)

```

- http://localhost:3000 にアクセス