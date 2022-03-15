FROM golang:1.17.8-buster

RUN apt-get update &&\
  apt install sudo
WORKDIR /app

COPY . /app/