# GoAPIService

## Foreword

This is not meant to be run. This is a proof of concept library and there
are dependencies outside of this repository that are needed to make the code
operational. Specifically Postgresql with a pre-specified table schema. There
are however instructions on setting this up in the data portion of this project
but that is not the goal here and these notes may not be maintained.

## Goal

The purpose of this repository is to be a playground for Go Lang Web API design.
Best practices may not always be adhered to and the overall structure of this
project is likely to change frequently. I will be experimenting with different
tools and libraries. This package is not meant to be exported or used in other
projects. If additional new control flow designs are implimented the exisiting
project will become nested and renamed to reflect its updated purpose.

## External Libraries

As it stands the existing project is built using Gin. There are a few libraries
that are used but [Gin](https://github.com/gin-gonic/gin) and
[lib/pq](https://github.com/lib/pq) are the only ones that are not part of the
Go standard libraries.

## Resources

* [Go SQL API](https://pkg.go.dev/database/sql)
* [Go Gin](https://go.dev/doc/tutorial/web-service-gin)
* [Nic Jackson MicroServices With Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)

## TODO

[ ] - Merge HTTP (Standard Library Only) Web server into this repository for
completeness.
