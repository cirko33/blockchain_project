const prettyJSONString = (input) => {
	return JSON.stringify(JSON.parse(input.toString()), null, 2);
}

module.exports = {
    prettyJSONString
}
