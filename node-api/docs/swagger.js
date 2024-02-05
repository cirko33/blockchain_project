const swaggerAutogen = require("swagger-autogen")();

const outputFile = "./swagger_output.json";
const endpointsFiles = ["./routes/bankAccount.js"];

swaggerAutogen(outputFile, endpointsFiles);
