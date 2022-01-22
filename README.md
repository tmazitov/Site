# VueGo site
This project is an example of a website with backend and frontend parts. The backend was written on Golang with help http standard library and httprouter. There I use clean architecture at the core. The frontend is based on Vue2.

In production i use docker and nginx-proxy. 

## Develop

### Start backend
```bash
cd app
```
```bash
go run main.go
```

### Start frontend
```bash
cd web
```
```bash
yarn serve
```
---
## Production

You must have Docker installed on your computer.

```bash
docker-compose build
```

```bash
docker-compose up -d
```