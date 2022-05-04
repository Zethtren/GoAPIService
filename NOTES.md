# Web API in Go Lang

This was an implimentation that was derived from examples on go.dev website.

As well as from videos by Nic Jackson on Youtube in his Building Microservices with Go series.

Currently I have only made it through the first three episodes of this series.

In order for this API to work you will need to have a postgresql database running on localhost port 5432 with username password and database all set to "postgres".


Their is a commented out section in /data/albums.go that will allow you to set environment variables instead and use any passwords of your choosing.

The other requirement is that there is an albums table with id as SERIAL, title VARCHAR, artists VARCHAR, and price DOUBLE PRECISION.

You can have additional values but they will need to be nullable as this script is not yet designed to handle non-null values.

