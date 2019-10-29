package server

import (
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
    "reflect"
	// "strconv"
	"unsafe"
	// "bytes"

)

// struct to maintain the state
type validationServer struct {
}

// Get a new instance of validationServer
func NewValidationServer() *validationServer {
	return &validationServer{}
}

func BytesToString(b []byte) string {
    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
    sh := reflect.StringHeader{bh.Data, bh.Len}
    return *(*string)(unsafe.Pointer(&sh))
}


// Implementation of the ValidationServer.SubmitRequest() RPC call
func (vs *validationServer) SubmitRequest(ctx context.Context, vrequest *mb.ValidationRequest) (*mb.ValidationResponse, error) {
	// DO NOT CHANGE THIS PRINTF STATEMENT
	log.Printf("Validated [MSGID:%d, MSG:%s]", vrequest.MsgId, vrequest.Msg)

	//Approach-sign the message with private key of Validation Server

	//loading private key of validation server
	privateKeyFile, err := os.Open("/home/hrishabh/go/src/new/pkiValidator/private_key.pem")
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}


	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	// log.Println("Private Key : ", privateKeyImported)

	//generating signature

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := []byte(vrequest.Msg)
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(
	    rand.Reader, 
	    privateKeyImported, 
	    newhash, 
	    hashed, 
	    &opts,
	)

	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}

	fmt.Printf("PSS Signature : %x\n", signature)

	sig := hex.EncodeToString(signature)
	fmt.Printf("sig: %s\n", sig)

	
	// tmp, _ := hex.DecodeString(sig)
	// fmt.Println(reflect.DeepEqual(tmp, signature))


	//appending msg to the signature

	// msg_to_send := sig + " " + vrequest.Msg

	return &mb.ValidationResponse{
		Msg:         vrequest.Msg,
		MsgId:       vrequest.MsgId,
		ReturnValue: mb.ValidationResponse_SUCCESS,
		Signature: []byte(sig),
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

	vServer := NewValidationServer()

	mb.RegisterValidationServer(grpcServer, vServer)

	log.Printf("Started validation server on port:%d, host:%s", port, host)
	grpcServer.Serve(lis)
}
