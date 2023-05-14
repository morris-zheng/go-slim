default:
	@echo 'Usage of make: [ serve | build | save | deploy ] image_name=example '

image_name = none

version = $(shell date +%Y-%m-%d_%H:%M:%S)

serve:
	go run cmd/main.go

image:
	docker build -f build/Dockerfile -t ${image_name}:latest .

save:
	docker save -o ${image_name}-${version}.img ${image_name}:latest

deploy: image save