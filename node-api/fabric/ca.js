const dotenv = require("dotenv");
const FabricCAServices = require("fabric-ca-client");
const path = require("path");

dotenv.config({ path: path.resolve("env", "development.env") });
const env = process.env;

const userId = env["CA_USER"] || "user1";
const userPasswd = env["CA_USER_PASSWORD"] || "user1pw";

const buildCAClient = (ccp, caHostName) => {
  const caInfo = ccp.certificateAuthorities[caHostName];
  const caTLSCACerts = caInfo.tlsCACerts.pem;
  const caClient = new FabricCAServices(
    caInfo.url,
    {
      trustedRoots: caTLSCACerts,
      verify: false,
    },
    caInfo.caName
  );

  console.log(`Built a CA Client named ${caInfo.caName}`);
  return caClient;
};

const enrollUser = async (caClient, wallet, orgMspId) => {
  try {
    const userInOrg = `${orgMspId.toLowerCase()}_${userId}`
    const identity = await wallet.get(userInOrg);
    if (identity) {
      console.log(
        "An identity for the admin user already exists in the wallet"
      );
      return;
    }

    const enrollment = await caClient.enroll({
      enrollmentID: userId,
      enrollmentSecret: userPasswd,
    });
    const x509Identity = {
      credentials: {
        certificate: enrollment.certificate,
        privateKey: enrollment.key.toBytes(),
      },
      mspId: orgMspId,
      type: "X.509",
    };

    await wallet.put(userId, x509Identity);
    console.log(
      "Successfully enrolled user and imported it into the wallet"
    );
  } catch (error) {
    console.error(`Failed to enroll user : ${error}`);
  }
};

module.exports = {
  buildCAClient,
  enrollUser,
};
