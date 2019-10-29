package main

import (
    "bufio"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "os"
)

func main() {

privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
if err != nil {
    fmt.Println(err.Error)
    os.Exit(1)
}

publicKey := &privateKey.PublicKey
fmt.Println("Private Key: ", privateKey)
fmt.Println("Public key: ", publicKey)

//saving pem to a file
pemPrivateFile, err := os.Create("private_key.pem")
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}

var pemPrivateBlock = &pem.Block{
    Type:  "RSA PRIVATE KEY",
    Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
}


err = pem.Encode(pemPrivateFile, pemPrivateBlock)
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
pemPrivateFile.Close()
 

//loading from a pem file
privateKeyFile, err := os.Open("private_key.pem")
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
fmt.Println("Private Key : ", privateKeyImported)


fmt.Println("********************************************************")
// publicKeyImported := &privateKeyImported.PublicKey
// fmt.Println("Private Key: ", privateKeyImported)
// fmt.Println("Public key: ", publicKeyImported)

pemPublicFile, err := os.Create("public_key.pem")
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}

var pemPublicBlock = &pem.Block{
    Type:  "RSA PUBLIC KEY",
    Bytes: x509.MarshalPKCS1PublicKey(publicKey),
}


err = pem.Encode(pemPublicFile, pemPublicBlock)
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
pemPublicFile.Close()
 

//loading from a pem file
publicKeyFile, err := os.Open("public_key.pem")
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}


pemfileinfoPublic, _ := publicKeyFile.Stat()
var sizePublic int64 = pemfileinfoPublic.Size()
pembytesPublic := make([]byte, sizePublic)
bufferPublic := bufio.NewReader(publicKeyFile)
_, err = bufferPublic.Read(pembytesPublic)
dataPublic, _ := pem.Decode([]byte(pembytesPublic))
publicKeyFile.Close()

publicKeyImported, err := x509.ParsePKCS1PublicKey(dataPublic.Bytes)
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
fmt.Println("Public Key : ", publicKeyImported)


}