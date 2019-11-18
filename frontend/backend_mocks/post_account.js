var faker = require('faker');

module.exports = {
  path: '/api/accounts',
  method: 'POST',
  status: function (req, res, next) {
    res.send(400, {message: 'Nope'});
  },
  template: {
    id: 1,
    name: faker.lorem.word(),
    startingBalance: 12.01,
    created: faker.date.recent(),
    updated: faker.date.recent()
  }
};
