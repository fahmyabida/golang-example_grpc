#!/bin/bash

protoc greet/greetpb/greet.protoc --go_out=plugins=grpc:.