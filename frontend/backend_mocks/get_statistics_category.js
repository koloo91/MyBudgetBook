var faker = require('faker');

module.exports = {
  path: '/api/statistics/category',
  template: {
    content: function (params) {
      return [
        {
          name: faker.name.findName(),
          sum: Math.abs(faker.finance.amount())
        },
        {
          name: faker.name.findName(),
          sum: Math.abs(faker.finance.amount())
        },
        {
          name: faker.name.findName(),
          sum: Math.abs(faker.finance.amount())
        },
        {
          name: faker.name.findName(),
          sum: Math.abs(faker.finance.amount())
        },
        {
          name: faker.name.findName(),
          sum: Math.abs(faker.finance.amount())
        },
        {
          name: faker.name.findName(),
          sum: Math.abs(faker.finance.amount())
        }
      ]
    }
  }
};
