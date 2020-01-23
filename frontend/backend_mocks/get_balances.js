var faker = require('faker');

module.exports = {
  path: '/mbb/api/balances',
  template: {
    content: function (params) {
      return [
        {
          accountId: '1',
          name: 'Sparkasse',
          balance: 5432
        },
        {
          accountId: '2',
          name: 'Konto',
          balance: 31000.23
        }
      ]
    },
  }
};
