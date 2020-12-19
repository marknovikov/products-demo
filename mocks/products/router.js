"use strict";

const express = require("express-promise-router");
const path = require("path");
const CsvView = require("./csv");

const newRouter = (controller) => {
  const router = express();

  router.get("/*.csv", async (req, resp) => {
    const rows = await controller.get(req.url);
    const view = await new CsvView(rows).render();

    const filename = path.basename(req.url);

    resp.writeHead(200, {
      "content-type": "text/csv; charset=utf-8",
      "content-disposition": `attachment; filename=${filename}`,
      "cache-control": "no-cache",
      connection: "keep-alive",
    });
    resp.end(view);
  });

  return router;
};

module.exports = newRouter;
