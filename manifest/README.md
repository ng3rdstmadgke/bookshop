# デプロイ手順


## イメージのビルドとプッシュ


```bash
VERSION=0.1
build.sh -v $VERSION -p
```


## デプロイ

```bash
aws eks update-kubeconfig --name baseport-prd
```

```bash
kubectl apply -f $PROJECT_DIR/manifest/ns.yaml

kubectl apply -f $PROJECT_DIR/manifest/catalogue/k8s/catalogue.yaml
kubectl apply -f $PROJECT_DIR/manifest/bff/k8s/bff.yaml
kubectl apply -f $PROJECT_DIR/manifest/frontend/k8s/frontend.yaml

kubectl apply -f $PROJECT_DIR/manifest/ingress.yaml
```

## アクセス

https://bookshop.prd.baseport.net/