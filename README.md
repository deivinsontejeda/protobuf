Google Protocol Buffer
========

A very simple example using [Protobuf protocol](https://developers.google.com/protocol-buffers/docs/proto).

This example is composed by two sides a client which send serialized data using Protobuf protocol and server which receive it and decode to processing. For this example just send the input to standard output, but you can think instead of send to standard output could send to store a database.

Once you cloned this repositorie need get up the server (written in Go) which will decode messages sent by clients. This server listen on port `2110`.

Install Protobuf & dependencies
========

Install Protobuf protocol compiler
```bash
brew install protobuf
```
More details about how to install Protobuf or further issue visit its [oficial page](https://developers.google.com/protocol-buffers/)


Install Go library
```bash
$ go get -u code.google.com/p/goprotobuf/...
```

Install [protobuf](https://github.com/localshred/protobuf) gem to support compile `.protoc` files for Ruby.

```bash
$ gem install protobuf
```

How to run?
======

In a consola go to `server` folder and run follow command:

```bash
$ go run proto_serve.go
```

In another console run the clients which read a CSV file and send its content to server.

```bash
$ go run proto_client.go
```

You must see how client send messages and severs how decode them :)

Now, you can repeat the previous step but now running ruby client.

```bash
$ ruby client.rb
```

References
=======

* [5 Reasons to Use Protocol Buffers Instead of JSON For Your Next Service](http://blog.codeclimate.com/blog/2014/06/05/choose-protocol-buffers/)

Happy Hacking!
