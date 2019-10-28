var faker = require('faker');

module.exports = {
  path: '/api/bookings',
  method: 'POST',
  template: {
    id: faker.random.uuid(),
    title: faker.lorem.word(),
    comment: '',
    amount: 12,
    date: '2019-10-23T23:17:12.0Z',
    categoryId: faker.random.uuid(),
    accountId: faker.random.uuid(),
    created: faker.date.recent(),
    updated: faker.date.recent()
  }
};
