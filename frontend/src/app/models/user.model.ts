export class User {
  name: string;
  accessToken: string;
  refreshToken: string;

  constructor(name: string, accessToken: string, refreshToken: string) {
    this.name = name;
    this.accessToken = accessToken;
    this.refreshToken = refreshToken;
  }
}
