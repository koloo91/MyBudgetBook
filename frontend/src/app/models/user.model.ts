export class User {
  name: string;
  token?: string;


  constructor(name: string, token: string) {
    this.name = name;
    this.token = token;
  }
}
