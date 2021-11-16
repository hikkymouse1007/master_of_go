package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net"
	"os"
	"time"

	//"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
)

func main() {
	// ref: https://stackoverflow.com/questions/45428126/how-to-create-a-big-int-with-a-secure-random
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	fmt.Println("rand.Reader", rand.Reader)
	fmt.Println("max:", max)
	fmt.Printf("The type of max is %T\n", max)// **big.Int

	serialNumber, err := rand.Int(rand.Reader, max)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The type of serialNumber is %T\n", serialNumber)// *big.Int
	fmt.Printf("The type of *serialNumber is %T\n", *serialNumber)// big.Int
	fmt.Printf("The type of &serialNumber is %T\n", &serialNumber)// **big.Int
	fmt.Println("serialNumber", serialNumber)

	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}
	fmt.Println("subject:", subject)

	// Sample: https://golang.hotexamples.com/jp/examples/crypto.x509/Certificate/SerialNumber/golang-certificate-serialnumber-method-examples.html
	template := x509.Certificate{
		SerialNumber: serialNumber, // pointer
		Subject: subject,
		NotBefore: time.Now(),
		NotAfter: time.Now(),
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}
	fmt.Printf(" template:", template)

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	fmt.Println("pk:", pk)
	fmt.Println("pk_pub:", &pk.PublicKey)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
