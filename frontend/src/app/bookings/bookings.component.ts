import {Component, OnInit} from '@angular/core';
import {BookingService} from '../services/booking.service';
import {forkJoin} from 'rxjs';
import {Booking} from '../models/booking.model';
import {MatDialog} from '@angular/material/dialog';
import {CreateBookingDialogComponent} from '../dialogs/create-booking-dialog/create-booking-dialog.component';
import {Category} from '../models/category.model';
import {CategoryService} from '../services/category.service';
import {Balance} from '../models/balance.model';
import {BalanceService} from '../services/balance.service';
import {UpdateBookingDialogComponent} from '../dialogs/update-booking-dialog/update-booking-dialog.component';
import {ErrorVo} from '../models/error.model';
import {ErrorService} from '../services/error.service';

@Component({
  selector: 'app-bookings',
  templateUrl: './bookings.component.html',
  styleUrls: ['./bookings.component.scss']
})
export class BookingsComponent implements OnInit {
  isLoading: boolean = true;

  bookings: Booking[] = [];
  categories: Category[] = [];
  balances: Balance[] = [];

  startDate: Date = new Date();
  endDate: Date = new Date();

  constructor(private bookingService: BookingService,
              private categoryService: CategoryService,
              private balanceService: BalanceService,
              private errorService: ErrorService,
              public dialog: MatDialog) {
  }

  ngOnInit() {
    const now = new Date();
    this.startDate = new Date(now.getFullYear(), now.getMonth(), 1, 0, 0, 0);
    this.endDate = new Date(now.getFullYear(), now.getMonth() + 1, 0, 23, 59, 59);

    this.loadData();
  }

  private loadData() {
    this.isLoading = true;

    const bookings$ = this.bookingService.getBookings(this.startDate, this.endDate);
    const categories$ = this.categoryService.getCategories();
    const balances$ = this.balanceService.getBalances();

    forkJoin(bookings$, categories$, balances$)
      .subscribe(([bookings, categories, balances]) => {
        this.isLoading = false;
        this.bookings = bookings;
        this.categories = categories;
        this.balances = balances;
      }, (err: ErrorVo) => {
        this.isLoading = false;
        this.errorService.showErrorMessage(err.message);
      });
  }

  showCreateDialog() {
    const dialogRef = this.dialog.open(CreateBookingDialogComponent, {
      width: '600px'
    });

    dialogRef.afterClosed().subscribe(result => {
      this.loadData();
    });
  }

  getCategory(categoryId: string): string {
    return this.categories.find(_ => _.id === categoryId).name || 'Unbekannt';
  }

  updateBooking(selectedBooking: Booking) {
    if (selectedBooking.standingOrderId) {
      const dialogRef = this.dialog.open(UpdateBookingDialogComponent, {
        width: '600px'
      });

      dialogRef.afterClosed().subscribe(updateAll => {
        this.displayUpdateBookingDialog(selectedBooking, updateAll);
      });
    } else {
      this.displayUpdateBookingDialog(selectedBooking, false);
    }
  }

  displayUpdateBookingDialog(booking: Booking, updateAll: boolean) {
    const dialogRef = this.dialog.open(CreateBookingDialogComponent, {
      width: '600px',
      data: {booking, updateAll}
    });

    dialogRef.afterClosed().subscribe(result => {
      this.loadData();
    });
  }

  onStartDateChange() {
    this.loadData();
  }

  onEndDateChange() {
    this.endDate.setHours(23, 59, 59);
    this.loadData();
  }

  selectPreviousMonth() {
    this.endDate = new Date(this.startDate.getFullYear(), this.startDate.getMonth(), 0, 23, 59, 59);
    this.startDate = new Date(this.startDate.getFullYear(), this.startDate.getMonth() - 1, 1, 0, 0, 0)
    this.loadData();
  }

  selectNextMonth() {
    this.startDate = new Date(this.endDate.getFullYear(), this.endDate.getMonth() + 1, 1, 0, 0, 0)
    this.endDate = new Date(this.startDate.getFullYear(), this.startDate.getMonth() + 1, 0, 23, 59, 59);
    this.loadData();
  }

  balancesClicked() {
    console.log('balances');
  }
}
