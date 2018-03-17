#!/bin/bash

go test -tags=integration -v -cover
go test -v -cover ./...
