var faker = require('faker');

module.exports = {
  path: '/mbb/api/statistics/month',
  template: {
    content: function (params) {
      return [
        {
          month: 1,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 2,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 3,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 4,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 5,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 6,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 7,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 8,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 9,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 10,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 11,
          expenses: -123.43,
          incomes: 12000.12
        },
        {
          month: 2,
          expenses: -123.43,
          incomes: 12000.12
        }
      ]
    }
  }
};
