#!/bin/bash

export GOPATH=$(go env GOPATH)
# proto-gen- under $GOPATH/bin  ($GOBIN)
export PATH=$PATH:$GOPATH/bin
