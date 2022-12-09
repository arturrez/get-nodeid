package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/staking"
	"github.com/ava-labs/avalanchego/utils/hashing"
)

func LoadTLSCert(keyPath, certPath string) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func main() {
	generate := flag.Bool("generate", false, "generate cert-key pair")
	keyPath := flag.String("key", "/home/ava_0/.avalanchego/staking/staker.key", "key path")
	certPath := flag.String("cert", "/home/ava_0/.avalanchego/staking/staker.crt", "crt path")

	flag.Parse()

	if *generate {
		cBytes, kBytes, err := staking.NewCertAndKeyBytes()
		if err != nil {
			panic(err)
		}
		if err = os.WriteFile(*certPath, cBytes, 0600); err != nil {
			panic(err)
		}
		if err = os.WriteFile(*keyPath, kBytes, 0600); err != nil {
			panic(err)
		}
	}

	cert, err := LoadTLSCert(*keyPath, *certPath)
	if err != nil {
		panic(err)
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	nodeID, err := ids.ToShortID(hashing.PubkeyBytesToAddress(cert.Leaf.Raw))
	if err != nil {
		panic(err)
	}
	fmt.Printf("NodeID-%s cert:%s key:%s\n", nodeID, *certPath, *keyPath)
}
