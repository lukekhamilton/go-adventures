.PHONY: deps clean build

ts := $(shell /bin/date "+%Y%m%d")
tss := $(shell /bin/cat .tmp )

APP=hello-sam
BUCKET=$(APP)-$(tss)

bucket:
	echo $(BUCKET)

deps:
	go get -u ./...

clean:
	rm -rf ./hello-world/hello-world

build:
	# TODO Fix naming?
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world

test:
	go test ./...

run: build
	sam local start-api

bootstrap:
	echo $(ts) > .tmp
	aws s3 mb s3://$(BUCKET)

test:
	curl http://127.0.0.1:3000/hello

destroy: package
	aws s3 rb s3://$(BUCKET)

package:
	sam package \
		--template-file template.yaml \
		--output-template-file package.yaml \
		--s3-bucket "$(BUCKET)" \

deploy: package
	sam deploy \
		--template-file package.yaml \
		--stack-name $(APP) \
		--capabilities CAPABILITY_IAM

