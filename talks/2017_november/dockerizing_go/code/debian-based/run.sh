#!/bin/bash

docker build -t vienna-go .

docker run -p 8080:8080 -it --rm --name vienna-go vienna-go
