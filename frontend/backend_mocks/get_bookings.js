var faker = require('faker');

module.exports = {
  path: '/api/bookings',
  template: {
    content: function (params) {
      return [
        {
          id: '1',
          title: faker.lorem.word(),
          comment: '',
          amount: 12,
          date: '2019-10-23T22:00:00.0Z',
          categoryId: '1',
          accountId: '1',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        },
        {
          id: '2',
          title: faker.lorem.word(),
          comment: '',
          amount: -50.00,
          date: '2019-10-23T23:17:12.0Z',
          categoryId: '2',
          accountId: '2',
          created: faker.date.recent(),
          updated: faker.date.recent()
        }
      ]
    },
  }
};
