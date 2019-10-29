var faker = require('faker');

module.exports = {
  path: '/api/categories',
  template: {
    content: function (params) {
      return [
        {
          id: '1',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '3',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        }
      ]
    }
  }
};
