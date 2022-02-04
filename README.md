# gogetter

Tool implement with Go for finding broken links on web page. The search starts from a given URL and recursively moves on to other URLs by following the links on the page. To avoid things getting out of hand, the recursive procedure will only go through pages which are under the same domain as the starting URL.

I have also implemented similar tool with Python, which can be found [here](https://github.com/Kaltsoon/dead-link-checker).

## Requirements

Go version 1.17 or higher.

## How to use

1. Compile the code by running `go build`

2. Execute the executable by running `./gogetter -url https://github.com/Kaltsoon/gogetter -maxdepth 1`. You can see the available options by running `./gogetter -h`.