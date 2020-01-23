var faker = require('faker');

module.exports = {
  path: '/mbb/api/accounts',
  template: {
    content: function (params) {
      return [
        {
          id: '1',
          name: faker.lorem.word(),
          startingBalance: 12,
          isMain: true,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          isMain: false,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          startingBalance: 12.01,
          created: faker.date.recent(),
          updated: faker.date.recent()
        }
      ]
    },
  }
};
