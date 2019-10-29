package server

import (
	// "strings"
	"context"
	"fmt"
	"log"
	"net"

	mb "daad/protos/master"

	"google.golang.org/grpc"

    "bufio"
    "crypto"
  	"crypto/rsa"
   	"crypto/rand"
    "crypto/x509"
	// "crypto/sha256"
    "encoding/pem"
    "encoding/hex"
    "os"
	"crypto/sha512"

)

// struct to maintain the state
type commitServer struct {
}

// Get a new instance of commitServer
func NewCommitServer() *commitServer {
	return &commitServer{}
}

func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		// log.Error(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		// log.Error(err)
	}
	return plaintext
} 


// Implementation of the CommitServer.SubmitRequest() RPC call
func (vs *commitServer) SubmitRequest(ctx context.Context, crequest *mb.ValidationResponse) (*mb.CommitResponse, error) {
	// DO NOT CHANGE THIS PRINTF STATEMENT

	log.Printf("Committed[MSGID:%d, MSG:%s]", crequest.MsgId, crequest.Msg)

	// INSERT CODE HERE

	//loading public key of validation server 

	publicKeyFile, err := os.Open("/home/hrishabh/go/src/new/pkiValidator/public_key.pem")
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}


	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKCS1PublicKey(data.Bytes)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	// log.Println("Public Key : ", publicKeyImported)


	//seperating signature with the msg
	// splt_string := strings.Split(crequest.Msg, " ")
	// fmt.Printf("Encoded String: %s\n", splt_string[0])

	enc_msg, err := hex.DecodeString(string(crequest.Signature))
	if err != nil {
		log.Println("Error in decode string")

		//could be because of adversary
		//handling for adversary
		log.Println("Not commited")
		return &mb.CommitResponse{
		ReturnValue: mb.CommitResponse_FAILURE,
		}, nil
	}

	// org_msg := strings.Join(splt_string[1:], " ")
	// fmt.Printf("Original String: %s\n", org_msg)


	//verifying the signature using public key of validation server

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	newhash := crypto.SHA256
	PSSmessage := []byte(crequest.Msg)
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	err = rsa.VerifyPSS(
	    publicKeyImported, 
	    newhash, 
	    hashed,
	    enc_msg,
	    &opts,
	)


	//check for error
	if err != nil {
		log.Println("Not commited")
		return &mb.CommitResponse{
		ReturnValue: mb.CommitResponse_FAILURE,
		}, nil
	}	

	return &mb.CommitResponse{
		ReturnValue: mb.CommitResponse_SUCCESS,
	}, nil

}

func Main(host string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		log.Fatalf("Failed to listen to port: %d, on host: %s", port, host)
	}

	// server option TLS or no TLS
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	vServer := NewCommitServer()

	mb.RegisterCommitServer(grpcServer, vServer)

	log.Printf("Started commit server on port:%d, host:%s", port, host)
	grpcServer.Serve(lis)
}
