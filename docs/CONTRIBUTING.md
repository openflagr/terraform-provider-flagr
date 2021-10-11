# Contributing

### Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 1.x
-	[Go](https://golang.org/doc/install) 1.17 (to build the provider plugin)

### Building The Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.17+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

Clone repository to: `$GOPATH/src/github.com/marceloboeira/terraform-provider-flagr`

```sh
mkdir -p $GOPATH/src/github.com/marceloboeira; cd $GOPATH/src/github.com/marceloboeira
git clone git@github.com:marceloboeira/terraform-provider-flagr
```

Enter the provider directory and build the provider

```sh
cd $GOPATH/src/github.com/marceloboeira/terraform-provider-flagr

make build
```

To test it locally you need to install the provider with:

```sh
make install
```

### Makefile

Makefile is your friend:

```
make build     Builds the local architecture binary to the root folder
make compose   Starts test dependencies with docker-compose
make format    Formats go and terraform code
make help      Lists the available commands
make install   Builds the local architecture binary and install it on the local terraform cache
make release   Builds release-binaries for all architectures
make test      Runs tests
make testacc   Runs acceptance tests
```

## Test suite

In order to test the provider, you can simply run `make test`.

```sh
make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

**Note**: you need to create the `.env` file following the `.env.example` conventions. The Flagr API Host/Path must be set so that acceptance tests run properly.

```sh
make testacc
```
