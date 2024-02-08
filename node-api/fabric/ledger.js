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

let gateways = [];

const getGateway = async (orgNum) => {
  if (gateways != null && gateways.length > 0) {
    return gateways[orgNum - 1];
  }

  try {
    const wallet = await buildWallet();
    for (let i = 1; i <= numOfOrgs; i++) {
      const ccp = buildCCPOrg(i);
      const caClient = buildCAClient(ccp, `ca.org${i}.example.com`);
      await enrollUser(caClient, wallet, `Org${i}MSP`);

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
  } catch (error) {
    console.error(`Failed to run the application: ${error}`);
  }

  return gateways[orgNum - 1];
};

const getNetwork = async (orgNum, channelNum) => {
  try {
    return await (await getGateway(orgNum)).getNetwork("channel" + channelNum);
  } catch (error) {
    console.error(`Failed to get network: ${error}`);
    return null;
  }
};

const getContract = async (orgNum, channelNum) => {
  if (orgNum < 1 || orgNum > numOfOrgs) {
    throw new Error("Invalid organization number");
  }

  if (channelNum < 1 || channelNum > numOfChannels) {
    throw new Error("Invalid channel number");
  }

  try {
    return await (await getNetwork(orgNum, channelNum)).getContract("basic" + channelNum);
  } catch (error) {
    console.error(`Failed to get contract: ${error}`);
    return null;
  }
};

module.exports = {
  getContract,
  getNetwork,
  getGateway,
};
