#@echo off
set GOPROXY=https://proxy.golang.org,direct
set http_proxy=http://127.0.0.1:7890
set https_proxy=%http_proxy%
cmd \k
