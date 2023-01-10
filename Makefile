VERSION ?= latest

# Create a variable that contains the current date in UTC
# Different flow if this script is running on Darwin or Linux machines.
ifeq (Darwin,$(shell uname))
	NOW_DATE = $(shell date -u +%d-%m-%Y)
else
	NOW_DATE = $(shell date -u -I)
endif

all: test

.PHONY: test
test:
	go test ./... -coverprofile coverage.out

.PHONY: build
build:
	go build ./cmd/config-shepherd/...

.PHONY: version
version:
	sed -i.bck "s|VERSION=\"[0-9]*.[0-9]*.[0-9]*.*\"|VERSION=\"${VERSION}\"|" "Dockerfile"
	rm -fr "Dockerfile.bck"
	git add "Dockerfile"
	git commit -m "Upgrade version to v${VERSION}"
	git tag v${VERSION}
