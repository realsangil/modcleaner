<h1 align="center">Welcome to modcleaner 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-v0.0.1-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/realsangil/modcleaner/blob/master/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

Modcleaner is a tool to erase the go module.  
The objects to be erased are as follows.
- $GOPATH/pkg/mod/github.com/realsangil/modcleaner@v0.0.1
- $GOPATH/pkg/mod/cache/download/github.com/realsangil/modcleaner/@v/v0.0.1.info
- $GOPATH/pkg/mod/cache/download/github.com/realsangil/modcleaner/@v/v0.0.1.lock
- $GOPATH/pkg/mod/cache/download/github.com/realsangil/modcleaner/@v/v0.0.1.mod
- $GOPATH/pkg/mod/cache/download/github.com/realsangil/modcleaner/@v/v0.0.1.zip
- $GOPATH/pkg/mod/cache/download/github.com/realsangil/modcleaner/@v/v0.0.1.ziphash
- $GOPATH/pkg/mod/cache/vcs/d5eb110f070eed22f7a69bc52fec68e273092374
- $GOPATH/pkg/mod/cache/vcs/d5eb110f070eed22f7a69bc52fec68e273092374.lock
- $GOPATH/pkg/mod/cache/vcs/d5eb110f070eed22f7a69bc52fec68e273092374.info



## Install

```sh
go get -u github.com/realsangil/modcleaner
```

## Usage

```sh
modcleaner -v pkg0 pkg1 pkg2... 

# root permission
sudo GOPATH=/path/to/gopath modcleaner -v pkg0 pkg1 pkg2... 
```

## Author

👤 **realsangil**

* Website: https://blog.realsangil.net
* Github: [@realsangil](https://github.com/realsangil)

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [realsangil](https://github.com/realsangil).<br />
This project is [MIT](https://github.com/realsangil/modcleaner/blob/master/LICENSE) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_