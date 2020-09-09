.PHONY:run
run:
	clear
	sh ./scripts/banner.sh
	@echo "------------------"
	go run main.go $(ARGS)

.PHONY:push
push:
	clear
	git add .
	git commit -m "$(COMMIT)"
	git push origin master