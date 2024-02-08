const dotenv = require("dotenv");
const { Gateway } = require("fabric-network");
const path = require("path");
const { buildCAClient, registerAndEnrollUser, enrollUser } = require("./ca");
const { buildCCPOrg, buildWallet } = require("./org");

dotenv.config({ path: path.resolve("env", "development.env") });
const env = process.env;

const numOfOrgs = parseInt(env["NUM_OF_ORGS"]) || 4;
const numOfChannels = parseInt(env["NUM_OF_CHANNELS"]) || 2;

const orgUserId = env["CA_USER"] || "user1";

let contracts = [];
let gateways = [];
let networks = [];

const getGateway = async (orgNum, channelNum) => {

  if (gateways != null && gateways.length > 0) {
    return gateways[(channelNum - 1) * numOfOrgs + orgNum - 1];
  }

  try {
    const wallet = await buildWallet();
    for (let i = 1; i <= numOfChannels; i++) {
      for (let j = 1; j <= numOfOrgs; j++) {
        const ccp = buildCCPOrg(j);
        const caClient = buildCAClient(ccp, `ca.org${j}.example.com`);
        await enrollUser(caClient, wallet, `Org${j}MSP`);

        const temp_gw = new Gateway();
        try {
          await temp_gw.connect(ccp, {
            wallet,
            identity: orgUserId,
            discovery: { enabled: true, asLocalhost: true },
          });

          gateways.push(temp_gw);
        } catch (error) {
          console.error(`Failed to connect: ${error}`);
        }
      }
    }
  } catch (error) {
    console.error(`Failed to run the application: ${error}`);
  }

  return gateways[(channelNum - 1) * numOfOrgs + orgNum - 1];
};

const getNetwork = async (orgNum, channelNum) => {
  if (networks != null && networks.length > 0) {
    return networks[(channelNum - 1) * numOfOrgs + orgNum - 1];
  }

  try {
    for (let i = 1; i <= numOfChannels; i++) {
      for (let j = 1; j <= numOfOrgs; j++) {
        networks.push(await (await getGateway(j, i)).getNetwork("channel" + i));
      }
    }

  } catch (error) {
    console.error(`Failed to get network: ${error}`);
  }

  return networks[(channelNum - 1) * numOfOrgs + orgNum - 1];
};

const getContract = async (orgNum, channelNum) => {
  if (orgNum < 1 || orgNum > numOfOrgs) {
    throw new Error("Invalid organization number");
  }

  if (channelNum < 1 || channelNum > numOfChannels) {
    throw new Error("Invalid channel number");
  }

  if (contracts != null && contracts.length > 0) {
    console.log("🚀 ~ getContract ~ contracts:", contracts)
    return contracts[(channelNum - 1) * numOfOrgs + orgNum - 1];
  }

  try {
    for (let i = 1; i <= numOfChannels; i++) {
      for (let j = 1; j <= numOfOrgs; j++) {
        contracts.push(await (await getNetwork(j, i)).getContract("basic" + i));
      }
    }
  } catch (error) {
    console.error(`Failed to get contract: ${error}`);
  }

  return contracts[(channelNum - 1) * numOfOrgs + orgNum - 1];
};

module.exports = {
  getContract,
  getNetwork,
  getGateway,
};
