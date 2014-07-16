require 'rubygems'
require 'bundler/setup'
require 'protobuf'
require 'socket'
require 'csv'
require './ProtoTest.pb'

csv = <<CSV
itemid,itemname,itemvalue,itemType
1,FirstItemName,222,0
2,SecondItemName,333,1
3,ThridItemName,444,2
4,FouthItemName,555,0
5,FifthItemName,666,3
CSV

message = ProtobufTest::TestMessage.new

message.clientName = 'Ruby Client'
message.clientId = 2
message.description = 'Ruby Client'

#message.messageItems << ProtobufTest::TestMessage::MsgItem.new({id: 1, itemName: 'Test1', itemValue: 2, ItemType: ProtobufTest::TestMessage::ItemType::TypeX})

puts "Parsing CSV content"
CSV.parse(csv, {headers: true, header_converters: :symbol}) do |row|
  attrs = {id: row[:itemid].to_i, itemName: row[:itemname], itemValue: row[:itemvalue].to_i, ItemType: row[:itemType].to_i}
  message.messageItems << ProtobufTest::TestMessage::MsgItem.new(attrs)
end

begin
  # open socket
  sock = TCPSocket.new('localhost', 2110)

  puts "Sending message to socket :2110"
  sock.write message.encode
  sock.close
rescue Errno::ECONNREFUSED
  raise RuntimeError.new "Oh! seems there is not socket listening"
end
