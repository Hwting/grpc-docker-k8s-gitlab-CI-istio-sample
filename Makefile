.PHONY: client server clean all docker dockerc dockers deploy cleandeploy

# set version and build
VERSION=$(shell git rev-parse --short=8 HEAD)
BUILD=$(shell date +%FT%T%z)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BINCLIENT=client
BINSERVER=server
# backup, need add VERSION and BUILD  variables in main func.
# LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

all: client server
docker: dockerc dockers

client:
	go build -o ./client/${BINCLIENT} ./client/

server:
	go build -o ./server/${BINSERVER} ./server/

dockerc:
	@cat ./client/Dockerfile | docker build -f - -t jacenr/client:${VERSION} .
	docker push jacenr/client:${VERSION}

dockers:
	@cat ./server/Dockerfile | docker build -f - -t jacenr/server:${VERSION} .
	docker push jacenr/server:${VERSION}

deploy:
	@cat ./k8s.yaml | sed 's/$$VERSION/${VERSION}/g' | kubectl apply -f -

cleandeploy:
	@cat ./k8s.yaml | sed 's/$$VERSION/${VERSION}/g' | kubectl delete -f -

clean:
	rm -f ./client/${BINCLIENT} ./server/${BINSERVER}
	docker image rm -f jacenr/client:${VERSION}
	docker image rm -f jacenr/server:${VERSION}
