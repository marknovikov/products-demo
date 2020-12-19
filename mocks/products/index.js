"use strict";

const app = require("express")();
const Controller = require("./controller");
const newRouter = require("./router");

app.use("/api/products", newRouter(new Controller()));

const srv = app.listen(3000, () => {
  console.log("products server has started");
});

const shutdown = (err) => {
  console.error(err);

  srv.close((err) => {
    if (err) {
      console.error(err);
      process.exit(1);
    }
  });
};

srv.on("error", (err) => shutdown(err));
process.on("SIGTERM", (err) => shutdown(err));
