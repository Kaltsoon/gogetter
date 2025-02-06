# 🔍 gogetter

[![Test](https://github.com/Kaltsoon/gogetter/actions/workflows/test.yml/badge.svg)](https://github.com/Kaltsoon/gogetter/actions/workflows/test.yml)

Tool implemented with Go for finding broken links on a web page. The search starts from the given URL and recursively moves on to other URLs by following the links on the page. To avoid things getting out of hand, the recursive procedure will only go through pages that are under the same domain as the starting URL.

I have also implemented a similar tool with Python, which can be found [here](https://github.com/Kaltsoon/dead-link-checker).

## Requirements

Go version 1.17 or higher.

## How to use

With Docker:

1. Run the container by running `docker run --rm -v /path/to/data:/usr/src/app/data kaltsoon/gogetter --url https://github.com/Kaltsoon/gogetter --maxdepth 1`. You can see the available options by running `docker run --rm kaltsoon/gogetter -h`. The executable will produce a JSON formatted report as a file under the `/path/to/data` directory.

Without Docker:

1. Compile the code by running `go build`

2. Execute the executable by running `./gogetter --url https://github.com/Kaltsoon/gogetter --maxdepth 1`. You can see the available options by running `./gogetter -h`. The executable will produce a JSON formatted report as a file under the `data` directory.

