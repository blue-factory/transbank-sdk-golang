#
# SO variables
#
# DOCKER_USER
# DOCKER_PASS
#
# PRIVATE_CERT_PATH
# PUBLIC_CERT_PATH
# COMMERCE_CODE
# COMMERCE_EMAIL 
# SERVICE
# ENVIRONMENT
#

#
# Internal variables
#
VERSION=0.0.2
NAME=webpay
SVC=$(NAME)-api
BIN_PATH=$(PWD)/bin
BIN=$(BIN_PATH)/$(SVC)
REGISTRY_URL=$(DOCKER_USER)

PORT=5040

clean c:
	@echo "[clean] Cleaning bin folder..."
	@rm -rf bin/

run ra:
	@echo "[running] Running service..."
	@PORT=$(PORT) \
	PRIVATE_CERT_PATH=$(PRIVATE_CERT_PATH) \
	PUBLIC_CERT_PATH=$(PUBLIC_CERT_PATH) \
	COMMERCE_CODE=$(COMMERCE_CODE) \
	COMMERCE_EMAIL=$(COMMERCE_EMAIL) \
	SERVICE=$(SERVICE) \
	ENVIRONMENT=$(ENVIRONMENT) \
	go run cmd/http/integration/main.go

build ba: 
	@echo "[build] Building service..."
	@cd cmd/$(NAME) && go build -o $(BIN)

linux la:
	@echo "[build-linux] Building service..."
	@cd cmd/$(NAME) && GOOS=linux GOARCH=amd64 go build -o $(BIN)

docker d:
	@echo "[docker] Building image..."
	@docker build -t $(SVC):$(VERSION) .
	
docker-login dl:
	@echo "[docker] Login to docker..."
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)

push p: linux docker docker-login
	@echo "[docker] pushing $(REGISTRY_URL)/$(SVC):$(VERSION)"
	@docker tag $(SVC):$(VERSION) $(REGISTRY_URL)/$(SVC):$(VERSION)
	@docker push $(REGISTRY_URL)/$(SVC):$(VERSION)

test t:
	@echo "[test] Testing $(NAME)..."
	@PORT=$(PORT) \
	PRIVATE_CERT_PATH=$(PRIVATE_CERT_PATH) \
	PUBLIC_CERT_PATH=$(PUBLIC_CERT_PATH) \
	COMMERCE_CODE=$(COMMERCE_CODE) \
	COMMERCE_EMAIL=$(COMMERCE_EMAIL) \
	SERVICE=$(SERVICE) \
	ENVIRONMENT=$(ENVIRONMENT) \
	go test -count=1 -v . 

.PHONY: clean c run r build b linux l docker d docker-login dl push p test t