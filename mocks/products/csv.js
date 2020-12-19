"use strict";

const fields = ["name", "price"];
const delimeter = ";";
const eol = "\n";
const nil = "null";

const formatHead = () => {
  return fields.map((field) => `"${field}"`).join(delimeter);
};

const formatRow = (row) => {
  return fields
    .map((field) => {
      let col = row[field];

      if (!col) {
        col = nil;
      }

      if (+col instanceof Number) {
        col = (+col).toFixed(2);
      }

      return `"${col}"`;
    })
    .join(delimeter);
};

class CsvView {
  constructor(rows) {
    this._rows = rows;
  }

  async render() {
    const out = [];
    out.push(formatHead());

    return this._rows
      .reduce((agg, row) => {
        agg.push(formatRow(row));
        return agg;
      }, out)
      .join(eol);
  }
}

module.exports = CsvView;
