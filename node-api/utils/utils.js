const prettyJSONString = (input) => {
	return (input == undefined ? "No data" : JSON.stringify(JSON.parse(input.toString()), null, 2));
}

module.exports = {
    prettyJSONString
}
