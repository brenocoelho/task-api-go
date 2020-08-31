# task-api

go run main.go
go build -o app 
go build -o app && ./app
go build && ./api



kubectl apply -f ./kubernetes/postgres/storageclass.yaml
kubectl apply -f ./kubernetes/postgres/storage.yml
kubectl apply -f ./kubernetes/postgres/configmap.yml
kubectl apply -f ./kubernetes/postgres/statefulset.yml
kubectl apply -f ./kubernetes/postgres/service.yml

kubectl create -f ./kubernetes/dashboard-admin.yaml
