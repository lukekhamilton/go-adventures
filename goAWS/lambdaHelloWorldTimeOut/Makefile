FUNCTION_NAME="luke-goAWS-HelloWorldTimeOut"

all: clean build package deploy

clean:
	rm ./lambda_handler
	rm ./handler.zip

build:
	GOOS=linux go build lambda_handler.go

package:
	zip handler.zip ./lambda_handler

deploy:
	aws lambda update-function-code --function-name ${FUNCTION_NAME} --zip-file fileb://handler.zip


create: build package
	aws lambda create-function \
  		--region ap-southeast-2 \
  		--function-name ${FUNCTION_NAME} \
  		--memory 128 \
  		--role arn:aws:iam::875311970826:role/service-role/guru-lab1-go-helloWorld-role \
  		--runtime go1.x \
  		--zip-file fileb://handler.zip \
  		--handler lambda_handler
