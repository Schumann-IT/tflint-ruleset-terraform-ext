default: build

test:
	go test ./...

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-schumann-it-terraform-ext ~/.tflint.d/plugins/tflint-ruleset-terraform-ext
