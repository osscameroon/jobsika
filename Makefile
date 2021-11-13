.DEFAULT_GOAL := help

## install-hooks: install hooks
.PHONY: install-hooks
install-hooks:
	ln -s $(PWD)/githooks/pre-push .git/hooks/pre-push

all: help
.PHONY: help
help: Makefile
	@echo " Choose a command..."
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
