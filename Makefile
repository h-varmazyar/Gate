GOPATH=${HOME}/go

%:
	@true

.PHONY: fmt
fmt:
	./scripts/fmt.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: test
test:
	./scripts/test.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: run
run:
	./scripts/fmt.sh $(filter-out $@,$(MAKECMDGOALS))
	./scripts/run.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: proto
proto:
	./scripts/proto.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: infra
infra:
	./scripts/infra.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: install
install:
	./scripts/install.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: build
build:
	./scripts/build.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: lint
lint:
	./scripts/lint.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: service
service:
	./scripts/service.sh $(filter-out $@,$(MAKECMDGOALS))

