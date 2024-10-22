package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"

	pb "handin2/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)


type Hospital struct {
	pb.UnimplementedSendAggregatedShareServiceServer
	AggregatedShareList []int64
}

func (h *Hospital) SendAggregatedShare(ctx context.Context, msg *pb.AggregatedShare) (*pb.Acknowledgement, error) {
	log.Printf("Received: %v", msg.AggregatedShareOfSecret)
	h.AggregatedShareList = append(h.AggregatedShareList, msg.AggregatedShareOfSecret)
	if len(h.AggregatedShareList) == 3 {
		log.Printf("Sum of aggregated shares: %v", h.sumAggregatedShares())
	}
	return &pb.Acknowledgement{Ack: true}, nil
}

func (h *Hospital) sumAggregatedShares() int64 {
	var sum int64
	for _, share := range h.AggregatedShareList {
		sum += share
	}
	return sum
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	log.Println("Loading server certificates...")

	// Load the certificate of the CA who signed the client's certificate
	pemClientCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		log.Fatalf("Failed to append CA certificate to pool")
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}
	log.Println("Loaded and appended CA certificate.")

	// Load the server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if  err != nil {
		log.Fatalf("Failed to load server certificate or key: %v", err)
		return nil, err
	}
	log.Println("Loaded server certificate and key.")

	// Create the credentials and return it
    config := &tls.Config{
        Certificates: []tls.Certificate{serverCert},
        ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
    }

	log.Println("Server TLS credentials configured successfully.")
    return credentials.NewTLS(config), nil
}

func main () {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tlsCreds, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCreds),
	)
	pb.RegisterSendAggregatedShareServiceServer(grpcServer, &Hospital{})


	fmt.Println("Hospital Server is running on port 8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}	
}

