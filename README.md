# gridfs-file-upload

gridfs-file-upload project contains two API's 
1. Upload API is POST API which upload the file in mongodb with GridFS.
2. Download API is also post api user pass file name and id and download the file from data base.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.
##### What is GridFS ?
A grid file system is a computer file system whose goal is improved reliability and availability by taking advantage of many smaller file storage areas.
In MongoDB, use GridFS for storing files larger than 16 MB.

#####GridFS stores files in two collections:
chunks stores the binary chunks. For details, see The chunks Collection.
files stores the file’s metadata. For details, see The files Collection.
GridFS places the collections in a common bucket by prefixing each with the bucket name. By default, GridFS uses two collections with a bucket named fs:<br>
<b>fs.files</b> </br>
<b>fs.chunks</b></br>

### Installation go(lang)

<b>Code Assignment </b> project is <b> Go language </b> based.
<br/>Install Go with [homebrew](https://brew.sh/):

```Shell
sudo brew install go
```

with [apt](https://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

[install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)


<br/>Install packages from github to my gopath/
```Shell
go get -u github.com/gorilla/mux
```

## Usage

TODO: Write usage instructions

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b code-assignment`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin code-assignment`
5. Submit a pull request.

## go-testing
The test directory contains tests of the Go tool chain and runtime.
It includes black box tests, regression tests, and error output tests.

To run just these tests, execute:

    $ go test -run NameOfTest
    $ go test -run Test_TriangleType
    
   Some Commands for test files -
    
    go test -run ''      # Run all tests.
    go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
    go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
    go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
   

Standard library tests should be written as regular Go tests in the appropriate package.

The tool chain and runtime also have regular Go tests in their packages.
The main reasons to add a new test to this directory are:

* it is most naturally expressed using the test runner; or
* it is also applicable to `gccgo` and other Go tool chains.