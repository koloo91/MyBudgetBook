import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PagedEntity} from '../models/paged-entity.model';
import {environment} from '../../environments/environment';
import {Account} from '../models/account.model';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(private http: HttpClient) {
  }

  getAccounts(): Observable<PagedEntity<Account>> {
    return this.http.get<PagedEntity<Account>>(`${environment.host}/api/accounts`)
  }

  createAccount(name: string, startingBalance: number): Observable<Account> {
    return this.http.post<Account>(`${environment.host}/api/accounts`, {name, startingBalance});
  }
}
