import {Component, OnInit} from '@angular/core';
import {map} from 'rxjs/operators';
import {BookingService} from '../services/booking.service';
import {Observable} from 'rxjs';
import {Booking} from '../models/booking.model';

@Component({
  selector: 'app-bookings',
  templateUrl: './bookings.component.html',
  styleUrls: ['./bookings.component.scss']
})
export class BookingsComponent implements OnInit {

  bookings: Observable<Booking[]>;

  constructor(private bookingService: BookingService) {
  }

  ngOnInit() {
    this.loadBookings();
  }

  private loadBookings() {
    this.bookings = this.bookingService.getBookings()
      .pipe(
        map(pagedEntity => pagedEntity.content)
      );
  }

  showCreateDialog() {

  }
}
