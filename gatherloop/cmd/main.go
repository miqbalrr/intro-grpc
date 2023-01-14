package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	gl "gatherloop/pb"

	"google.golang.org/grpc"
)

const (
	ServerAddress = "0.0.0.0:50052"
)

type SeminarData struct {
	sync.Mutex
	Seminar map[string]*gl.Seminar
}

var seminar = &SeminarData{
	Seminar: make(map[string]*gl.Seminar),
}

type GatherloopServer struct{}

func (s *GatherloopServer) BuatSeminar(ctx context.Context,
	in *gl.BuatSeminarRequest) (*gl.ServiceResponse, error) {

	sm := gl.Seminar{
		Pembicara: in.Pembicara,
		Acara:     in.Acara,
		Waktu:     in.Waktu,
		Lokasi:    in.Lokasi,
		Berbayar:  in.Berbayar,
	}

	if in.Waktu == "" || in.Lokasi == "" {
		return nil, errors.New("waktu dan lokasi tidak boleh kosong")
	}

	key := fmt.Sprintf("%s-%s", in.Waktu, in.Lokasi)

	seminar.Mutex.Lock()
	seminar.Seminar[key] = &sm
	seminar.Mutex.Unlock()

	return &gl.ServiceResponse{
		Status: "OK",
	}, nil
}

func (s *GatherloopServer) GetSeminar(ctx context.Context,
	in *gl.GetSeminarRequest) (*gl.GetSeminarResponse, error) {

	res := &gl.GetSeminarResponse{
		Seminars: make([]*gl.Seminar, 0),
	}

	for _, s := range seminar.Seminar {
		res.Seminars = append(res.Seminars, s)
	}

	return res, nil
}

func main() {
	srv := grpc.NewServer()
	var glServer *GatherloopServer

	gl.RegisterGatherloopServiceServer(srv, glServer)

	log.Println("Starting GatherLoop server at", ServerAddress)

	l, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", ServerAddress, err)
	}

	log.Fatal(srv.Serve(l))
}
