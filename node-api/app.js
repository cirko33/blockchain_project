const express = require("express");
const logger = require("morgan");
const cors = require("cors");
const swaggerUi = require("swagger-ui-express");
const swaggerFile = require("./docs/swagger_output.json");

const app = express();

const formatsLogger = app.get("env") === "development" ? "dev" : "short";

app.use(logger(formatsLogger));
app.use(cors());
app.use(express.json());

const bankAccountRoutes = require("./routes/main");

app.use(bankAccountRoutes);
app.use("/doc", swaggerUi.serve, swaggerUi.setup(swaggerFile));

app.use((req, res) => {
  res.status(404).json({ message: "Not found" });
});

app.use((err, req, res, next) => {
  res.status(500).json({ message: err.message });
});

// module.exports = app;
app.listen(3000, () => {
  console.log("Server running. Use our API on port: 3000");
});
