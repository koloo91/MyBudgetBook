var faker = require('faker');

module.exports = {
  path: '/mbb/api/accounts',
  method: 'POST',
  template: {
    id: 1,
    name: faker.lorem.word(),
    startingBalance: 12.01,
    created: faker.date.recent(),
    updated: faker.date.recent()
  }
};
