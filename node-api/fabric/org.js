const { Client } = require("fabric-common");
const { Wallet, Wallets } = require("fabric-network");
const fs = require("fs");
const path = require("path");
const dotenv = require("dotenv");

dotenv.config({ path: path.resolve(__dirname, "env", ".env") });

const buildCCPOrg = (n) => {
  const organizationsPath =
    process.env["ORGANIZATIONS_PATH"] ||
    path.resolve(__dirname, "..", "..", "network", "organizations");
  const ccpPath = path.resolve(
    organizationsPath,
    "peerOrganizations",
    `org${n}.example.com`,
    `connection-org${n}.json`
  );

  const fileExists = fs.existsSync(ccpPath);
  if (!fileExists) {
    throw new Error(`no such file or directory: ${ccpPath}`);
  }
  const contents = fs.readFileSync(ccpPath, "utf8");

  const ccp = JSON.parse(contents);
  console.log(`Loaded the network configuration located at ${ccpPath}`);

  return ccp;
};

const buildWallet = async (walletPath) => {
  let wallet;
  if (walletPath) {
    wallet = await Wallets.newFileSystemWallet(walletPath);
    console.log(`Built a file system wallet at ${walletPath}`);
  } else {
    wallet = await Wallets.newInMemoryWallet();
    console.log("Built an in memory wallet");
  }

  return wallet;
};

const prettyJSONString = (inputString) => {
  if (inputString) {
    return JSON.stringify(JSON.parse(inputString), null, 2);
  } else {
    return inputString;
  }
};

module.exports = { buildCCPOrg, buildWallet, prettyJSONString };
