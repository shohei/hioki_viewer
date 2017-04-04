# HIOKI viewer
![](image/temp.png)

# Target device
[LR8402](https://www.hioki.co.jp/shop/products/detail/122)

![LR8402](https://www.hioki.co.jp/shop/html/upload/save_image/lr8402.jpg)

# How to use
```sh
$ go get github.com/shohei/hioki_viewer
$ cd $GOPATH/github.com/shohei/hioki_viewer
$ go build server.go
$ ./server --logger-ip=192.168.100.206 # IP address of HIOKI logger
```
## Install on Raspberry Pi
```sh
$ sudo mv start_server.sh /etc/init.d
$ sudo chmod a+x /etc/init.d/start_server.sh
$ sudo update-rc.d start_server.sh defaults
$ sudo reboot #-> Here we go
```