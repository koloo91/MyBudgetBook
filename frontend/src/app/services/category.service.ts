import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PagedEntity} from '../models/paged-entity.model';
import {environment} from '../../environments/environment';
import {Category} from '../models/category.model';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {

  constructor(private http: HttpClient) {
  }

  getCategories(): Observable<PagedEntity<Category>> {
    return this.http.get<PagedEntity<Category>>(`${environment.host}/api/categories`)
  }

  createCategory(name: string): Observable<Category> {
    return this.http.post<Category>(`${environment.host}/api/categories`, {name});
  }

  updateCategory(id: string, name: string): Observable<Category> {
    return this.http.put<Category>(`${environment.host}/api/categories/${id}`, {name});
  }

}
