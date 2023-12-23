##Các bước build image
docker run --name ThucHanh -e MYSQL_ROOT_PASSWORD=123 -p 3306:3306 -d mysql:latest

docker run --name bemount -itv F:\ThucHanh_Go\go-project:/app --network=my-net -p 8888:8888 be

docker run --name femount -itv F:\ThucHanh_Go\vue-project:/app -p 8080:8080 fe-app

docker run -i --network=my-net -t -p 8888:8888 be