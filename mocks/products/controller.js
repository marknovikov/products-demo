"use strict";

const randomPrice = () => {
  return Math.round(Math.random() * 1000 * 100) / 100;
};

const rows = [
  {
    name: "t-shirt",
    price: randomPrice(),
  },
  {
    name: "phone",
    price: randomPrice(),
  },
  {
    name: "football",
    price: randomPrice(),
  },
  {
    name: "black sneakers",
    price: randomPrice(),
  },
  {
    name: "white sneakers",
    price: randomPrice(),
  },
  {
    name: "cup",
    price: randomPrice(),
  },
  {
    name: "mug",
    price: randomPrice(),
  },
  {
    name: "galsses",
    price: randomPrice(),
  },
  {
    name: "keyabord",
    price: randomPrice(),
  },
  {
    name: "robot assasin",
    price: randomPrice(),
  },
  {
    name: "helicopter",
    price: randomPrice(),
  },
  {
    name: "yacht",
    price: randomPrice(),
  },
  {
    name: "golfball",
    price: randomPrice(),
  },
  {
    name: "book",
    price: randomPrice(),
  },
  {
    name: "notebook",
    price: randomPrice(),
  },
  {
    name: "speaker",
    price: randomPrice(),
  },
  {
    name: "headphones",
    price: randomPrice(),
  },
  {
    name: "plate",
    price: randomPrice(),
  },
  {
    name: "table",
    price: randomPrice(),
  },
  {
    name: "chair",
    price: randomPrice(),
  },
  {
    name: "cd",
    price: randomPrice(),
  },
  {
    name: "dvd",
    price: randomPrice(),
  },
  {
    name: "blue-ray",
    price: randomPrice(),
  },
];

class Controller {
  async get(_) {
    const copy = rows.slice();

    const changePriceCount = 7;
    for (let i = 0; i < changePriceCount; i++) {
      const idx = Math.floor(Math.random() * copy.length);
      const item = copy[idx];

      item.price = randomPrice();
    }

    return copy;
  }
}

module.exports = Controller;
