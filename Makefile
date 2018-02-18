#! /usr/bin/make
ifeq ($(OS),Windows_NT)
	BUILD_TARGET_FILES = algtmapi.exe main.go
else
	BUILD_TARGET_FILES ?= algtmapi main.go
endif

DESIGN_PACKAGE = algtmapi/design
.DEFAULT_GOAL := prepare

all: cleandep depend clean mkdir precompile gen precompile build 

prepare: cleandep depend clean mkdir precompile gen precompile

depend:
	@dep ensure

cleandep:
	@rm -rf vendor

clean:
	@rm -rf app
	@ls swagger | grep -v -E "swaggerui$$" | xargs -I{} rm -rf swagger/{}
	@rm -rf assets

mkdir:
	@mkdir swagger/specs

confinit:
	@cp config/local.toml.example config/local.toml 
	@cp config/test.toml.example config/test.toml 

gen:
	@goagen app -d $(DESIGN_PACKAGE) 
	@goagen swagger -d $(DESIGN_PACKAGE) -o swagger
	@rm -rf swagger/specs
	@mv swagger/swagger swagger/specs
	@goagen controller -d $(DESIGN_PACKAGE) --pkg controller -o controller

precompile:
	@go-bindata -pkg=swaggerassets -o=swagger/swaggerassets/swagger.go swagger/specs/... swagger/swaggerui/...
	@go-bindata -pkg=assets -o=assets/bindata.go config/...

build:
	@go build -o $(BUILD_TARGET_FILES)

go-run:
	@go run main.go

run: precompile go-run

test:
	@go test -tags=test algtmapi/presentation/... -cover
