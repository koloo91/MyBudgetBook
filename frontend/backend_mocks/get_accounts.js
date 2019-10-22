var faker = require('faker');

module.exports = {
  path: '/api/accounts',
  template: {
    content: function (params) {
      return [
        {
          id: faker.random.uuid(),
          name: faker.lorem.word(),
          startingBalance: 12,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: faker.random.uuid(),
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        }
      ]
    },
  }
};
