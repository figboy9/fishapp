#!/bin/sh

docker run --rm -v ${PWD}:/proto ezio1119/protoc \
-I/proto \
-I/go/src/github.com/envoyproxy/protoc-gen-validate \
--doc_out=. \
--doc_opt=markdown,README.md \
user.proto chat.proto event.proto post.proto event.proto image.proto