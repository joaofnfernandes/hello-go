# Our image is based on the golang image
FROM golang:1.7-alpine
MAINTAINER joao.fernandes@docker.com

# We add the files in this directory
# to the /go/src/app directory in the image
ADD . /go/src/app

# Change the default path to this dir
WORKDIR /go/src/app

# The image has the go binaries installed
# Use it to compile our main.go file
RUN go install

# By default, launch the executable we compiled
CMD ["app"]