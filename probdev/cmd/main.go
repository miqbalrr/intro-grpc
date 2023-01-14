package main

import (
	"context"
	"fmt"
	"log"

	gl "probdev/pb"

	"google.golang.org/grpc"
)

const (
	GatherLoopServerAddress = "0.0.0.0:50052"
)

func main() {
	conn, err := grpc.Dial(GatherLoopServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s with err %+v", GatherLoopServerAddress, err)
	}

	glClient := gl.NewGatherloopServiceClient(conn)

	res, err := glClient.BuatSeminar(context.TODO(), &gl.BuatSeminarRequest{
		Pembicara: "iqbal",
		Acara:     "Makan makan",
		Waktu:     "Lusa",
		Lokasi:    "Rumah nya Aka",
		Berbayar:  false,
	})
	if err != nil {
		log.Fatalf("Buat Seminar Error = %+v ", err)
	}

	fmt.Println(res)

	seminars, _ := glClient.GetSeminar(context.TODO(), &gl.GetSeminarRequest{})

	fmt.Println(seminars)
}
