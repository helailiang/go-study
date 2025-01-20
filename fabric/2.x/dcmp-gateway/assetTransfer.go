/*
Copyright 2021 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	mspID        = "Hllorg2MSP"
	certPath     = "E:\\logs\\pki\\5ceca3eb7363236\\certs\\peerOrganizations\\hllorg2-5ceca3eb7363236.docker.bsnbase.com\\users\\Admin@hllorg2-5ceca3eb7363236.docker.bsnbase.com\\msp\\signcerts\\cert.pem"
	keyPath      = "E:\\logs\\pki\\5ceca3eb7363236\\certs\\peerOrganizations\\hllorg2-5ceca3eb7363236.docker.bsnbase.com\\users\\Admin@hllorg2-5ceca3eb7363236.docker.bsnbase.com\\msp\\keystore"
	tlsCertPath  = "E:\\logs\\pki\\5ceca3eb7363236\\certs\\peerOrganizations\\hllorg2-5ceca3eb7363236.docker.bsnbase.com\\peers\\peer2.hllorg2-5ceca3eb7363236.vdc2785556.docker.bsnbase.com\\tls\\ca.crt"
	peerEndpoint = "10.0.7.26:7051"
	gatewayPeer  = "peer1.hllorg2-5ceca3eb7363236.vdc2785556.docker.bsnbase.com"
)

var now = time.Now()
var chaincodeName = "hllContractCode1010"
var channelName = "test"

func main() {
	fmt.Println(os.Getwd())
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(1005*time.Second),
		client.WithEndorseTimeout(1005*time.Second),
		client.WithSubmitTimeout(1005*time.Second),
		client.WithCommitStatusTimeout(100*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	// Context used for event listening
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	// Listen for events emitted by subsequent transactions
	//startChaincodeEventListening(ctx, network, chaincodeName)
	//initLedger(contract)
	//createAsset(contract)
	getAllAssets(contract)

}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	// TLS 证书
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
// 交易的客户端标识
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	files, err := os.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

// Evaluate a transaction to query ledger state.
func getAllAssets(contract *client.Contract) {
	fmt.Println("\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("Get", "a")
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}

// Submit a transaction synchronously, blocking until it has been committed to the ledger.
func createAsset(contract *client.Contract) {
	fmt.Printf("\n--> 调用set() \n")

	_, err := contract.SubmitTransaction("set", "Tom", "38800")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

// Format JSON data
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}
