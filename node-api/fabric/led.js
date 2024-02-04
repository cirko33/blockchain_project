// Convert ESM imports to CommonJS requires
const dotenv = require("dotenv");
const { Gateway } = require("fabric-network");
const Contract = require("fabric-network/lib/contract");
const Network = require("fabric-network/lib/network");
const path = require("path");
const { buildCAClient, registerAndEnrollUser, enrollAdmin } = require("./ca");
const { buildCCPOrg, buildWallet } = require("./org");

dotenv.config({ path: path.resolve("env", ".env") });
const env = process.env;

const channelName = env["CHANNEL_NAME"] || "mychannel";
const chaincodeName = env["CHAINCODE_NAME"] || "basic";
const organizationId = parseInt(env["ORGANIZATION_ID"] || "1");
const mspOrg = `Org${organizationId}MSP`;
const walletPath = path.join(__dirname, env["WALLET_PATH"] || "wallet");
const orgUserId = env["ORGANIZATION_USER_ID"] || "appUser";
const affiliation = env["AFFILIATION"] || `manufacturer.department1`;

let contract = null;
let gateway = null;
let network = null;

const getGateway = async () => {
  if (gateway != null) {
    return gateway;
  }

  try {
    const ccp = buildCCPOrg(organizationId);
    const caClient = buildCAClient(ccp, `ca.org${organizationId}.example.com`);
    const wallet = await buildWallet(walletPath);
    await enrollAdmin(caClient, wallet, mspOrg);
    await registerAndEnrollUser(
      caClient,
      wallet,
      mspOrg,
      orgUserId,
      affiliation
    );
    gateway = new Gateway();

    try {
      await gateway.connect(ccp, {
        wallet,
        identity: orgUserId,
        discovery: { enabled: true, asLocalhost: true },
      });
    } catch (error) {
      console.error(`******** FAILED to connect: ${error}`);
    }
  } catch (error) {
    console.error(`******** FAILED to run the application: ${error}`);
  }

  return gateway;
};

const getNetwork = async () => {
  if (network != null) {
    return network;
  }

  try {
    network = await (await getGateway()).getNetwork(channelName);
  } catch (error) {
    console.error(`********* Failed to get network: ${error}`);
  }

  return network;
};

const getContract = async () => {
  if (contract != null) {
    return contract;
  }

  try {
    contract = await (await getNetwork()).getContract(chaincodeName);
  } catch (error) {
    console.error(`********* Failed to get contract: ${error}`);
  }

  return contract;
};

module.exports = {
  getContract,
  getNetwork,
  getGateway,
};
