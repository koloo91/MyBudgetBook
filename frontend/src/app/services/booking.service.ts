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

  createBooking(title: string, date: string, amount: number, categoryId: string, accountId: string, standingOrderPeriod?: string): Observable<Booking> {
    return this.http.post<Booking>(`${environment.host}/api/bookings`, {
      title,
      date,
      amount,
      categoryId,
      accountId,
      standingOrderPeriod
    });
  }

  updateBooking(id: string, title: string, date: string, amount: number, categoryId: string, accountId: string, updateAll: boolean) {
    return this.http.put<Booking>(`${environment.host}/api/bookings/${id}`, {
      title,
      date,
      amount,
      categoryId,
      accountId
    }, {
      params: {
        updateStrategy: updateAll ? 'ALL' : 'ONE'
      }
    });
  }

  delete(id: string, updateAll: boolean) {
    return this.http.delete(`${environment.host}/api/bookings/${id}`,
      {
        params: {
          updateStrategy: updateAll ? 'ALL' : 'ONE'
        }
      })
  }
}
