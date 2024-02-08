const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const router = Router();


router.post("/create-card", async (req, res) => {
    try {
      let id;
        let personId;
        let balance;
        let currency;
  
      try {
        id = req.body["id"];
        if (id.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "Id is a mandatory field!" });
      }
  
      try {
        personId = req.body["personId"];
        if (personId.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "PersonId is a mandatory field!" });
      }
  
      try {
        balance = req.body["balance"];
        if (balance.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "Balance is a mandatory field!" });
      }
  
      try {
        currency = req.body["currency"];
        if (currency.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "Currency is a mandatory field!" });
      }
  
      const contract = await getContract();
      const result = await contract.submitTransaction(
        "CreateBankAccount",
        id,
        personId,
        currency,
        balance
      );
      try {
        return res.send(prettyJSONString(result));
      } catch (e) {
        return res.send({ result: prettyJSONString(result) });
      }
    } catch (e) {
      console.error(`Error occurred: ${e}`);
      return res.send("Method invoke failed! " + e.message);
    }
  });


  router.post("/remove-card", async (req, res) => {
    try {
      let id;
        let personId;
        let balance;
        let currency;
  
      try {
        id = req.body["id"];
        if (id.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "Id is a mandatory field!" });
      }
  
      try {
        personId = req.body["personId"];
        if (personId.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "PersonId is a mandatory field!" });
      }
  
      try {
        balance = req.body["balance"];
        if (balance.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "Balance is a mandatory field!" });
      }
  
      try {
        currency = req.body["currency"];
        if (currency.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "Currency is a mandatory field!" });
      }
  
      const contract = await getContract();
      const result = await contract.submitTransaction(
        "CreateBankAccount",
        id,
        personId,
        currency,
        balance
      );
      try {
        return res.send(prettyJSONString(result));
      } catch (e) {
        return res.send({ result: prettyJSONString(result) });
      }
    } catch (e) {
      console.error(`Error occurred: ${e}`);
      return res.send("Method invoke failed! " + e.message);
    }
  });

  module.exports = router;