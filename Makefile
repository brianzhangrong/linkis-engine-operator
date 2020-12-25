IMAGE_NAME = harbor.ihomefnt.com/ai/engine-ops

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o engine-ops . 
	docker build -t $(IMAGE_NAME) . 
	docker push  $(IMAGE_NAME) 

# .PHONY: test
# test:
# 	docker build -t $(IMAGE_NAME)-candidate .
# 	IMAGE_NAME=$(IMAGE_NAME)-candidate test/run
