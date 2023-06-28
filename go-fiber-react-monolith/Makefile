run: build-ui
	go run .

watch:
	gowatch

build-ui:
	npm --prefix ui install
	npm --prefix ui run build 

.PHONY: run build-ui watch
.SILENT: run build-ui watch