package main

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/deivinsontejeda/protobuf/ProtobufTest"
	"io"
	"net"
	"os"
	"strconv"
)

var (
	filename = flag.String("f", "./data.csv", "Enter the filename to read from")
	HTTPAddr = flag.String("http", "127.0.0.1:2110", "Address to send HTTP requests")
)

type Headers []string

type client struct {
	Id          int32
	Name        string
	Description string
}

func (h Headers) getHeaderIndex(headerName string) int {
	if len(headerName) >= 2 {
		for index, s := range h {
			if s == headerName {
				return index
			}
		}
	}
	return -1
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}

func sendData(data []byte, dst *string) {
	conn, err := net.Dial("tcp", *dst)
	checkError(err)
	n, err := conn.Write(data)
	checkError(err)
	fmt.Println("Sent " + strconv.Itoa(n) + " bytes")
}

// Read data from a CSV
func readData(fname *string) ([]byte, error) {
	file, err := os.Open(*fname)
	checkError(err)
	defer file.Close()
	csvreader := csv.NewReader(file)

	var headers Headers
	headers, err = csvreader.Read()
	checkError(err)

	ITEMIDINDEX := headers.getHeaderIndex("itemid")
	ITEMNAMEINDEX := headers.getHeaderIndex("itemname")
	ITEMVALUEINDEX := headers.getHeaderIndex("itemvalue")
	ITEMTYPEINDEX := headers.getHeaderIndex("itemType")

	ProtoMessage := new(ProtobufTest.TestMessage)

	client := client{
		Id:          2,
		Name:        "GoClient",
		Description: "Go Client using Protobuf",
	}

	ProtoMessage.ClientName = proto.String(client.Name)
	ProtoMessage.ClientId = proto.Int32(client.Id)
	ProtoMessage.Description = proto.String(client.Description)

	//loop through the records
	for {
		record, err := csvreader.Read()
		if err != io.EOF {
			checkError(err)
		} else {
			break
		}
		//Populate items
		testMessageItem := new(ProtobufTest.TestMessage_MsgItem)
		itemid, err := strconv.Atoi(record[ITEMIDINDEX])
		checkError(err)
		testMessageItem.Id = proto.Int32(int32(itemid))
		testMessageItem.ItemName = &record[ITEMNAMEINDEX]
		itemvalue, err := strconv.Atoi(record[ITEMVALUEINDEX])
		checkError(err)
		testMessageItem.ItemValue = proto.Int32(int32(itemvalue))
		itemtype, err := strconv.Atoi(record[ITEMTYPEINDEX])
		checkError(err)
		iType := ProtobufTest.TestMessage_ItemType(itemtype)
		testMessageItem.ItemType = &iType

		ProtoMessage.MessageItems = append(ProtoMessage.MessageItems, testMessageItem)

		fmt.Println(record)
	}

	//fmt.Println(ProtoMessage.Messageitems)
	return proto.Marshal(ProtoMessage)
}

func main() {
	flag.Parse()

	data, err := readData(filename)
	checkError(err)
	sendData(data, HTTPAddr)
}
