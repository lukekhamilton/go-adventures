.PHONY: deps clean build

# Generate a timestamp in memory
ts := $(shell /bin/date "+%Y%m%d")

# Get generated timestamp from file
tss := $(shell /bin/cat .tmp )

# API Gateway deployment id for building the url
id := $(shell aws apigateway get-rest-apis | jq '.[][]| .id' | tr -d '""')
# App settings
APP=hello-sam
BUCKET=$(APP)-$(tss)
URL=https://$(id).execute-api.ap-southeast-2.amazonaws.com/Stage/hello

debug:
	echo $(id)

bucket:
	echo $(BUCKET)

deps:
	go get -u ./...

clean:
	rm -rf ./hello-world/hello-world

validate:
	sam validate --template template.yaml

build: validate
	# TODO Fix naming?
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world

# Run the golang test's
test:
	go test ./...

# Run a local smoke test
local_test:
	echo "Make sure you run 'make run' in a seperate terminal first"
	curl http://127.0.0.1:3000/hello

# Run a smoke test aginest the deployed
smoke_test: deploy
	curl $(URL)

# Run the lambda locally via a docker container thats build into the sam cli toolbox.
run: build
	sam local start-api

bootstrap:
	echo $(ts) > .tmp
	aws s3 mb s3://$(BUCKET)


destroy: package
	aws s3 rb s3://$(BUCKET)
	# TODO Destroy the lambda and API gateway setup too

# Setup local dev
setup:
	pip3 install awscli
	pip install --user aws-sam-cli

package: build
	sam package \
		--template-file template.yaml \
		--output-template-file package.yaml \
		--s3-bucket "$(BUCKET)" \

# deploy to aws
deploy: package
	sam deploy \
		--template-file package.yaml \
		--stack-name $(APP) \
		--capabilities CAPABILITY_IAM

