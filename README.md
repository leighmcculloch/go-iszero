# iszero

[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/go-iszero)](https://goreportcard.com/report/github.com/leighmcculloch/go-iszero)
[![Codecov](https://img.shields.io/codecov/c/github/leighmcculloch/go-iszero.svg)](https://codecov.io/gh/leighmcculloch/go-iszero)
[![Build Status](https://img.shields.io/travis/leighmcculloch/go-iszero.svg)](https://travis-ci.org/leighmcculloch/go-iszero)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/leighmcculloch/go-iszero)

A go package for checking if a value is zero.

I put this package together from code within the Go source, because I needed to be able to identify if a field in a struct was empty/its-zero-value, so that when encoding a struct I could determine if the field was empty and could be ommitted.

## Usage

See the example in [Go Docs](https://godoc.org/github.com/leighmcculloch/go-iszero).
