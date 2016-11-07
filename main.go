package main

import (
	"fmt"
	"os"
)

const (
	defaultName     = "Jane Doe"
	messageTemplate = `
Hello %s

If you're reading this message, this means that you were able to

1. Create a new Docker image

    Your images is based on the Golang image, and copies this source code
    file to a directory on the image. Check the Dockerfile to see how it's
    done.

2. Launch a new container 

    An image is basically a filesystem and some configurations to
    specify what executables and parameters to use when trying to
    run something on that filesystem.

    Images are immutable. Everything you do inside a container is
    not persisted after the container stops running.

    This is cool because you can use the same image to launch a bunch
    of containers, that do their own thing and create or delete files.
    If you want to persist the changes done in the container, you can
    create a new image based on the container.

3. Compile and run this source code

	Your container compiled this Go source code file, and then ran
	it to produce this output.

	This is pretty cool because you were able to do a bunch of things
	without having to install the Go compiler or runtime in your
	machine.
`
)

func main() {
	name := os.Getenv("NAME")
	if name == "" {
		name = defaultName
	}
	fmt.Printf(messageTemplate, name)
}
