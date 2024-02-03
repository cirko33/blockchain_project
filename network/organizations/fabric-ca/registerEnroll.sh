#!/bin/bash
function createPeerForOrg() {
  peer_num=$1
  org_num=$2
  org_name="org${org_num}"
  peer_name="peer${peer_num}"

  infoln "Registering ${peer_name} for ${org_name}"
  set -x
  fabric-ca-client register --caname ca-${org_name} --id.name ${peer_name} --id.secret ${peer_name}pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/${org_name}/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Generating the ${peer_name} msp for ${peer_name} in ${org_name}"
  set -x
  fabric-ca-client enroll -u https://${peer_name}:${peer_name}pw@localhost:${ca_port} --caname ca-${org_name} -M ${org_path}/peers/${peer_name}.${org_name}.example.com/msp --csr.hosts ${peer_name}.${org_name}.example.com --tls.certfiles ${PWD}/organizations/fabric-ca/${org_name}/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${org_path}/msp/config.yaml ${org_path}/peers/${peer_name}.${org_name}.example.com/msp/config.yaml

  infoln "Generating the ${peer_name}-tls certificates for ${peer_name} in ${org_name}"
  set -x
  fabric-ca-client enroll -u https://${peer_name}:${peer_name}pw@localhost:${ca_port} --caname ca-${org_name} -M ${org_path}/peers/${peer_name}.${org_name}.example.com/tls --enrollment.profile tls --csr.hosts ${peer_name}.${org_name}.example.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/${org_name}/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${org_path}/peers/${peer_name}.${org_name}.example.com/tls/tlscacerts/* ${org_path}/peers/${peer_name}.${org_name}.example.com/tls/ca.crt
  cp ${org_path}/peers/${peer_name}.${org_name}.example.com/tls/signcerts/* ${org_path}/peers/${peer_name}.${org_name}.example.com/tls/server.crt
  cp ${org_path}/peers/${peer_name}.${org_name}.example.com/tls/keystore/* ${org_path}/peers/${peer_name}.${org_name}.example.com/tls/server.key

  mkdir -p ${org_path}/msp/tlscacerts
  cp ${org_path}/peers/${peer_name}.${org_name}.example.com/tls/tlscacerts/* ${org_path}/msp/tlscacerts/ca.crt

  mkdir -p ${org_path}/tlsca
  cp ${org_path}/peers/${peer_name}.${org_name}.example.com/tls/tlscacerts/* ${org_path}/tlsca/tlsca.${org_name}.example.com-cert.pem

  mkdir -p ${org_path}/ca
  cp ${org_path}/peers/${peer_name}.${org_name}.example.com/msp/cacerts/* ${org_path}/ca/ca.${org_name}.example.com-cert.pem
}

function createOrg() {
  org_num=$1
  org_name="org${org_num}"
  org_path="${PWD}/organizations/peerOrganizations/${org_name}.example.com"
  ca_port="$((6 + ${org_num}))054"

  infoln "Enrolling the CA admin"
  mkdir -p organizations/peerOrganizations/${org_name}.example.com/

  export FABRIC_CA_CLIENT_HOME=${org_path}/

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:${ca_port} --caname ca-${org_name} --tls.certfiles ${PWD}/organizations/fabric-ca/${org_name}/tls-cert.pem
  { set +x; } 2>/dev/null

  echo "NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-${ca_port}-ca-${org_name}.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-${ca_port}-ca-${org_name}.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-${ca_port}-ca-${org_name}.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-${ca_port}-ca-${org_name}.pem
    OrganizationalUnitIdentifier: orderer" >${org_path}/msp/config.yaml

  infoln "Registering user"
  set -x
  fabric-ca-client register --caname ca-${org_name} --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/${org_name}/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Registering the org admin"
  set -x
  fabric-ca-client register --caname ca-${org_name} --id.name ${org_name}admin --id.secret ${org_name}adminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/${org_name}/tls-cert.pem
  { set +x; } 2>/dev/null

  for (( i=0; i<($PEER_NUMBER); i++ )); do
    createPeerForOrg $i $org_num
  done

  infoln "Generating the user msp"
  set -x
  fabric-ca-client enroll -u https://user1:user1pw@localhost:${ca_port} --caname ca-${org_name} -M ${org_path}/users/User1@${org_name}.example.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/${org_name}/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${org_path}/msp/config.yaml ${org_path}/users/User1@${org_name}.example.com/msp/config.yaml

  infoln "Generating the org admin msp"
  set -x
  fabric-ca-client enroll -u https://${org_name}admin:${org_name}adminpw@localhost:${ca_port} --caname ca-${org_name} -M ${org_path}/users/Admin@${org_name}.example.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/${org_name}/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${org_path}/msp/config.yaml ${org_path}/users/Admin@${org_name}.example.com/msp/config.yaml
}

function createOrderer() {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/ordererOrganizations/example.com
  ca_port=6054
  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/example.com

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:${ca_port} --caname ca-orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  echo "NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-${ca_port}-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-${ca_port}-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-${ca_port}-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-${ca_port}-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer" >${PWD}/organizations/ordererOrganizations/example.com/msp/config.yaml

  infoln "Registering orderer"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Registering the orderer admin"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Generating the orderer msp"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:${ca_port} --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp --csr.hosts orderer.example.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/example.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/config.yaml

  infoln "Generating the orderer-tls certificates"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:${ca_port} --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls --enrollment.profile tls --csr.hosts orderer.example.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt
  cp ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/signcerts/* ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt
  cp ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/keystore/* ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.key

  mkdir -p ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

  mkdir -p ${PWD}/organizations/ordererOrganizations/example.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem

  infoln "Generating the admin msp"
  set -x
  fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:${ca_port} --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/example.com/users/Admin@example.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/example.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/example.com/users/Admin@example.com/msp/config.yaml
}
