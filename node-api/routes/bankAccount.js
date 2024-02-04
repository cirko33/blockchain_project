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

    console.log("test0");
    const contract = await getContract();
    console.log("id1",id1);
    console.log("id2",id2);
    console.log("test1");
    const result = await contract.submitTransaction(
      "CheckAccountCurrencies",
      id1,
      id2
    );
    console.log("test2");
    try {
      console.log("test5");
      return res.send(JSON.stringify(result));
    } catch (e) {
      console.log("test6");
      return res.send({ result: JSON.stringify(result) });
    }
  } catch (e) {
    console.error(`Error occurred: ${e}`);
    return res.send("Method invoke failed!");
  }
});

module.exports = router;
