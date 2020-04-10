#!/bin/sh

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculator.proto --go_out=plugins=grpc:.