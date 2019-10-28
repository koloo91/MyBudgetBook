import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PagedEntity} from '../models/paged-entity.model';
import {environment} from '../../environments/environment';
import {Booking} from '../models/booking.model';

@Injectable({
  providedIn: 'root'
})
export class BookingService {

  constructor(private http: HttpClient) {
  }

  getBookings(): Observable<PagedEntity<Booking>> {
    return this.http.get<PagedEntity<Booking>>(`${environment.host}/api/bookings`)
  }

  createBooking(title: string, comment: string, date: string, amount: number, categoryId: string, accountId: string): Observable<Booking> {
    return this.http.post<Booking>(`${environment.host}/api/bookings`, {
      title,
      comment,
      date,
      amount,
      categoryId,
      accountId
    });
  }

  updateBooking(id: string, title: string, comment: string, date: string, amount: number, categoryId: string, accountId: string) {
    return this.http.put<Booking>(`${environment.host}/api/bookings/${id}`, {
      title,
      comment,
      date,
      amount,
      categoryId,
      accountId
    });
  }
}
