# pingpong-envoy

to test envoyproxy and docker-compose with ping-pong

run : `docker-compose up` \
stop & remove : `docker-compose down`

structure : \
next (3000) <-> envoy-cli (10000) <-> envoy-service (8080) <-> service (50051)

more info : \
envoy cli - envoy client
