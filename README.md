# Go crawler 

[WIP]

## Usage

@todo


## Exemples

### Simple exemple

@todo

### Complete exemple

@todo


## Installation from binaries

@todo


## Installation from sources

### Install Go

```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source ~/.gvm/scripts/gvm
sudo apt-get install bison
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.11 --prefer-binary
gvm use go1.11 --default
```

### Install packages

```
make deps
```

### Build

```
make build
```

### Cross compile

Use gox :

```
go get github.com/mitchellh/gox
gox
```
