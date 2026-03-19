#! /bin/bash

docker build -t pascal-sort .

docker run --name pascal-sort --rm pascal-sort
