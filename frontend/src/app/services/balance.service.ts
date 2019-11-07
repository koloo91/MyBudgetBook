import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PagedEntity} from '../models/paged-entity.model';
import {environment} from '../../environments/environment';
import {Balance} from '../models/balance.model';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BalanceService {

  constructor(private http: HttpClient) {
  }

  getBalances(): Observable<Balance[]> {
    return this.http.get<PagedEntity<Balance>>(`${environment.host}/api/balances`)
      .pipe(
        map(_ => _.content)
      );
  }
}
