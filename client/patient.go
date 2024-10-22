package main

import (
    "context"
    "crypto/tls"
    "crypto/x509"
    "flag"
    "fmt"
    pb "handin2/grpc"
    "log"
    "math/rand"
    "net"
    "os"
    "sync"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
)

type Patient struct {
    pb.UnimplementedSendShareServiceServer
    patientName    string
    portAddress    string
    otherPatients  map[string]string
    ServerAddress  string
    receivedShares []int64
    privateInput   int64
}

func (p *Patient) SendShare(ctx context.Context, msg *pb.Share) (*pb.Acknowledgement, error) {
    log.Printf("Received: %v", msg.ShareOfSecret)
    p.receivedShares = append(p.receivedShares, msg.ShareOfSecret)
    if len(p.receivedShares) == 3 {
        aggregatedShare := p.sumAggregatedShares()
        p.SendAggregatedShareToHospital(context.Background(), aggregatedShare)
    }
    return &pb.Acknowledgement{Ack: true}, nil
}

func (p *Patient) sumAggregatedShares() int64 {
    var sum int64
    for _, share := range p.receivedShares {
        sum += share
    }
    return sum
}

func (p *Patient) StartPatientServer(wg *sync.WaitGroup) {
    defer wg.Done()
    lis, err := net.Listen("tcp", p.portAddress)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

	tlsCreds, err := loadTLSCredentials(p.patientName)
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}


    grpcServer := grpc.NewServer(
		grpc.Creds(tlsCreds),
	)
	pb.RegisterSendShareServiceServer(grpcServer, p)
    log.Println("Patient Server is running on port", p.portAddress)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

// Send a share to another patient
func (p *Patient) SendShareToOtherPatient(ctx context.Context, share int64, otherPatientName string) {
    log.Printf("[%s] Attempting to send share to [%s] at %s", p.patientName, otherPatientName, p.otherPatients[otherPatientName])
	
	tlsCreds, err := loadTLSCredentials(p.patientName)
    if err != nil {
        log.Fatalf("[%s] Failed to load TLS credentials: %v", p.patientName, err)
    }
    
    conn, err := grpc.Dial(p.otherPatients[otherPatientName], grpc.WithTransportCredentials(tlsCreds))
    if err != nil {
        log.Fatalf("[%s] Failed to dial [%s]: %v", p.patientName, otherPatientName, err)
    }
    defer conn.Close()

    client := pb.NewSendShareServiceClient(conn)
    ack, err := client.SendShare(ctx, &pb.Share{ShareOfSecret: share})
    if err != nil {
        log.Fatalf("[%s] failed to send share to [%s]: %v", p.patientName, otherPatientName, err)
    }

    log.Printf("[%s] successfully sent share to [%s]. Acknowledgement: %v", p.patientName, otherPatientName, ack.Ack)
}

// Send the final aggregated share from the patient to the hospital
func (p *Patient) SendAggregatedShareToHospital(ctx context.Context, aggregatedShare int64) {
    tlsCreds, err := loadTLSCredentials(p.patientName)
    if err != nil {
        log.Fatalf("Failed to load TLS credentials: %v", err)
    }
    
    conn, err := grpc.Dial(p.ServerAddress, grpc.WithTransportCredentials(tlsCreds))
    if err != nil {
        log.Fatalf("Failed to dial: %v", err)
    }
    defer conn.Close()

    client := pb.NewSendAggregatedShareServiceClient(conn)
    ack, err := client.SendAggregatedShare(ctx, &pb.AggregatedShare{AggregatedShareOfSecret: aggregatedShare})
    if err != nil {
        log.Fatalf("[%s] failed to send aggregated share: %v", p.patientName, err)
    }
    log.Printf("[%s] sent aggregated share to hospital: %v", p.patientName, ack.Ack)
}

func (p *Patient) calculateShares() (int64, int64, int64) {
    // Start off by generating two random integers
    share1 := rand.Int63()
    share2 := rand.Int63()

    // Calculate the third share
    share3 := p.privateInput - share1 - share2

    return share1, share2, share3
}

func loadTLSCredentials(patientName string) (credentials.TransportCredentials, error) {
    // log.Printf("[%s] Loading CA certificate...", patientName)
	
	// Load the certificate of the CA who signed the client's certificate
    pemServerCA, err := os.ReadFile("cert/ca-cert.pem")
    if err != nil {
		log.Fatalf("[%s] Failed to read CA certificate: %v", patientName, err)
        return nil, err
    }

    certPool := x509.NewCertPool()
    if !certPool.AppendCertsFromPEM(pemServerCA) {
		log.Fatalf("[%s] Failed to append CA certificate to pool", patientName)
        return nil, fmt.Errorf("failed to add server CA's certificate")
    }
	// log.Printf("[%s] Successfully loaded CA certificate.", patientName)

    // Load the client's certificate and private key
    clientCert, err := loadClientCertificate(patientName)
    if err != nil {
		log.Fatalf("[%s] Failed to load client certificate: %v", patientName, err)
        return nil, err
    }
	// log.Printf("[%s] Loaded client certificate and key.", patientName)

    // Create the credentials and return it
    config := &tls.Config{
        Certificates: 	[]tls.Certificate{clientCert},
		RootCAs:      	certPool,
		ClientAuth:   	tls.RequireAndVerifyClientCert,
        ClientCAs:      certPool,
    }

	// log.Printf("[%s] TLS credentials configured successfully.", patientName)
    return credentials.NewTLS(config), nil
}

func loadClientCertificate(patientName string) (tls.Certificate, error) {
    certFile := fmt.Sprintf("cert/%s-cert.pem", patientName)
    keyFile := fmt.Sprintf("cert/%s-key.pem", patientName)
	// log.Printf("[%s] Loading client certificate and key from %s and %s...", patientName, certFile, keyFile)
    
	clientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
		log.Fatalf("[%s] Failed to load client certificate and key: %v", patientName, err)
        return tls.Certificate{}, err
    }
	// log.Printf("[%s] Successfully loaded client certificate and key.", patientName)
    return clientCert, nil
}

func main() {
    patientName := flag.String("name", "", "Name of the patient")
    portAddress := flag.String("port", "", "Port address of the patient")
    serverAddress := flag.String("server", "localhost:8080", "Address of the server")
    privateInput := flag.Int64("input", 0, "Private input of the patient")
    flag.Parse()

    if *patientName == "" {
        log.Fatalf("Patient name must be specified")
    }

    otherPatients := map[string]string{
        "Alice":   "localhost:9000",
        "Bob":     "localhost:9001",
        "Charlie": "localhost:9002",
    }
    delete(otherPatients, *patientName)

    patient := &Patient{
        patientName:    *patientName,
        portAddress:    *portAddress,
        otherPatients:  otherPatients,
        ServerAddress:  *serverAddress,
        receivedShares: []int64{},
        privateInput:   *privateInput,
    }

    var wg sync.WaitGroup
    wg.Add(1)
    go patient.StartPatientServer(&wg)
    time.Sleep(10 * time.Second)

    // log.Println("All patients are up and running")

    share1, share2, share3 := patient.calculateShares()
    log.Printf("[%s] calculated shares: %d, %d, %d", patient.patientName, share1, share2, share3)

    if patient.patientName == "Alice" {
        patient.receivedShares = append(patient.receivedShares, share1)
        time.Sleep(10 * time.Second)
        patient.SendShareToOtherPatient(context.Background(), share2, "Bob")
        patient.SendShareToOtherPatient(context.Background(), share3, "Charlie")
    } else if patient.patientName == "Bob" {
        patient.receivedShares = append(patient.receivedShares, share2)
        time.Sleep(10 * time.Second)
        patient.SendShareToOtherPatient(context.Background(), share1, "Alice")
        patient.SendShareToOtherPatient(context.Background(), share3, "Charlie")
    } else if patient.patientName == "Charlie" {
        patient.receivedShares = append(patient.receivedShares, share3)
        time.Sleep(10 * time.Second)
        patient.SendShareToOtherPatient(context.Background(), share1, "Alice")
        patient.SendShareToOtherPatient(context.Background(), share2, "Bob")
    }
    wg.Wait()
}