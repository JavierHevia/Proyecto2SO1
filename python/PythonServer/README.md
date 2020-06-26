# Servidor python Dockerfile kubernetes


Sistemas operativos 1
Eduardo Javier Hevia


# Pasos a seguir.


## Verificar si ya existe alguna imagen
sudo docker images


## Eliminar todas las imagenes creadas
sudo docker system prune
sudo docker system prune -a


## Remover contenedores
sudo docker ps  IdContenedor
sudo docker stop 
sudo docker rm


# Construir nuestra imagen 

sudo docker build -t servidorpython .

# Crear contenedor y se queda en consola
sudo docker run --rm -it -p 49160:8080 servidorpython
Se sale con Ctrl + c 
Miramos la IP de nuestra VM y le agregamos el puerto "49160"

# Subir nuestra imagen a Docker hub
docker login
	Nos pide nuestro usuario y nuestra contraseña
	
## Luego debemos de crear un tag en Dockerhub
	docker tag NombredeImagen USUARIODOCKERhub/NombredeImagen:v1
	
## Luego le damos push a nuestra imagen

	docker push USUARIODOCKERhub/NombredeImagen:v1
	
*Ir a revisar https://hub.docker.com/

#Crear kubernetes a partir de una imagen que esta en Docker hub

## Podemos ver los pods ya creados
	kubectl get pods

## Podemos ver la informacion detallada de los pods
	kubectl get rc
	kubectl get rc NOMBREDELPOD

## Creamos nuestro archivo python.yaml del POD

apiVersion: v1
kind: ReplicationController

metadata:
    name: my-sevpython
spec:
    replicas: 1
    selector:
        app: sevpython
    template:
        metadata:
            name: sevpython
            labels:
                app: sevpython
        spec:
            containers:
                - name: sevpython
                  image: javierhevia22/servidorpython #cambiar esto por lo que necesitamos
                  ports:
                    - containerPort: 8080
					
## Para crear nuestro POD desde .yaml le damos
	kubectl create -f python.yaml
	
## Miramos si esta creados
	kubectl get pods
	
## Miramos su puerto asignado aleatoriamente 
	kubectl get svc 
	
## Miramos la IP asignada
	kubectl describe nodes | grep ExternalIP
	
## Ver en NODE que se le fue asignado
	kubectl get pods -o wide
	
## Ver informacion del NODE
	kubectl describe nodes nameNode
	
# Para eliminar el POD anteriormente creado le damos
	kubectl delete rc my-nginx
	
#Creamos un servicio para exponer el POD pythonservice.yaml

apiVersion: v1
kind: Service
metadata:
    name: python-service
spec:
    type: NodePort
    selector:
        app: sevpython
    ports:
      - port: 8080
	  
## Creamos el servicio pythonservice.yaml
	kubectl create -f pythonservice.yaml
 
## Debemos de obtener la IP
	kubectl get pods -o wide
	kubectl describe nodes nameNode
		*buscamos ExternalIP para saber como vamos a acceder
	kubectl get svc
		*buscamos nuetro servicio y obtenemos el PORT 8080:30116 (30116) es el que se agrega con la ip antes obtenida
	


