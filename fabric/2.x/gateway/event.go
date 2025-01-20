package main
//
//import (
//	"bytes"
//	"context"
//	"crypto/x509"
//	"encoding/json"
//	"fmt"
//	"github.com/hyperledger/fabric-gateway/pkg/client"
//	"github.com/hyperledger/fabric-gateway/pkg/identity"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials"
//	"os"
//	"path"
//	"time"
//)
//
///**
// * @Description $
// * @Author hll
// * @Date 2023/10/9 14:55
// **/
//const (
//	channelName      = "hll8"
//	chaincodeName    = "hll8gateway"
//	mspID            = "Node2MSP"
//	cryptoPath       = "certs/peerOrganizations/node2.private.bsnbase.com"
//	certPath         = cryptoPath + "/users/Admin@node2.private.bsnbase.com/msp/signcerts/cert.pem"
//	keyPath          = cryptoPath + "/users/Admin@node2.private.bsnbase.com/msp/keystore/"
//	tlsCertPath      = "certs/peerOrganizations/node2.private.bsnbase.com/peers/peer1.node2.private.bsnbase.com/tls/ca.crt"
//	node2TlsCertPath = "certs/peerOrganizations/node2.private.bsnbase.com/peers/peer1.node2.private.bsnbase.com/tls/ca.crt"
//	peerEndpoint     = "10.0.51.35:13002"
//	gatewayPeer      = "peer1.node2.private.bsnbase.com"
//)
//
//func main() {
//	fmt.Println(os.Getwd())
//	// The gRPC client connection should be shared by all Gateway connections to this endpoint
//	clientConnection := newGrpcConnection()
//	defer clientConnection.Close()
//
//	id := newIdentity()
//	sign := newSign()
//
//	// Create a Gateway connection for a specific client identity
//	gw, err := client.Connect(
//		id,
//		client.WithSign(sign),
//		client.WithClientConnection(clientConnection),
//		// Default timeouts for different gRPC calls
//		client.WithEvaluateTimeout(1005*time.Second),
//		client.WithEndorseTimeout(1005*time.Second),
//		client.WithSubmitTimeout(1005*time.Second),
//		client.WithCommitStatusTimeout(100*time.Minute),
//	)
//	if err != nil {
//		panic(err)
//	}
//	defer gw.Close()
//	network := gw.GetNetwork(channelName)
//	_ = network.GetContract(chaincodeName)
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	BlockEventListening(ctx, network)
//	startChaincodeEventListening(ctx, network, chaincodeName)
//
//	time.Sleep(500 * time.Second)
//}
//
//func BlockEventListening(ctx context.Context, network *client.Network) {
//	fmt.Println("\n*** Start Block Event listening")
//
//	events, err := network.BlockEvents(ctx)
//	if err != nil {
//		panic(fmt.Errorf("failed to start Block Event listening: %w", err))
//	}
//
//	go func() {
//		for event := range events {
//			//asset := formatJSON(event.GetData().Data)
//			fmt.Printf("\n<-- Block Event received: BlockNumber: %d \n", event.GetHeader().GetNumber())
//		}
//	}()
//}
//
//func startChaincodeEventListening(ctx context.Context, network *client.Network, chaincodeName string) {
//	fmt.Println("\n*** Start chaincode event listening")
//
//	events, err := network.ChaincodeEvents(ctx, chaincodeName)
//	if err != nil {
//		panic(fmt.Errorf("failed to start chaincode event listening: %w", err))
//	}
//
//	go func() {
//		for event := range events {
//			asset := formatJSON(event.Payload)
//			fmt.Printf("\n<-- Chaincode event received: %s - %s\n", event.EventName, asset)
//		}
//	}()
//}
//func formatJSON(data []byte) string {
//	var prettyJSON bytes.Buffer
//	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
//		panic(fmt.Errorf("failed to parse JSON: %w", err))
//	}
//	return prettyJSON.String()
//}
//
//// newGrpcConnection creates a gRPC connection to the Gateway server.
//func newGrpcConnection() *grpc.ClientConn {
//	// TLS 证书
//	certificate, err := loadCertificate(tlsCertPath)
//	if err != nil {
//		panic(err)
//	}
//
//	certPool := x509.NewCertPool()
//	certPool.AddCert(certificate)
//	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)
//
//	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
//	if err != nil {
//		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
//	}
//
//	return connection
//}
//
//// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
//// 交易的客户端标识
//func newIdentity() *identity.X509Identity {
//	certificate, err := loadCertificate(certPath)
//	if err != nil {
//		panic(err)
//	}
//
//	id, err := identity.NewX509Identity(mspID, certificate)
//	if err != nil {
//		panic(err)
//	}
//
//	return id
//}
//
//func loadCertificate(filename string) (*x509.Certificate, error) {
//	certificatePEM, err := os.ReadFile(filename)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read certificate file: %w", err)
//	}
//	return identity.CertificateFromPEM(certificatePEM)
//}
//
//// newSign creates a function that generates a digital signature from a message digest using a private key.
//func newSign() identity.Sign {
//	files, err := os.ReadDir(keyPath)
//	if err != nil {
//		panic(fmt.Errorf("failed to read private key directory: %w", err))
//	}
//	privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))
//
//	if err != nil {
//		panic(fmt.Errorf("failed to read private key file: %w", err))
//	}
//
//	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
//	if err != nil {
//		panic(err)
//	}
//
//	sign, err := identity.NewPrivateKeySign(privateKey)
//	if err != nil {
//		panic(err)
//	}
//
//	return sign
//}
