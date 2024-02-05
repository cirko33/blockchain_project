const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const router = Router();

router.get("/get-all-banks", async (req, res) => {
    try {
        const contract = await getContract();
        const result = await contract.submitTransaction(
            "GetAllBanks"
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