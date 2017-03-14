# testgoserv [![Build Status](https://travis-ci.org/guzenok/testgoserv.svg?branch=develop)](https://travis-ci.org/guzenok/testgoserv)

Simple http-server "go_serv" in docker-image [guzenok/testgoserv](https://hub.docker.com/r/guzenok/testgoserv/), and docker-compose.yml to try it with nginx balancer.

### Example:

Commands
```
docker-compose up -d
for i in `seq 0 9`; do echo "On request $i balancer return \"$(curl http://localhost:8080 2>/dev/null)\""; done
docker-compose down
```
output
```
Starting testcicd_app1_1
Starting testcicd_app2_1
Starting testcicd_balancer_1

On request 0 balancer return "App-Node-1 get 1-nd response"
On request 1 balancer return "App-Node-2 get 1-nd response"
On request 2 balancer return "App-Node-1 get 2-nd response"
On request 3 balancer return "App-Node-1 get 3-nd response"
On request 4 balancer return "App-Node-2 get 2-nd response"
On request 5 balancer return "App-Node-1 get 4-nd response"
On request 6 balancer return "App-Node-1 get 5-nd response"
On request 7 balancer return "App-Node-2 get 3-nd response"
On request 8 balancer return "App-Node-2 get 4-nd response"
On request 9 balancer return "App-Node-1 get 6-nd response"

Stopping testcicd_balancer_1 ... done
Stopping testcicd_app2_1 ... done
Stopping testcicd_app1_1 ... done

```