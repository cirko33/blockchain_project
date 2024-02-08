const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const router = Router();


router.post("/create-card", async (req, res) => {
    try {
        let cardNumber;
        let id;
        let bankAccountId;
  
      try {
        id = req.body["id"];
        if (id.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "Id is a mandatory field!" });
      }
  
      try {
        cardNumber = req.body["cardNumber"];
        if (cardNumber.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "cardNumber is a mandatory field!" });
      }
  
      try {
        bankAccountId = req.body["bankAccountId"];
        if (bankAccountId.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "bankAccountId is a mandatory field!" });
      }
  
      const contract = await getContract(req.org, req.channel);
      const result = await contract.submitTransaction(
        "CreateCard",
        cardNumber,
        id,
        bankAccountId
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
        let bankAccountId;
  
      try {
        id = req.body["id"];
        if (id.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "id is a mandatory field!" });
      }
  
      try {
        bankAccountId = req.body["bankAccountId"];
        if (bankAccountId.length < 1) throw "";
      } catch (_) {
        return res.status(400).send({ message: "bankAccountId is a mandatory field!" });
      }
  
      const contract = await getContract(req.org, req.channel);
      const result = await contract.submitTransaction(
        "RemoveCard",
        id,
        bankAccountId
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