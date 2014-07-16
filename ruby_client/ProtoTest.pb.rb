##
# This file is auto-generated. DO NOT EDIT!
#
require 'protobuf/message'

module ProtobufTest

  ##
  # Message Classes
  #
  class TestMessage < ::Protobuf::Message
    class ItemType < ::Protobuf::Enum
      define :TypeX, 0
      define :TypeY, 1
      define :TypeZ, 2
    end

    class MsgItem < ::Protobuf::Message; end

  end



  ##
  # Message Fields
  #
  class TestMessage
    class MsgItem
      required :int32, :id, 1
      optional :string, :itemName, 2
      optional :int32, :itemValue, 3
      optional ::ProtobufTest::TestMessage::ItemType, :itemType, 4
    end

    required :string, :clientName, 1
    required :int32, :clientId, 2
    optional :string, :description, 3, :default => "NONE"
    repeated ::ProtobufTest::TestMessage::MsgItem, :messageItems, 4
  end

end

