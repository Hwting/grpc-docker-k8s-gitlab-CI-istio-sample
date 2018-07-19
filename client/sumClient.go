package main

import (
	"log"
	"time"
	//"os"
	"fmt"
	pb "github.com/jacenr/gRPCinDockerDemo/grpcDemo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func handHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mm, ok := r.Form["m"]
	if !ok {
		fmt.Fprintf(w, "no such item: m\n")
		return
	}
	m0, err := strconv.Atoi(mm[0])
	if err != nil {
		fmt.Fprintf(w, "invlide m\n")
		return
	}
	m := int32(m0)
	nm, ok := r.Form["n"]
	if !ok {
		fmt.Fprintf(w, "no such item: n\n")
		return
	}
	n0, err := strconv.Atoi(nm[0])
	if err != nil {
		fmt.Fprintf(w, "invlide n\n")
		return
	}
	n := int32(n0)
	s := handReq(m, n)
	fmt.Fprintf(w, "s: %v\n", s)
}

func handReq(m int32, n int32) int32 {
	// connect to server
	addr := "svc-server:80"
	//if len(os.Args) > 1 {
	//        addr = os.Args[1]
	//}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSumClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//r, err := c.GetSum(ctx, &pb.Req{M:3,N:2})
	r, err := c.GetSum(ctx, &pb.Req{M: m, N: n})
	if err != nil {
		log.Fatalf("could not get resp: %v", err)
	}

	//log.Printf("the sum: %v", r.S)
	return r.S
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handHTTP))
	http.ListenAndServe(":8000", mux)
}
