var faker = require('faker');

module.exports = {
  path: '/api/categories',
  template: {
    content: function (params) {
      return [
        {
          id: 'e13fa5e2-5640-48b8-b40e-d96acaa969d5',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2cb2572b-3bb9-4912-8167-672de9a5c7a5',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '6fb7b689-59b1-4255-9c67-9dcfe5cb9858',
          name: faker.lorem.word(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        }
      ]
    }
  }
};
