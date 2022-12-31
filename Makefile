unit-tests:
	go test -v ./...

lint:
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run \
		-it \
		--rm \
		-v ~/go:/root/go \
		-v golangci-lint-cache:/root/.cache \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s $(OUTPUT_OPTIONS) -v
		#-c ./.golangci.yml