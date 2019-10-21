var faker = require('faker');

module.exports = {
  path: '/api/accounts',
  method: 'POST',
  template: {
    id: faker.random.uuid(),
    name: faker.lorem.word(),
    created: faker.date.recent(),
    updated: faker.date.recent()
  }
};
