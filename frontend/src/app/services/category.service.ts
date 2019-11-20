import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PagedEntity} from '../models/paged-entity.model';
import {environment} from '../../environments/environment';
import {Category} from '../models/category.model';
import {catchError, map} from 'rxjs/operators';
import {BaseService} from './base.service';

@Injectable({
  providedIn: 'root'
})
export class CategoryService extends BaseService {

  constructor(private http: HttpClient) {
    super()
  }

  getCategories(): Observable<Category[]> {
    return this.http.get<PagedEntity<Category>>(`${environment.host}/api/categories`)
      .pipe(
        map(_ => _.content),
        catchError(this.handleError)
      );
  }

  createCategory(name: string): Observable<Category> {
    return this.http.post<Category>(`${environment.host}/api/categories`, {name})
      .pipe(
        catchError(this.handleError)
      );
  }

  updateCategory(id: string, name: string): Observable<Category> {
    return this.http.put<Category>(`${environment.host}/api/categories/${id}`, {name})
      .pipe(
        catchError(this.handleError)
      );
  }

}
