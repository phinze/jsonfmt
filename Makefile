TEST?=./...
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods -nilfunc -printf -rangeloops -shift -structtags -unsafeptr

default: test

# test runs the unit tests and vets the code
test:
	go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4
	@$(MAKE) vet

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) . ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for reviewal."; \
	fi

# bin generates the releaseable binaries for Jsonfmt
bin:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

# dev creates binaries for testing Jsonfmt locally. These are put
# into ./bin/ as well as $GOPATH/bin
dev:
	@JSONFMT_DEV=1 sh -c "'$(CURDIR)/scripts/build.sh'"


