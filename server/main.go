package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	tdxClient "github.com/google/go-tdx-guest/client"
	tdxProto "github.com/google/go-tdx-guest/proto/tdx"

	pb "remote-attest-test-server-go/proto/attest"
)

type server struct {
	pb.UnimplementedAttestServer
}

func (s *server) Attest(ctx context.Context, req *pb.AttestRequest) (*pb.AttestReply, error) {
	log.Printf("Received Attest request for report_data=%x", req.GetReportData())
	quoteProvider, err := tdxClient.GetQuoteProvider()
	if err != nil {
		log.Fatalf("Failed to get quote provider: %v", err)
	}
	var reportData [64]byte
	copy(reportData[:], req.GetReportData())
	quote, err := tdxClient.GetQuote(quoteProvider, reportData)
	if err != nil {
		log.Fatalf("Failed to get quote: %v", err)
	}

	quoteProto := quote.(*tdxProto.QuoteV4)

	// for debug
	for i, rtmr := range quoteProto.TdQuoteBody.Rtmrs {
		log.Printf("rtmr[%d]: %x", i, rtmr)
	}

	quoteBytes, err := json.MarshalIndent(quote, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal quote to JSON: %v", err)
	}

	quoteStr := string(quoteBytes)
	// log.Println("Full Quote JSON:", quoteStr)

	return &pb.AttestReply{
		Quote: quoteStr,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server listening on port 50051...")

	grpcServer := grpc.NewServer()

	pb.RegisterAttestServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
