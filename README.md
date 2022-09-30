Run MongoDB using Docker
- docker run -e MONGO_INITDB_ROOT_USERNAME=<username> -e MONGO_INITDB_ROOT_PASSWORD=<password> -p 27020:27020 --expose 27020 -d --name mongodb mongo:6.0.1

Build Image from Dockerfile:
- docker build -f Dockerfile <image-name:tag> .

Run Backend using Docker (?):
- docker run -e MONGO_DB_URI=mongodb://admin:yadayada12345@localhost:27017 -e MONGO_DB_NAME=go_pos -p 8081:8081 --expose 8081 --name backe
nd test:test