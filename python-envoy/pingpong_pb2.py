# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pingpong.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='pingpong.proto',
  package='pingpong',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\x0epingpong.proto\x12\x08pingpong\"\x1b\n\x0bPingRequest\x12\x0c\n\x04ping\x18\x01 \x01(\t\"\x1c\n\x0cPongResponse\x12\x0c\n\x04pong\x18\x01 \x01(\t2L\n\x0fPingPongService\x12\x39\n\x08pingPong\x12\x15.pingpong.PingRequest\x1a\x16.pingpong.PongResponseb\x06proto3'
)




_PINGREQUEST = _descriptor.Descriptor(
  name='PingRequest',
  full_name='pingpong.PingRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ping', full_name='pingpong.PingRequest.ping', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=28,
  serialized_end=55,
)


_PONGRESPONSE = _descriptor.Descriptor(
  name='PongResponse',
  full_name='pingpong.PongResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pong', full_name='pingpong.PongResponse.pong', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=85,
)

DESCRIPTOR.message_types_by_name['PingRequest'] = _PINGREQUEST
DESCRIPTOR.message_types_by_name['PongResponse'] = _PONGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

PingRequest = _reflection.GeneratedProtocolMessageType('PingRequest', (_message.Message,), {
  'DESCRIPTOR' : _PINGREQUEST,
  '__module__' : 'pingpong_pb2'
  # @@protoc_insertion_point(class_scope:pingpong.PingRequest)
  })
_sym_db.RegisterMessage(PingRequest)

PongResponse = _reflection.GeneratedProtocolMessageType('PongResponse', (_message.Message,), {
  'DESCRIPTOR' : _PONGRESPONSE,
  '__module__' : 'pingpong_pb2'
  # @@protoc_insertion_point(class_scope:pingpong.PongResponse)
  })
_sym_db.RegisterMessage(PongResponse)



_PINGPONGSERVICE = _descriptor.ServiceDescriptor(
  name='PingPongService',
  full_name='pingpong.PingPongService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=87,
  serialized_end=163,
  methods=[
  _descriptor.MethodDescriptor(
    name='pingPong',
    full_name='pingpong.PingPongService.pingPong',
    index=0,
    containing_service=None,
    input_type=_PINGREQUEST,
    output_type=_PONGRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PINGPONGSERVICE)

DESCRIPTOR.services_by_name['PingPongService'] = _PINGPONGSERVICE

# @@protoc_insertion_point(module_scope)
