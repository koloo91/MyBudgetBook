import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PagedEntity} from '../models/paged-entity.model';
import {environment} from '../../environments/environment';
import {Balance} from '../models/balance.model';
import {catchError, map} from 'rxjs/operators';
import {BaseService} from './base.service';

@Injectable({
  providedIn: 'root'
})
export class BalanceService extends BaseService {

  constructor(private http: HttpClient) {
    super()
  }

  getBalances(): Observable<Balance[]> {
    return this.http.get<PagedEntity<Balance>>(`${environment.host}/mbb/api/balances`)
      .pipe(
        map(_ => _.content),
        catchError(this.handleError)
      );
  }
}
