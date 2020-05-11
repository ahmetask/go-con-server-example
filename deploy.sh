go mod vendor
docker build -t go-con-server .
kubectl delete deployment go-con-server
kubectl delete service go-con-server
kubectl apply -f deployment.yml
kubectl apply -f service.yml
kubectl apply -f role.yml
kubectl apply -f rolebinding.yml
kubectl apply -f cluster_role.yml