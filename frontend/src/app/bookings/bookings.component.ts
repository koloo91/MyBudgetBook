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
  categoriesObservable: Observable<Category[]>;
  categories?: Category[];

  startDate: Date = new Date();
  endDate: Date = new Date();

  constructor(private bookingService: BookingService,
              private categoryService: CategoryService,
              public dialog: MatDialog) {
  }

  ngOnInit() {
    const now = new Date();
    this.startDate = new Date(now.getFullYear(), now.getMonth(), 1, 0, 0, 0);
    this.endDate = new Date(now.getFullYear(), now.getMonth() + 1, 0, 23, 59, 59);

    this.loadBookings();
    this.loadCategories();
  }

  private loadBookings() {
    this.bookings = this.bookingService.getBookings(this.startDate, this.endDate)
      .pipe(
        map(pagedEntity => pagedEntity.content)
      );
  }

  private loadCategories() {
    this.categoriesObservable = this.categoryService.getCategories()
      .pipe(
        map(pagedEntity => pagedEntity.content)
      );

    this.categoriesObservable.subscribe(categories => this.categories = categories);
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

  getCategory(categoryId: string): string {
    return this.categories.find(_ => _.id === categoryId).name || 'Unbekannt';
  }

  updateBooking(selectedBooking: Booking) {
    const dialogRef = this.dialog.open(CreateBookingDialogComponent, {
      width: '600px',
      data: selectedBooking
    });

    dialogRef.afterClosed().subscribe(result => {
      this.loadBookings();
    });
  }

  onStartDateChange() {
    this.loadBookings();
  }

  onEndDateChange() {
    this.endDate.setHours(23, 59, 59);
    this.loadBookings();
  }

  selectPreviousMonth() {
    this.endDate = new Date(this.startDate.getFullYear(), this.startDate.getMonth(), 0, 23, 59, 59);
    this.startDate = new Date(this.startDate.getFullYear(), this.startDate.getMonth() - 1, 1, 0, 0, 0)
    this.loadBookings();
  }

  selectNextMonth() {
    this.startDate = new Date(this.endDate.getFullYear(), this.endDate.getMonth() + 1, 1, 0, 0, 0)
    this.endDate = new Date(this.startDate.getFullYear(), this.startDate.getMonth() + 1, 0, 23, 59, 59);
    this.loadBookings();
  }
}
