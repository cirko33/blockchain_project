const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const { prettyJSONString } = require("../utils/utils");

const router = Router();

router.get("/get-all-banks", async (req, res) => {
    try {
        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "GetAllBanks"
        );

        try {
            return res.send(prettyJSONString(result));
        } catch (e) {
            return res.send({ result: prettyJSONString(result) });
        }
    } catch (e) {
        console.error(`Error occurred: ${e}`);
        res.status(500).send("Method invoke failed!");
    }
});

router.get("/get-bank", async (req, res) => {
    try {
        let id;

        try {
            id = req.query["id"];
            if (id.length < 1) throw "";
        } catch (_) {
            return res.status(400).send({ message: "Id is a mandatory field!" });
        }

        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "GetBank", id
        );

        try {
            return res.send(prettyJSONString(result));
        } catch (e) {
            return res.send({ result: prettyJSONString(result) });
        }
    } catch (e) {
        console.error(`Error occurred: ${e}`);
        res.status(500).send("Method invoke failed!");
    }
});

router.post("/create-bank", async (req, res) => {
    try {
        let id;
        let location;
        let pib;

        try {
            id = req.body["id"];
            if (id.length < 1) throw "";
            location = req.body["location"]
            if (location.length < 1) throw "";
            pib = req.body["pib"]
            if (pib.length < 1) throw "";
        } catch (_) {
            return res.status(400).send({ message: "Id, location and pib are mandatory fields!" });
        }

        const contract = await getContract();
        const result = await contract.submitTransaction(
            "CreateBank", id, location, pib
        );

        try {
            return res.send(prettyJSONString(result));
        } catch (e) {
            return res.send({ result: prettyJSONString(result) });
        }
    } catch (e) {
        console.error(`Error occurred: ${e}`);
        res.status(500).send("Method invoke failed!");
    }
});

module.exports = router;