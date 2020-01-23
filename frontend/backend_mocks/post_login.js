var faker = require('faker');

module.exports = {
  path: '/lgn/api/login',
  method: 'POST',
  template: {
    accessToken: 'accessToken',
    refreshToken: 'refreshToken',
    type: 'Bearer'
  }
};
