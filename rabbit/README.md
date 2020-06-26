docker login 
docker build -t golb .
docker tag golb javierhevia/serviciogo
docker push javierhevia/serviciogo


kubectl create deployment golb -n proyecto --image=javierhevia/serviciogo
kubectl -n proyecto expose deployment golb --port 8080 --target-port=8080 --type NodePort --name=golb-service
kubectl -n proyecto apply -f go-ingress.yml


docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management