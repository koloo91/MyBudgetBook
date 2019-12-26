import {Injectable} from '@angular/core';
import {BaseService} from './base.service';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PagedEntity} from '../models/paged-entity.model';
import {environment} from '../../environments/environment';
import {catchError, map} from 'rxjs/operators';
import {MonthStatistic} from '../models/month_statistic.model';
import {CategoryStatistic} from '../models/category_statistic.model';

@Injectable({
  providedIn: 'root'
})
export class StatisticService extends BaseService {

  constructor(private http: HttpClient) {
    super();
  }

  getMonthStatistics(year: number = new Date().getFullYear()): Observable<MonthStatistic[]> {
    return this.http.get<PagedEntity<MonthStatistic>>(`${environment.host}/api/statistics/month`, {
      params: {
        year: `${year}`
      }
    }).pipe(
      map(_ => _.content),
      catchError(this.handleError)
    );
  }

  getCategoryStatistics(year: number = new Date().getFullYear()): Observable<CategoryStatistic[]> {
    return this.http.get<PagedEntity<CategoryStatistic>>(`${environment.host}/api/statistics/category`, {
      params: {
        year: `${year}`
      }
    }).pipe(
      map(_ => _.content),
      catchError(this.handleError)
    );
  }
}
