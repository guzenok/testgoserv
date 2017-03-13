build: go_serv
	docker build -t testgoserv .

go_serv:
	CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o go_serv ./app/

test:
	go test ./app/