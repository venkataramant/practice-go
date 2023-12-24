package main

import (
	"fmt"
	"log"
	"os"

	"github.com/venkataramant/grpc/pb/backoffice"
	"github.com/venkataramant/grpc/pb/rides"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	displayRides()
	displayInvoice()
}
func displayInvoice() {
	lineItem1 := &backoffice.LineItem{
		SKU:    "mysku123",
		Amount: 65,
		Price:  30,
	}
	lineItem2 := &backoffice.LineItem{
		SKU:    "mysku14",
		Amount: 100,
		Price:  10,
	}
	inv := &backoffice.Invoice{
		ID:       "myid123",
		Time:     timestamppb.Now(),
		Customer: "myCustomer",
		Items:    []*backoffice.LineItem{lineItem1, lineItem2},
	}
	fmt.Printf("%+v", inv)
	print("\ninvoice::\n")
	fmt.Print(inv)
	print("\ninvoice2::\n")
	fmt.Print(&inv)

	print("\ninvoice3::\n")
	fmt.Print(*inv)
}
func displayRides() {

	myRequest := rides.StartRequest{
		Id:       "abcdefghijklmnopqrt",
		DriverId: "007",
		Location: &rides.Location{
			Lat: 51.4871871,
			Lng: -0.1266743,
		},
		PassengerIds: []string{"A", "B"},
		Time:         timestamppb.Now(),
		Type:         rides.RideType_POOL,
	}

	fmt.Printf("%+v", myRequest)
	println("output::\n")
	fmt.Println(&myRequest)

	data, err := proto.Marshal(&myRequest)
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		fmt.Println(data)
	}

	var clientSideRequest rides.StartRequest
	if err2 := proto.Unmarshal(data, &clientSideRequest); err2 != nil {
		log.Fatalf("Unmarshall error %s", err2)
	}
	print("Print clientSideRequest:: \n")
	fmt.Println(&clientSideRequest)
	fmt.Println("Size of Proto", len(data))
	jData, err3 := protojson.Marshal(&myRequest)
	if err3 != nil {
		log.Fatalf("Unnable to convert to json %s", err3)
	}
	fmt.Println(string(jData))
	os.Stdout.Write(jData)
}
