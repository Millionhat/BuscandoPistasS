# BuscandoPistasS

Se me olvido crear el repo antes de empezar a hacer este proyecto

Para iniciar el proyecto se requiere inicializar el nodo de la db, inicializar el server/Api Go e instalar los paquetes de vue


Instalacion e inicializacion paquetes de vue:
en la carpeta del proyecto con vue.js instalado utilizar los comandos siguientes:
cd frontend ,
npm install,
npm run serve


Inicializar nodo de db:
en la carpeta del proyecto con cockroachDB instalado utilizar el siguiente comando en un consola diferente a la de vue:
cockroach  start --insecure --listen-addr=localhost:8081 --http-addr=localhost:3000

Inicializar Go Api:
en una consola diferente a la del nodo y la de vue ir a la carpeta del proyecto
cd backend,
cd main,
go run main.go

