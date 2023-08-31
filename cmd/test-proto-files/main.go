package main

import (
	"fmt"
	"github.com/MatejaMaric/testing-grpc/pkg/pb"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ProtoType string

const (
	ProtoWire = "proto"
	ProtoText = "prototext"
	ProtoJSON = "protojson"
)

func main() {
	log.Println("Started Protobuf test...")

	obj := &pb.Person{
		Name:        "Test",
		DateOfBirth: 1990,
	}

	if err := ProtobufToFile(ProtoWire, "protowire_test_file", obj); err != nil {
		log.Fatalln(err)
	}
	if err := ProtobufToFile(ProtoText, "prototext_test_file", obj); err != nil {
		log.Fatalln(err)
	}
	if err := ProtobufToFile(ProtoJSON, "protojson_test_file", obj); err != nil {
		log.Fatalln(err)
	}
}

func ProtobufToFile(pt ProtoType, filename string, msg protoreflect.ProtoMessage) error {
	var proto_output []byte
	var err error

	switch pt {
	case ProtoWire:
		proto_output, err = proto.Marshal(msg)
	case ProtoText:
		proto_output, err = prototext.Marshal(msg)
	case ProtoJSON:
		proto_output, err = protojson.Marshal(msg)
	}
	if err != nil {
		return fmt.Errorf("failed marshaling the object with %s: %w", pt, err)
	}

	if err := os.WriteFile(filename, proto_output, 0644); err != nil {
		return fmt.Errorf("failed writing %s output to file: %w", pt, err)
	}

	return nil
}
