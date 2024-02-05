const swaggerAutogen = require("swagger-autogen")();

const outputFile = "./swagger_output.json";
const endpointsFiles = [
    "./routes/bank.js",
    "./routes/bankAccount.js",
    "./routes/card.js",
];

swaggerAutogen(outputFile, endpointsFiles);
