var faker = require('faker');

module.exports = {
  path: '/api/categories/:id',
  method: 'PUT',
  template: {
    id: faker.random.uuid(),
    name: faker.lorem.word(),
    created: faker.date.recent(),
    updated: faker.date.recent()
  }
};
