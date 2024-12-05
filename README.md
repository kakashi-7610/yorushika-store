# yorushika-store
ヨルシカのグッズを販売するシステム（仮）

## store-controller
リクエストを受け付け処理を振り分ける

### docker-compose疎通
```
docker compose up -d
```

### kubernetes疎通
ubuntuの/home/shanks/配下にyorushika-storeを配置
```
minikube start
minikube mount /home/shanks/yorushika-store:/home/shanks/yorushika-store
minikube apply -k /kubernetes
minikube tunnel
```

## member-manager
会員管理を行うサービス

## product-manager
商品管理を行うサービス

## clearing-manager
決済管理を行うサービス