.PHONY: run
run: build
	docker-compose -f docker-compose.yml --profile backend up --build

.PHONY: runarm
runarm: buildarm
	docker-compose -f docker-compose.yml --profile backend up --build

.PHONY: build
build:
	cd service.github && sudo make docker && cd .. && \
	cd service.openai && sudo make docker && cd .. && \
	cd service.agentsmith && sudo make docker & cd ..

.PHONY: buildarm
buildarm:
	cd service.openai && sudo make dockerarm && cd .. && \
	cd service.github && sudo make dockerarm && cd .. && \
	cd service.agentsmith && sudo make docker & cd ..

.PHONY: run-infra
run-infra:
	docker-compose -f docker-compose.yml --profile infra up --build
