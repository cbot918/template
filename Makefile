# connect to db
DOCKER_NETWORK="infra_default"
psql:
	docker run -it --rm --network=$(DOCKER_NETWORK) postgres psql -h postgres -U postgres

.PHONY: psql