all: api build up

api:
	docker run --rm \
		-v $(shell pwd -P):/go/src/gitlab.com/tonyhb/keepupdated \
		-v $(shell pwd -P)/cmd/api:/go/bin \
		golang:1.7.1-alpine \
		go install gitlab.com/tonyhb/keepupdated/cmd/api


build:
	docker-compose -f compose-dev.yml -p ku build

up:
	docker-compose -f compose-dev.yml -p ku up -d

logs:
	docker-compose -f compose-dev.yml -p ku logs -f

clean:
	rm cmd/api/api


.PHONY: up logs
