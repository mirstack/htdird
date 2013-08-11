# htdird [![Build Status](https://travis-ci.org/mirstack/htdird.png?branch=master)](https://travis-ci.org/mirstack/htdird)

**Static HTTP server**

This is a tiny HTTP server that serves static files from specific directory. The server is just 
around 100 lines of code written in Golang.

## Installation

To install `htdird` go to [releases][releases] page and pick latest package. Download
it and unpack with desired prefix:

    $ unzip -d /usr/local htdird-*.zip

[releases]: https://github.com/mirstack/htdird/releases

### Installing from source

The package has no external dependencies and is easily go-installable:

    $ go install gitbub.com/mirstack/htdird

You can also download it and build manually.

    $ git clone https://github.com/mirstack/htdird.git
    $ cd htdird
    $ go build . && go install .
    
## Usage

The idea behind the project is extremally simple, just start serving given directory on addres that 
you want:

    $ htdird :8080 /path/to/directory
    
Server doesn't provide anything else, doesn't deal with authentication, nor SSL nor any other options.
And why is that so? Because this tiny bit has been made to use internally within very secured virtual
private formation with no direct access from the outside world. 

It can be also used as a lowest layer in serving static assets in environments that dont provide access
to static servers (for example on Heroku), an example `Procfile`:

    $ bin/htdird :$PORT ./public
    
Where `bin/htdird` is added to your repository.

## Hacking

If you wanna hack on `htdird` just clone the repo and play with the code. You can run the tests at 
any time with standard `go test` tool:

    $ go test .

### Releasing

To build a new release use bundled make target called `pack`:

    $ make pack

It will wrap the binary and other needed files into a zip archive and save
it to `pkg/htdird-x.y.z-{os}-{arch}.zip`.

## Contribute

1. Fork the project.
2. Write awesome code (in a feature branch).
3. Test, test, test...!
4. Commit, push and send Pull Request.
