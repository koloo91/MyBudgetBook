var faker = require('faker');

module.exports = {
  path: '/api/bookings',
  template: {
    content: function (params) {
      return [
        {
          id: faker.random.uuid(),
          title: faker.lorem.word(),
          comment: '',
          amount: 12,
          date: '2019-10-23T22:00:00.0Z',
          categoryId: faker.random.uuid(),
          accountId: faker.random.uuid(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: faker.random.uuid(),
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: faker.random.uuid(),
          accountId: faker.random.uuid(),
          created: faker.date.recent(),
          updated: faker.date.recent()
        }
      ]
    },
  }
};
