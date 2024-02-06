const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const router = Router();

router.get("/get-all-persons", async (req, res) => {
    try {
        const contract = await getContract();
        const result = await contract.submitTransaction(
            "GetAllPersons"
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

router.get("/get-person", async (req, res) => {
    try {
        let id;

        try {
            id = req.body["id"];
            if (id.length < 1) throw "";
        } catch (_) {
            return res.status(400).send({ message: "Id is a mandatory field!" });
        }

        const contract = await getContract();
        const result = await contract.submitTransaction(
            "GetPerson", id
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

router.post("/create-person", async (req, res) => {
    try {
        let id;
        let name;
        let surname;
        let email;

        try {
            id = req.body["id"];
            if (id.length < 1) throw "";
            name = req.body["name"]
            if (name.length < 1) throw "";
            surname = req.body["surname"]
            if (surname.length < 1) throw "";
            email = req.body["email"]
            if (email.length < 1) throw "";
        } catch (_) {
            return res.status(400).send({ message: "Id, name, surname, and email are mandatory fields!" });
        }

        const contract = await getContract();
        const result = await contract.submitTransaction(
            "CreatePerson", id, name, surname, email
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