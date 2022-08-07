package main

import (
	"encoding/json"
	"fmt"
	"log"

	pbuser "proto-buffs-basic/pbout/protos"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	// user "pbout/protos/users.pb.go"
)

func main() {
	fmt.Println("Go - Protobufs Basic")

	user := &pbuser.Users{
		Fullname:  "Abhishek Patil",
		Age:       24,
		Status:    pbuser.Status_ACTIVE,
		CreatedAt: timestamppb.Now(),
	}

	fmt.Println(user)
	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(user.Status)
	// printing out our raw protobuf object
	fmt.Println(data)

	// Google marshalling, problems, emun as int, timestamp as another strcut etc.
	jsonBytes, _ := json.MarshalIndent(user, "", "	")
	fmt.Println(string(jsonBytes))

	// enums as string but default values and nulls are ommitting
	jsonBytes, _ = protojson.Marshal(user)
	fmt.Println(string(jsonBytes))

	// marshalling with all fields
	m := protojson.MarshalOptions{
		Indent: "	",
		EmitUnpopulated: true,
	}

	jsonBytes, _ = m.Marshal(user)
	fmt.Println(string(jsonBytes))

	m = protojson.MarshalOptions{
		Indent: "	",
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}

	jsonBytes, _ = m.Marshal(user)
	fmt.Println(string(jsonBytes))

	um := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	//Unmarshal
	var nuser pbuser.Users

	um.Unmarshal(jsonBytes, &nuser)
	fmt.Println(nuser)
}
