const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const { prettyJSONString } = require("../utils/utils");

const router = Router();

router.post("/create-bank-account", async (req, res) => {
  try {
    let id;
    let bankId;
	  let personId;
	  let balance;
	  let currency;

    try {
      id = req.body["id"];
      if (id.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "id is a mandatory field!" });
    }

    try {
      personId = req.body["personId"];
      if (personId.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "personId is a mandatory field!" });
    }

    try {
      bankId = req.body["bankId"];
      if (bankId.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "bankId is a mandatory field!" });
    }

    try {
      balance = req.body["balance"];
      if (balance.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "balance is a mandatory field!" });
    }

    try {
      currency = req.body["currency"];
      if (currency.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "currency is a mandatory field!" });
    }

    const contract = await getContract(req.org, req.channel);
    const result = await contract.submitTransaction(
      "CreateBankAccount",
      id,
      personId,
      bankId,
      currency,
      balance,
    );
    try {
      return res.send(prettyJSONString(result));
    } catch (e) {
      return res.send({ result: prettyJSONString(result) });
    }
  } catch (e) {
    console.error(`Error occurred: ${e}`);
    res.status(500).send("Method invoke failed! " + e.message);
  }
});

router.post("/transfer-funds", async (req, res) => {
  try {
    let fromAccountId;
    let toAccountId;
    let amount;
    let convert;

    try {
      fromAccountId = req.body["fromAccountId"];
      if (fromAccountId.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "FromAccountId is a mandatory field!" });
    }

    try {
      toAccountId = req.body["toAccountId"];
      if (toAccountId.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "ToAccountId is a mandatory field!" });
    }

    try {
      amount = req.body["amount"];
      if (amount.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "Amount is a mandatory field!" });
    }

    try {
      convert = req.body["convert"];
      if (convert.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "Convert is a mandatory field!" });
    }

    const contract = await getContract(req.org, req.channel);

    const sameCurrency = await contract.submitTransaction(
      "CheckAccountCurrencies",
      fromAccountId,
      toAccountId
    );

    try {
      if(!sameCurrency && !convert)
      return res.status(400).send({ message: "Accounts have different currencies!" });
    } catch (e) {
      return res.send({ result: prettyJSONString(result) });
    }


    const result = await contract.submitTransaction(
      "TransferFunds",
      fromAccountId,
      toAccountId,
      amount
    );

    try {
      return res.send(prettyJSONString(result));
    } catch (e) {
      return res.send({ result: prettyJSONString(result) });
    }
  } catch (e) {
    console.error(`Error occurred: ${e}`);
    res.status(500).send("Method invoke failed! " + e.message);
  }
});

router.post("/withdraw-funds", async (req, res) => {
  try {
    let accountId;
    let amount;

    try {
      accountId = req.body["accountId"];
      if (accountId.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "AccountId is a mandatory field!" });
    }

    try {
      amount = req.body["amount"];
      if (amount.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "Amount is a mandatory field!" });
    }
    const contract = await getContract(req.org, req.channel);
    const result = await contract.submitTransaction(
      "WithdrawFunds",
      accountId,
      amount
    );
    try {
      return res.send(prettyJSONString(result));
    } catch (e) {
      return res.send({ result: prettyJSONString(result) });
    }
  } catch (e) {
    console.error(`Error occurred: ${e}`);
    res.status(500).send("Method invoke failed! " + e.message);
  }
});

router.post("/deposit-funds", async (req, res) => {
  try {
    let accountId;
    let currency;
    let amount;

    try {
      accountId = req.body["accountId"];
      if (accountId.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "AccountId is a mandatory field!" });
    }

    try {
      currency = req.body["currency"];
      if (currency.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "Currency is a mandatory field!" });
    }

    try {
      amount = req.body["amount"];
      if (amount.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "Amount is a mandatory field!" });
    }

    const contract = await getContract(req.org, req.channel);
    const result = await contract.submitTransaction(
      "DepositFunds",
      accountId,
      currency,
      amount
    );
    try {
      return res.send(prettyJSONString(result));
    } catch (e) {
      return res.send({ result: prettyJSONString(result) });
    }
  } catch (e) {
    console.error(`Error occurred: ${e}`);
    res.status(500).send("Method invoke failed! " + e.message);
  }
});


router.patch("/check-account-currencies", async (req, res) => {
  try {
    let fromAccountId;
    let recipientId;

    try {
      fromAccountId = req.body["fromAccountId"];
      if (fromAccountId.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "fromAccountId is a mandatory field!" });
    }

    try {
      recipientId = req.body["recipientId"];
      if (recipientId.length < 1) throw "";
    } catch (_) {
      return res.status(400).send({ message: "recipientId is a mandatory field!" });
    }

    const contract = await getContract(req.org, req.channel);
    const result = await contract.submitTransaction(
      "CheckAccountCurrencies",
      fromAccountId,
      recipientId
    );
    try {
      return res.send(prettyJSONString(result));
    } catch (e) {
      return res.send({ result: prettyJSONString(result) });
    }
  } catch (e) {
    console.error(`Error occurred: ${e}`);
    res.status(500).send("Method invoke failed! " + e.message);
  }
});

module.exports = router;
