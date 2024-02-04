const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const router = Router();

router.patch("/check-account-currencies", async (req, res) => {
  try {
    let id1;
    let id2;

    try {
      id1 = req.body["id1"];
      if (id1.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "Id1 is a mandatory field!" });
    }

    try {
      id2 = req.body["id2"];
      if (id2.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "Id2 is a mandatory field!" });
    }

    const contract = await getContract();
    const result = await contract.submitTransaction(
      "CheckAccountCurrencies",
      id1,
      id2
    );
    try {
      return res.send(JSON.stringify(result));
    } catch (e) {
      return res.send({ result: JSON.stringify(result) });
    }
  } catch (e) {
    console.error(`Error occurred: ${e}`);
    return res.send("Method invoke failed!");
  }
});

module.exports = router;
