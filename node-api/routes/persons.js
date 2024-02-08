const { Router } = require("express");

const { getContract } = require("../fabric/ledger");

const { prettyJSONString } = require("../utils/utils");

const router = Router();

router.get("/get-all-persons", async (req, res) => {
    /*#swagger.parameters['org'] = {
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
            "GetAllPersons"
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

router.get("/get-person", async (req, res) => {
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
            "GetPerson", id
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

router.post("/create-person", async (req, res) => {
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

    #swagger.parameters['name'] = {
        in: 'query',                            
        required: true,                     
        type: 'string'
    }

    #swagger.parameters['surname'] = {
        in: 'query',                            
        required: true,                     
        type: 'string'
    }

    #swagger.parameters['email'] = {
        in: 'query',                            
        required: true,                     
        type: 'string'
    }

    */
    try {
        let id;
        let name;
        let surname;
        let email;

        try {
            id = req.body["id"];
            if (id.length < 1) throw "";
            name = req.body["name"];
            if (name.length < 1) throw "";
            surname = req.body["surname"];
            if (surname.length < 1) throw "";
            email = req.body["email"];
            if (email.length < 1) throw "";
        } catch (_) {
            return res.status(400).send({ message: "Id, name, surname, and email are mandatory fields!" });
        }

        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "CreatePerson", id, name, surname, email
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

router.get("/search-persons-by-name", async (req, res) => {
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

    #swagger.parameters['name'] = {
        in: 'query',                            
        required: true,                     
        type: 'string'
    }

    */
    try {
        let name;

        try {
            name = req.query.name;
        } catch (_) {
            return res.status(400).send({ message: "Name is a mandatory field!" });
        }

        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "SearchPersonsByName", name
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

router.get("/search-persons-by-surname", async (req, res) => {
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

    #swagger.parameters['surname'] = {
        in: 'query',                            
        required: true,                     
        type: 'string'
    }

    */
    try {
        let surname;

        try {
            surname = req.query.surname;
        } catch (_) {
            return res.status(400).send({ message: "Surname is a mandatory field!" });
        }

        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "SearchPersonsBySurname", surname
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

router.get("/search-persons-by-surname-email", async (req, res) => {
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

    #swagger.parameters['surname'] = {
        in: 'query',                            
        required: true,                     
        type: 'string'
    }

    #swagger.parameters['email'] = {
        in: 'query',                            
        required: true,                     
        type: 'string'
    }

    */
    try {
        let surname;
        let email;

        try {
            surname = req.query.surname;
            email = req.query.email;
        } catch (_) {
            return res.status(400).send({ message: "Surname and email are mandatory fields!" });
        }

        const contract = await getContract(req.org, req.channel);
        const result = await contract.submitTransaction(
            "SearchPersonsBySurnameAndEmail", surname, email
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