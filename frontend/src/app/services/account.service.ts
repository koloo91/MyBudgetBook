import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PagedEntity} from '../models/paged-entity.model';
import {environment} from '../../environments/environment';
import {Account} from '../models/account.model';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(private http: HttpClient) {
  }

  getAccounts(): Observable<Account[]> {
    return this.http.get<PagedEntity<Account>>(`${environment.host}/api/accounts`)
      .pipe(
        map(_ => _.content)
      );
  }

  createAccount(name: string, startingBalance: number): Observable<Account> {
    return this.http.post<Account>(`${environment.host}/api/accounts`, {name, startingBalance});
  }
}
