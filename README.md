# Fasttrack Code Test Solution

Golang implementation for Fasttrack super simple quiz.

### Implementation Stack

* Language: Golang
* Database: Just in-memory
* Communication protocol: gRPC
* Client: CLI using [promptui](https://github.com/manifoldco/promptui)


### Installation (MacOS)

Protobuf (for development only):

    $ brew install protobuf

protoc-gen libraries (for development only):

    $ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

Clone repository:

    $ git clone git@github.com:gersakbogdan/fasttrackquiz.git


### Run (docker container)

[Docker Compose](https://docs.docker.com/compose/) is used to make the testing as smooth as possible.

    // build and run quiz server inside docker container
    $ make build

    // run quiz client
    $ make run

### Run (without docker)

    // create executable
    $ go install ./cmd/quizer

    // run server & client
    $ quizer server
    $ quizer client

### Result

![Golang quiz example](./screenshots/quizer.gif)

### ToDo

* Document code using GoDoc standards (https://blog.golang.org/godoc-documenting-go-code)
* Add unit & integration tests
* Create Rest API client (integrate grpc-gateway)
* Service monitoring
