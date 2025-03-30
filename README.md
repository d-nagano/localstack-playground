# LocalStack検証

## 環境構築
1. リポジトリをクローン
   ```bash
   git clone https://github.com/d-nagano/localstack-playground.git
   cd localstack-playground
   ```

1. コンテナを起動
   ```bash
   docker compose build
   docker compose up -d
   ```

1. S3バケット作成
   ```bash
   awslocal s3api create-bucket --bucket sample-bucket
   ```

1. S3バケット確認
   ```bash
   awslocal s3api list-buckets
   ```
