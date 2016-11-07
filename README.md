# Hello with Go from a container

This repo shows one of the many use cases for Docker.
We're going to compile and hello world app written in
Go, without installing Go in our machines.

## How to run this

Clone this repository to your local machine:

```
git clone https://github.com/joaofnfernandes/hello-go.git
```

Then [install Docker for your OS](https://www.docker.com/products/docker#/mac).

If you have Mac or Linux, you can use the Makefile
in this repo to build and run our Go hello world:

```
cd hello-go
make
```

If you have Windows, or want to manually build and
run the containers:

```
cd hello-go
docker build --tag hello-go .
docker run --rm hello-go
```

## Pass environment variables

Our hello-world prints `hello Jane Doe` by default, but it allows you
to customize the name that is printed.
If the `NAME` environment is set, that name will be printed instead of "Jane Doe".

When running your app in a container you can pass environment variables to the container. You can do it like this:

```
docker run --rm --env "NAME=Dave Lauper" hello-go
```

