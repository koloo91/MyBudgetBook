var faker = require('faker');

module.exports = {
  path: '/api/balances',
  template: {
    content: function (params) {
      return [
        {
          accountId: '1',
          balance: 5432
        },
        {
          accountId: '2',
          balance: 31000.23
        }
      ]
    },
  }
};
