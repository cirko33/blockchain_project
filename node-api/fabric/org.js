const { Wallets } = require("fabric-network");
const fs = require("fs");
const path = require("path");
const dotenv = require("dotenv");

dotenv.config({ path: path.resolve(__dirname, "env", "development.env") });

const buildCCPOrg = (n) => {
  const organizationsPath =
    process.env["ORGANIZATIONS_PATH"] ||
    path.resolve(__dirname, "..", "network", "organizations");
  const ccpPath = path.resolve(
    organizationsPath,
    "peerOrganizations",
    `org${n}.example.com`,
    `connection-org${n}.json`
  );

  const fileExists = fs.existsSync(ccpPath);
  if (!fileExists) {
    throw new Error(`No such file or directory: ${ccpPath}`);
  }
  const contents = fs.readFileSync(ccpPath, "utf8");

  const ccp = JSON.parse(contents);
  console.log(`Loaded the network configuration located at ${ccpPath}`);

  return ccp;
};

const buildWallet = async () => {
  let wallet = await Wallets.newInMemoryWallet();
  console.log("Built an in memory wallet");
  return wallet;
};

module.exports = { buildCCPOrg, buildWallet };
