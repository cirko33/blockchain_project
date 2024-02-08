# blockchain_project
## Requirements
- **Linux Distro (Debian or Ubuntu)**
- **HyperLedger Fabric 2.2.6**
- **Python 3.10 or newer**
- **Node 20 or newer**

## Installation
Install **HyperLedger Fabric 2.2.6**:  `./install-fabric.sh --fabric-version 2.2.6`   
Install **Python 3.10**:  [Guide](https://docs.python-guide.org/starting/install3/linux/)  
Install **Node 20**: ``` curl https://raw.githubusercontent.com/creationix/nvm/master/install.sh | bash && nvm install 20.11.0 ```

## Build and run
- You can change number of organizations and peers in the `network.sh` file at beginning.  
- To build and run network run script `./run_network.sh [-i] [-t]` (-i flag for ledger initialization with mock data, -t is for testing but must first have -i)
- To invoke changes on peers run from `network` folder `./invoke.sh FUNC_NAME [ARG1 ARG2 ...]` and change FUNC_NAME to Smart Contract function that you want to invoke and add arguments.  
- To query data on peer run from `network` folder `./query.sh FUNC_NAME [ARG1 ARG2 ...]` and change FUNC_NAME to SmartContract function that you want to query and add arguments.  
- To run API application place yourself into folder `node_api` and run `npm i && npm start`  
- To test with API go to [link](http://localhost:3000/doc) and test with Swagger. 
- To test with script later run `./test-chaincode.sh` from `network` folder.

## Specification
Create a web application that facilitates communication with the Hyperledger Fabric network through calls to chaincode functions. The application should support the following functionalities:

- User enrollment/login
- Querying chaincode
- Invoking chaincode

The web application must use an SDK (NodeSDK, JavaSDK, GoSDK, etc.) to interact with the chaincode.

In the system, initially, there are several banks modeled as follows:

- Bank ID (unique identifier)
- Headquarters
- Year of establishment
- Tax ID (PIB)
- Lists of users and accounts (students decide how to model this relationship)

Additionally, user and account data are stored in the system. The minimum required fields for the structures mentioned in the project specification are provided, and it is allowed and desirable to expand the specification.

User has:

- User ID (unique identifier)
- First name
- Last name
- Email address
- Accounts owned by the user

Account has:

- Amount of money
- Currency
- List of cards associated with the account

The described structures are written to the WorldState during the INIT function of the chaincode or some other invoke function that sets the initial state with a minimum of 4 banks, each having at least 3 users with one or two accounts.

The chaincode functionalities include:

- Adding a new user
- Creating one or more accounts for a user
- Transferring funds between accounts within one or more banks
- Depositing money into an account
- Withdrawing money from an account

Project task requirements:

- The application for interacting with the chaincode must have the option to work with at least four certificates. This means that the application allows logging in as four different Hyperledger network users belonging to different organizations.
- The mapping of organizations to the project specification is open to interpretation. It is not mandatory to create a new user in Hyperledger Fabric for every new client registration on the web application. Working with multiple clients from the same bank using one Hyperledger Fabric certificate is allowed.
- Creation of a Fabric network with two channels:
  - The specification includes four organizations, each with four peers.
  - Generating the necessary cryptographic materials.
  - Specifying Docker containers for Fabric network elements.
  - Each channel must be joined by at least one orderer.
  - All peers exist on both channels.
  - Creating all other necessary participants in the network along with their cryptographic materials (CA, Orderer).
  - Mandatory use of CouchDB.
  - Writing additional queries that demonstrate the advantages of CouchDB and explaining why these queries are interesting during the defense. Also, providing information on how the same thing would be done in LevelDB.
  - All peers have the roles of endorsers and committers.
- Chaincode functionalities include:
  - Money transfer: Can only occur if the transferring user has sufficient funds. If the accounts have different currencies, ask the user if they want to proceed using the exchange rate provided by the National Bank of Serbia. If the user confirms, perform the conversion.
  - Money deposit: Can only be done in the currency in which the account was created.
  - Query functions allowing searches by name, surname, account number, by surname and address simultaneously.
  - Testing all implemented functionalities (through function calls or REST calls).
  - Handling all possible errors (e.g., if there is no user or account with the specified key in the WorldState, insufficient funds, etc.).
- Objects in the WorldState are stored in JSON format, and the web application also returns JSON data.
  - The SDK-based application should enable:
    - Enroll and login
    - Query
    - Invoke