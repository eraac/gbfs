COVER_PROFILE="c.out"
COVER_HTML="coverage.html"

run:
	@go run example/*.go

test:
	@go test -cover -v

test-html:
	@go test -cover -v -coverprofile=${COVER_PROFILE}
	@go tool cover -html=${COVER_PROFILE} -o ${COVER_HTML}
	@open ${COVER_HTML}

clean:
	@rm ${COVER_PROFILE} || true
	@rm ${COVER_HTML} || true

code-quality:
	@echo "== GOLINT =="
	@find . -type d | xargs -L 1 golint
	@echo "== GO VET =="
	@go tool vet .
