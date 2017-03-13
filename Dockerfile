FROM scratch
COPY go_serv /
ENTRYPOINT ["/go_serv"]

EXPOSE 8080
