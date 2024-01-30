./network.sh down

cd organizations/fabric-ca/org1
rm ca-cert.pem fabric-ca-server.db IssuerPublicKey IssuerRevocationPublicKey tls-cert.pem 
rm -r msp

cd ../org2
rm ca-cert.pem fabric-ca-server.db IssuerPublicKey IssuerRevocationPublicKey tls-cert.pem 
rm -r msp

cd ../org3
rm ca-cert.pem fabric-ca-server.db IssuerPublicKey IssuerRevocationPublicKey tls-cert.pem 
rm -r msp

cd ../org4
rm ca-cert.pem fabric-ca-server.db IssuerPublicKey IssuerRevocationPublicKey tls-cert.pem 
rm -r msp
