#!/bin/bash

protoc greet/greetpb/greet.protoc --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.protoc --go_out=plugins=grpc:.