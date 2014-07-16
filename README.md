protobuf
========

Practical guide using [Protobuf protocol](https://developers.google.com/protocol-buffers/docs/proto) 

Install Protobuf 
========

Install Protobuf protocol compiler
```bash
brew install protobuf
```

Install Go library
```bash
$ go get -u code.google.com/p/goprotobuf/...
```

Install [protobuf](https://github.com/localshred/protobuf) gem to support compile `.protoc` files in Ruby

```bash
$ gem install protobuf
```

How to run?
======

First at all get up the server which will decode messages sending to it using protobuf protocol.

```bash
$ go run proto_serve.go
```

And another console run the clients

```bash
$ go run proto_client.go
```

You must see how client send messages and severs how decode them :)

Now, you can repeat but now running ruby client.

```bash
$ ruby client.rb
```

enjoy!


