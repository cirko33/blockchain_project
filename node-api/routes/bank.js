const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const { prettyJSONString } = require("../utils/utils");

const router = Router();

router.get("/get-all-banks", async (req, res) => {
    /* #swagger.parameters['org'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'

    } 

    #swagger.parameters['channel'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'
    } 

    */
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
    /* #swagger.parameters['org'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'

    }

    #swagger.parameters['channel'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'
    }

    #swagger.parameters['id'] = {
        in: 'query',                            
        required: true,                     
        type: 'integer'
    }

    */
    try {
        let id;

        try {
            id = req.query.id;
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
    /* 
    #swagger.parameters['org'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'
    } 

    #swagger.parameters['channel'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'
    } 

    #swagger.parameters['body'] = {
        in: 'body',
        required: true,
        type: 'object',
        schema: { 
            id: 1,
            location: "Novi Sad",
            pib: "123456789",
            foundationYear: 1990
        }
    }

    */
    try {
        let id;
        let location;
        let pib;

        try {
            id = parseInt(req.body.id);
            if (id.length < 1) throw "";
            location = req.body.location;
            if (location.length < 1) throw "";
            pib = req.body.pib;
            if (pib.length < 1) throw "";
            foundationYear = req.body.foundationYear;
            if (foundationYear.length < 1) throw "";

        } catch (_) {
            return res.status(400).send({ message: "Id, location and pib are mandatory fields!" });
        }

        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "CreateBank", id, location, pib, foundationYear
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

router.get("/search-banks-older-than", async (req, res) => {
    /* 
    #swagger.parameters['org'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'

    } 

    #swagger.parameters['channel'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'
    }

    #swagger.parameters['year'] = {
        in: 'query',
        required: true,
        type: 'integer',
        format: 'string'
    }

    */
    try {
        let year;

        try {
            year = req.query.year;
        } catch (_) {
            return res.status(400).send({ message: "Year is a mandatory field!" });
        }

        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "GetBanksOlderThan", year
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

router.get("/search-banks-by-location", async (req, res) => {
    /* 
    #swagger.parameters['org'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'

    } 

    #swagger.parameters['channel'] = {
        in: 'query',                            
        required: false,                     
        type: 'integer'
    } 

    #swagger.parameters['location'] = {
        in: 'query',
        required: true,
        type: 'string',
        format: 'string'
    }

    */
    try {
        let location;

        try {
            location = req.query.location;
        } catch (_) {
            return res.status(400).send({ message: "Location is a mandatory field!" });
        }

        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "GetBanksByLocation", location
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