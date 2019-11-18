import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {AccountService} from '../../services/account.service';
import {CategoryService} from '../../services/category.service';
import {BookingService} from '../../services/booking.service';
import {Observable} from 'rxjs';
import {Category} from '../../models/category.model';
import {Account} from '../../models/account.model';
import {ErrorService} from '../../services/error.service';

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-booking-dialog.component.html',
  styleUrls: ['./create-booking-dialog.component.scss']
})
export class CreateBookingDialogComponent implements OnInit {

  title: string;
  date: Date = new Date();
  amount: number;
  categoryId: string;
  accountId: string;
  isLoading = false;

  categories: Observable<Category[]>;
  accounts: Observable<Account[]>;

  isStandingOrder: boolean = false;
  standingOrderPeriod: string = 'MONTHLY';

  constructor(public dialogRef: MatDialogRef<CreateBookingDialogComponent>,
              private accountService: AccountService,
              private categoryService: CategoryService,
              private bookingService: BookingService,
              private errorService: ErrorService,
              @Inject(MAT_DIALOG_DATA) public data?: any) {

    if (data && data.booking) {
      const booking = data.booking;
      this.title = booking.title;
      this.date = new Date(booking.date);
      this.amount = booking.amount;
      this.categoryId = booking.categoryId;
      this.accountId = booking.accountId;
    }
  }

  ngOnInit() {
    this.categories = this.categoryService.getCategories();
    this.accounts = this.accountService.getAccounts();
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  onOkClick() {
    this.isLoading = true;

    if (this.data && this.data.booking) {
      this.updateBooking();
    } else {
      this.createBooking();
    }
  }

  updateBooking() {
    this.bookingService.updateBooking(
      this.data.booking.id,
      this.title,
      this.addTimeToSelectedDate(this.date).toISOString(),
      this.amount,
      this.categoryId,
      this.accountId,
      this.data.updateAll)
      .subscribe(booking => {
        this.dialogRef.close({success: true});
      }, err => {
        this.errorService.showErrorMessage(err.error);
      });
  }

  createBooking() {
    this.bookingService.createBooking(
      this.title,
      this.addTimeToSelectedDate(this.date).toISOString(),
      this.amount,
      this.categoryId,
      this.accountId,
      this.isStandingOrder ? this.standingOrderPeriod : null)
      .subscribe(booking => {
        console.log(booking);
        this.dialogRef.close({success: true});
      }, err => {
        this.errorService.showErrorMessage(err.error);
      });
  }

  addTimeToSelectedDate(date: Date): Date {
    let now = new Date();
    now.setFullYear(date.getFullYear(), date.getMonth(), date.getDate());
    return now;
  }

  onDeleteClick() {
    if (!this.data || !this.data.booking) {
      return;
    }
    this.isLoading = true;
    this.bookingService.delete(this.data.booking.id, this.data.updateAll)
      .subscribe(() => {
        this.dialogRef.close({success: true});
      }, (err: any) => {
        this.isLoading = false;
        this.errorService.showErrorMessage(err.message);
      });
  }
}
