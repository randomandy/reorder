# reorder
Efficiently ordering booking locations

# Index
- [Deployment & Installation](https://github.com/randomandy/reorder#deployment--installation)

# Deployment & Installation

## Go

You will need a working Go installation.
Download Go and follow the installation instructions here:

https://golang.org/doc/install

Once completed, set up your Go workspace where source code and binaries are going to be located. For example on any Linux, BSD, UNIX, etc:

	mkdir ~/golang

This will create a directory called `golang` in your current user's home directory.

Finally you need to set your Go workspace via $GOPATH environment variable.
Easiest would be to add this to your .bashrc or .bash_profile:

	export GOPATH=$HOME/golang

## Checkout

If you don't have any other Go projects yet, create a src structure for the git checkout, followed by cloning the repo:

	mkdir -p $GOPATH/src/github.com/randomandy
	cd $GOPATH/src/github.com/randomandy
	git clone git@github.com:randomandy/reorder.git

(Git will create a new directory called `reorder` in this structure)

## Compile & Run

For this project I'm not using any third party dependencies, so you are all set. To compile and run the project directly:

	go run main.go

You can also just compile it, and use the binary in ($GOPATH/bin/) afterwards:

	go install


Since this project requires a JSON file with bookings for parsing, the file can be passed as argument when running the project. If no argument is passed, it will automatically look for a file called bookingordering.json in the current directory. (For the example below I have copied my JSON file to /tmp)

	$GOPATH/bin/reorder --json /tmp/bookingordering.json

Or to run it directly from source:

	cd $GOPATH/src/github.com/randomandy/reorder
	go run main.go --json /tmp/bookingordering.json

There is also a usage help in case you get lost:

	$GOPATH/bin/reorder --help







