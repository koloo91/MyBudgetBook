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

  getBookings(startDate: Date, endDate: Date): Observable<PagedEntity<Booking>> {
    let params = {
      startDate: startDate.toISOString(),
      endDate: endDate.toISOString()
    };
    return this.http.get<PagedEntity<Booking>>(`${environment.host}/api/bookings`, {params: params})
  }

  createBooking(title: string, comment: string, date: string, amount: number, categoryId: string, accountId: string, standingOrderPeriod?: string): Observable<Booking> {
    return this.http.post<Booking>(`${environment.host}/api/bookings`, {
      title,
      comment,
      date,
      amount,
      categoryId,
      accountId,
      standingOrderPeriod
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
