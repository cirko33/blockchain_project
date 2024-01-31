from sys import argv



def make_peers(port1, port2, orgs, peers):
  text = ""
  depends_on = ""

  cli = """
  cli:
    container_name: cli
    image: hyperledger/fabric-tools:latest
    tty: true
    stdin_open: true
    environment:
      - GOPATH=/opt/gopath
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      - FABRIC_LOGGING_SPEC=INFO
      #- FABRIC_LOGGING_SPEC=DEBUG
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: /bin/bash
    volumes:
      - /var/run/:/host/var/run/
      - ../organizations:/opt/gopath/src/github.com/hyperledger/fabric/peer/organizations
      - ../scripts:/opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/
    depends_on:
{0}
    networks:
      - test

volumes:
{1}
  """

  start = """
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

version: '2.1'
networks:
  test:
    name: fabric_test

services:
  orderer.example.com:
    extends:
      file: common-all.yaml
      service: common-orderer
  """

  form = """  
  peer{0}.org{1}.example.com:
    container_name: peer{0}.org{1}.example.com
    image: hyperledger/fabric-peer:latest
    extends:
      file: common-all.yaml
      service: common-peer
    environment:
      - CORE_PEER_ID=peer{0}.org{1}.example.com
      - CORE_PEER_ADDRESS=peer{0}.org{1}.example.com:{2}
      - CORE_PEER_LISTENADDRESS=0.0.0.0:{2}
      - CORE_PEER_CHAINCODEADDRESS=peer{0}.org{1}.example.com:{3}
      - CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:{3}
      - CORE_PEER_GOSSIP_BOOTSTRAP=peer{0}.org{1}.example.com:{2}
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer{0}.org{1}.example.com:{2}
      - CORE_PEER_LOCALMSPID=Org{1}MSP
      - CORE_OPERATIONS_LISTENADDRESS=peer{0}.org{1}.example.com:{4}
    volumes:
      - /var/run/docker.sock:/host/var/run/docker.sock
      - ../organizations/peerOrganizations/org{1}.example.com/peers/peer{0}.org{1}.example.com/msp:/etc/hyperledger/fabric/msp
      - ../organizations/peerOrganizations/org{1}.example.com/peers/peer{0}.org{1}.example.com/tls:/etc/hyperledger/fabric/tls
      - peer{0}.org{1}.example.com:/var/hyperledger/production
    ports:
      - {2}:{2}
      - {4}:{4}
  """
  volumes = ""

  for i in range(0, peers):
    for j in range(1, orgs + 1):
        depends_on += "      - peer{0}.org{1}.example.com\n".format(i, j)
        volumes += "  peer{0}.org{1}.example.com:\n".format(i, j)
        temp = port1 + 1000 * j + 5 * i
        text += form.format(i, j, temp, temp + 1, port2 + 1000 * j + 1 * i)

  final = start + text + cli.format(depends_on, volumes)
  return final

def make_couch(port, orgs, peers):
  start = """
version: '2.1'
networks:
  test:
    name: fabric_test

services:
"""

  couch = """
##### couchdb{0}{1} #####
  couchdb{0}{1}:
    extends:
      file: common-all.yaml
      service: common-couch
    container_name: couchdb{0}{1}
    ports:
      - "{2}:5984"

  peer{0}.org{1}.example.com:
    environment:
      - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb{0}{1}:5984
    depends_on:
      - couchdb{0}{1}
"""
  text = ""
  for i in range(0, peers):
    for j in range(1, orgs + 1):
      text += couch.format(i, j, port + 1000 * j + i)
  return start + text

def make_ca(port, orgs):
  text = """
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
version: '2.1'
networks:
  test:
    name: fabric_test

services:
  ca_orderer:
    image: hyperledger/fabric-ca:latest
    environment:
      - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
      - FABRIC_CA_SERVER_CA_NAME=ca-orderer
      - FABRIC_CA_SERVER_TLS_ENABLED=true
      - FABRIC_CA_SERVER_PORT=6054
      - FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS=0.0.0.0:16054
    ports:
      - "6054:6054"
      - "16054:16054"
    command: sh -c 'fabric-ca-server start -b admin:adminpw -d'
    volumes:
      - ../organizations/fabric-ca/ordererOrg:/etc/hyperledger/fabric-ca-server
    container_name: ca_orderer
    networks:
      - test
"""
  ca_template = """
  ca_org{0}:
    image: hyperledger/fabric-ca:latest
    environment:
      - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
      - FABRIC_CA_SERVER_CA_NAME=ca-org{0}
      - FABRIC_CA_SERVER_TLS_ENABLED=true
      - FABRIC_CA_SERVER_PORT={1}
      - FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS=0.0.0.0:{2}
    ports:
      - "{1}:{1}"
      - "{2}:{2}"
    command: sh -c 'fabric-ca-server start -b admin:adminpw -d'
    volumes:
      - ../organizations/fabric-ca/org{0}:/etc/hyperledger/fabric-ca-server
    container_name: ca_org{0}
    networks:
      - test
"""
  for i in range(1, orgs + 1):
    text += ca_template.format(i, port + 1000 * i, port + 1000 * i + 10000)

  return text

if __name__ == "__main__":
  peer_port1, peer_port2 = 6051, 6400
  couch_port=6984
  ca_port=6054
  orgs, peers = 4, 4

  if len(argv) == 3:
     orgs, peers = int(argv[1]), int(argv[2])
     

  peer = make_peers(peer_port1, peer_port2, orgs, peers)
  couch = make_couch(couch_port, orgs, peers)
  ca = make_ca(ca_port, orgs)

  print(peer, file=open("network/docker/docker-compose-test-net.yaml", "w"))
  print(couch, file=open("network/docker/docker-compose-couch.yaml", "w"))
  print(ca, file=open("network/docker/docker-compose-ca.yaml", "w"))