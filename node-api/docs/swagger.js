const swaggerAutogen = require("swagger-autogen")();

const outputFile = "./swagger_output.json";
const endpointsFiles = ["./routes/main.js"];

swaggerAutogen(outputFile, endpointsFiles);
