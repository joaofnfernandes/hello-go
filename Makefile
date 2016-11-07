.PHONY: build run

default: run

# Runs a new container with the image with just built
# When the app finishes, remove the container
run: build
	docker run --rm hello-go

# Builds the Docker image using the Dockerfile
# in this directory, and assigns it the name 'hello-go'
build:
	@docker build --tag hello-go .
