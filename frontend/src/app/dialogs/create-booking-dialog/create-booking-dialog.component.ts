import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {AccountService} from '../../services/account.service';
import {CategoryService} from '../../services/category.service';
import {BookingService} from '../../services/booking.service';
import {forkJoin} from 'rxjs';
import {Category} from '../../models/category.model';
import {Account} from '../../models/account.model';
import {ErrorService} from '../../services/error.service';
import {ErrorVo} from '../../models/error.model';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {Booking} from '../../models/booking.model';

interface BookingDialogData {
  booking: Booking;
  updateAll: boolean;
}

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-booking-dialog.component.html',
  styleUrls: ['./create-booking-dialog.component.scss']
})
export class CreateBookingDialogComponent implements OnInit {

  isLoading = true;

  categories: Category[];
  accounts: Account[];

  bookingFormGroup: FormGroup;

  constructor(public dialogRef: MatDialogRef<CreateBookingDialogComponent>,
              private accountService: AccountService,
              private categoryService: CategoryService,
              private bookingService: BookingService,
              private errorService: ErrorService,
              @Inject(MAT_DIALOG_DATA) public data: BookingDialogData) {

    const booking = data.booking;

    this.bookingFormGroup = new FormGroup({
      title: new FormControl(booking.title, Validators.required),
      date: new FormControl(new Date(booking.date), Validators.required),
      amount: new FormControl(booking.amount, Validators.required),
      categoryId: new FormControl(booking.categoryId, Validators.required),
      accountId: new FormControl(booking.accountId, Validators.required),
      isStandingOrder: new FormControl(false),
      standingOrderPeriod: new FormControl(booking.standingOrderPeriod)
    });
  }

  ngOnInit() {
    const categories$ = this.categoryService.getCategories();
    const accounts$ = this.accountService.getAccounts();

    forkJoin([categories$, accounts$])
      .subscribe(([categories, accounts]) => {
        this.categories = categories;
        this.accounts = accounts;
        this.isLoading = false;
      }, (error: ErrorVo) => {
        this.errorService.showErrorMessage(error.message);
        this.isLoading = false;
      });
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  onOkClick() {
    this.isLoading = true;

    if (this.data.booking && this.data.booking.id.length > 0) {
      this.updateBooking();
    } else {
      this.createBooking();
    }
  }

  updateBooking() {
    this.bookingFormGroup.patchValue({
      date: this.addTimeToSelectedDate(this.bookingFormGroup.controls['date'].value).toISOString()
    });

    console.log(this.bookingFormGroup.value);
    this.bookingService.updateBooking(
      this.data.booking.id,
      this.bookingFormGroup.value,
      this.data.updateAll)
      .subscribe(booking => {
        this.dialogRef.close({success: true});
      }, (err: ErrorVo) => {
        this.errorService.showErrorMessage(err.message);
      });
  }

  createBooking() {
    const isStandingOrder = this.bookingFormGroup.controls['isStandingOrder'].value as boolean;

    this.bookingFormGroup.patchValue({
      date: this.addTimeToSelectedDate(this.bookingFormGroup.controls['date'].value).toISOString(),
      standingOrderPeriod: !isStandingOrder ? null : this.bookingFormGroup.controls['standingOrderPeriod'].value
    });

    console.log(this.bookingFormGroup.value);
    this.bookingService.createBooking(
      this.bookingFormGroup.value)
      .subscribe(booking => {
        this.dialogRef.close({success: true});
      }, (err: ErrorVo) => {
        this.errorService.showErrorMessage(err.message);
      });
  }

  addTimeToSelectedDate(date: Date): Date {
    let now = new Date();
    now.setFullYear(date.getFullYear(), date.getMonth(), date.getDate());
    now.setHours(12, 0, 0, 0);
    return now;
  }

  onDeleteClick() {
    if (!this.data.booking || this.data.booking.id.length === 0) {
      return;
    }
    this.isLoading = true;
    this.bookingService.delete(this.data.booking.id, this.data.updateAll)
      .subscribe(() => {
        this.dialogRef.close({success: true});
      }, (err: ErrorVo) => {
        this.isLoading = false;
        this.errorService.showErrorMessage(err.message);
      });
  }
}
