.PHONY:run
run:
	clear
	sh ./scripts/banner.sh
	@echo "------------------"
	go run main.go $(ARGS)

.PHONY:build
build:
	go build

.PHONY:global
global: build
	sudo cp licensor /usr/local/bin
	rm licensor

.PHONY:keep-binary
keep-binary: build
	sh ./scripts/keep-binary.sh

.PHONY:push
push:
	clear
	git add .
	git commit -m "$(COMMIT)"
	git push origin master

.PHONY:test
test: test-env
	sh ./scripts/test.sh
	make drop-test-env

.PHONY:test-env
test-env:
	clear
	sh ./scripts/test-env.sh

.PHONY:drop-test-env
drop-test-env:
	MODE="DROP" sh ./scripts/test-env.sh