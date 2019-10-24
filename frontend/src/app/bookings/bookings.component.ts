import {Component, OnInit} from '@angular/core';
import {map} from 'rxjs/operators';
import {BookingService} from '../services/booking.service';
import {Observable} from 'rxjs';
import {Booking} from '../models/booking.model';
import {MatDialog} from '@angular/material/dialog';
import {CreateBookingDialogComponent} from './create-booking-dialog/create-booking-dialog.component';
import {Category} from '../models/category.model';
import {CategoryService} from '../services/category.service';

@Component({
  selector: 'app-bookings',
  templateUrl: './bookings.component.html',
  styleUrls: ['./bookings.component.scss']
})
export class BookingsComponent implements OnInit {

  bookings: Observable<Booking[]>;
  categories: Observable<Category[]>;

  constructor(private bookingService: BookingService,
              private categoryService: CategoryService,
              public dialog: MatDialog) {
  }

  ngOnInit() {
    this.loadBookings();
    this.loadCategories();
  }

  private loadBookings() {
    this.bookings = this.bookingService.getBookings()
      .pipe(
        map(pagedEntity => pagedEntity.content)
      );
  }

  private loadCategories() {
    this.categories = this.categoryService.getCategories()
      .pipe(
        map(pagedEntity => pagedEntity.content)
      );
  }

  showCreateDialog() {
    const dialogRef = this.dialog.open(CreateBookingDialogComponent, {
      width: '600px'
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
      this.loadBookings();
    });
  }

  getCategory(categoryId: string): Observable<string> {
    return this.categories.pipe(
      map(categories => categories.find(_ => _.id === categoryId)),
      map(_ => _.id)
    );
  }
}
