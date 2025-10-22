MAIN_FILE := ./cmd/main.go

.PHONY: templ-gen
templ-gen:
	go tool templ generate

.PHONY: run
run: templ-gen
	go run ${MAIN_FILE}