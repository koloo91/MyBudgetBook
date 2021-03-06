var faker = require('faker');

module.exports = {
  path: '/mbb/api/bookings/:id',
  method: 'PUT',
  template: {
    id: faker.random.uuid(),
    title: faker.lorem.word(),
    amount: 12,
    date: '2019-10-23T23:17:12.0Z',
    categoryId: faker.random.uuid(),
    accountId: faker.random.uuid(),
    created: faker.date.recent(),
    updated: faker.date.recent()
  }
};
