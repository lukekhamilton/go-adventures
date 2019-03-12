FILE=swagger.yml

generate: validate generate_client generate_server deps

validate:
	swagger validate ${FILE}

deps:
	go get -u ./...

generate_client:
	swagger generate client -f ${FILE} # -o logs/swagger_pet_store_client.log

generate_server:
	swagger generate server -f ${FILE} # -o logs/swagger_pet_store_server.log

generate_model:
	swagger generate model --spec=${FILE}

clean:
	rm -rf cmd client models restapi logs/*

doco:
	redoc-cli serve ${FILE}

generate_doco:
	spectacle ${FILE}

server_doco:
	http-server public/lean:
