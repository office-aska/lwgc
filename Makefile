FIRESTORE_REPO_VERSION := v1.10.0
OS_NAME := `echo $(shell uname -s) | tr A-Z a-z`
MACHINE_TYPE := $(shell uname -m)

.PHONY: bootstrap_firestore_repo
bootstrap_firestore_repo:
	mkdir -p bin
	curl -L -o ./bin/firestore-repo.tar.gz https://github.com/go-generalize/firestore-repo/releases/download/$(FIRESTORE_REPO_VERSION)/firestore-repo_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf firestore-repo.tar.gz && rm *.tar.gz

.PHONY: run_server
run_server:
	go run app/*.go

.PHONY: prune
prune:
	git remote prune origin
	git fetch origin --prune 'refs/tags/*:refs/tags/*'
	git branch --merged | egrep -v '^ *(develop|master|main|\*)' | xargs git branch -d
