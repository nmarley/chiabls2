# Go bindings for Chia BLS signatures library

This is a CGo wrapper around the Chia BLS Signatures implementation in C++.

## Pre-Requisites

You will need to install a modern version of Go (1.12+), which is outside the scope of this README.

This wraps the [Chia BLS signatures](https://github.com/Chia-Network/bls-signatures) library, so you will need to install that in a place that the CGo compiler can find.

## Install

It is necessary to first build and install the Chia BLS library per the instructions in that repo.

## Testing

A Makefile is included for simplicity. By default, it runs conventional Go
formatting / linter tools and tests.

```sh
make
```

You may need to install the golint and goimports tools if not already
installed:

```sh
go get golang.org/x/tools/cmd/goimports
go get -u golang.org/x/lint/golint
```

## Usage

Please see the [example Go program to demonstrate usage of these Go bindings](https://github.com/nmarley/go-bls-signatures-example).

## Contributing

Feel free to dive in! [Open an issue](https://github.com/nmarley/chiabls2/issues/new) or submit PRs.

## License

[ISC](LICENSE)
