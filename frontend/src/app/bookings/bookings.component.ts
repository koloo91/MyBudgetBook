import {Component, OnInit} from '@angular/core';
import {BookingService} from '../services/booking.service';
import {forkJoin} from 'rxjs';
import {Booking} from '../models/booking.model';
import {MatDialog} from '@angular/material/dialog';
import {CreateBookingDialogComponent} from '../dialogs/create-booking-dialog/create-booking-dialog.component';
import {Category} from '../models/category.model';
import {CategoryService} from '../services/category.service';
import {Balance} from '../models/balance.model';
import {Account} from '../models/account.model';
import {BalanceService} from '../services/balance.service';
import {UpdateBookingDialogComponent} from '../dialogs/update-booking-dialog/update-booking-dialog.component';
import {ErrorService} from '../services/error.service';
import {ErrorVo} from '../models/error.model';
import {AccountService} from '../services/account.service';

@Component({
  selector: 'app-bookings',
  templateUrl: './bookings.component.html',
  styleUrls: ['./bookings.component.scss']
})
export class BookingsComponent implements OnInit {
  isLoading: boolean = true;

  bookings: Booking[] = [];
  categories: Category[] = [];
  accounts: Account[] = [];
  balances: Balance[] = [];

  startDate: Date = new Date();
  endDate: Date = new Date();

  constructor(private bookingService: BookingService,
              private categoryService: CategoryService,
              private accountService: AccountService,
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
    const accounts$ = this.accountService.getAccounts();
    const balances$ = this.balanceService.getBalances();

    forkJoin([bookings$, categories$, accounts$, balances$])
      .subscribe(([bookings, categories, accounts, balances]) => {
        this.isLoading = false;
        this.bookings = bookings;
        this.categories = categories;
        this.accounts = accounts;
        this.balances = balances;
      }, (err: ErrorVo) => {
        this.isLoading = false;
        this.errorService.showErrorMessage(err.message);
      });
  }

  showCreateDialog() {
    const dialogRef = this.dialog.open(CreateBookingDialogComponent, {
      width: '600px',
      data: {booking: new Booking(), updateAll: false}
    });

    dialogRef.afterClosed().subscribe(result => {
      if (!result || !result.success) {
        return;
      }
      this.loadData();
    });
  }

  getCategory(categoryId: string): string {
    const maybeCategory = this.categories.find(_ => _.id === categoryId);
    if (!maybeCategory) {
      return 'Unbekannt';
    }
    return maybeCategory.name;
  }

  updateBooking(selectedBooking: Booking) {
    if (selectedBooking.standingOrderId) {
      const dialogRef = this.dialog.open(UpdateBookingDialogComponent, {
        width: '600px'
      });

      dialogRef.afterClosed().subscribe(updateAll => {
        if (!updateAll) {
          return;
        }
        this.displayUpdateBookingDialog(selectedBooking, updateAll.updateAll);
      });
    } else {
      this.displayUpdateBookingDialog(selectedBooking, false);
    }
  }

  displayUpdateBookingDialog(booking: Booking, updateAll: boolean) {
    const dialogRef = this.dialog.open(CreateBookingDialogComponent, {
      width: '600px',
      data: {booking: booking, updateAll: updateAll}
    });

    dialogRef.afterClosed().subscribe(result => {
      if (!result || !result.success) {
        return;
      }
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

  dateIsInFuture(date: string): boolean {
    const endOfDay = new Date();
    endOfDay.setHours(23, 59, 59);
    return new Date(date).getTime() > endOfDay.getTime();
  }

  getMainAccountName(): string {
    const mainAccount = this.accounts.find(_ => _.isMain)
    if (!mainAccount) {
      return 'Kein Hauptkonto ausgewählt';
    }

    return mainAccount.name;
  }

  getMainAccountBalance(): number {
    const mainAccount = this.accounts.find(_ => _.isMain)
    if (!mainAccount) {
      return 0.0;
    }

    return this.balances.find(_ => _.accountId === mainAccount.id).balance;
  }
}
