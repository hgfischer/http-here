# http-here

[![Build Status](https://travis-ci.org/hgfischer/http-here.svg?branch=master)](https://travis-ci.org/hgfischer/http-here)

`http-here` is a simple http server for testing, local development, and learning.

## Installing

Install the `go` compiler from your preferred source. I recommend the official packages from http://golang.org.

Once you have `go` installed, setup the environment variable `$GOPATH` to somewhere inside your `$HOME`. For example:

```
export $GOPATH="/home/user/go"
mkdir -p $GOPATH
```

Then install `http-here` with the following command:

```
go get github.com/hgfischer/http-here
```

This will download, build, and install `http-here` to `$GOPATH/bin/http-here`.

For easy access, you can copy the `http-here` binary to `/usr/local/bin`, or somewhere else inside `$PATH`, or you
can add `$GOPATH/bin` to your `$PATH` environment var.

## Usage

Currently there are a few options that you can check bellow or when running `http-here -h`:

```
Usage: ./http-here [flags]

Flags:
  -a=":8080": Address to start listening for HTTP connections
  -c=-1: Set cache time, in seconds, for cache-control max-age header
  -cors=false: Enable CORS via the 'Access-Control-Allow-Origin' header
```

# License 

Copyright (c) 2014, Herbert G. Fischer
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the organization nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL HERBERT G. FISCHER BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

