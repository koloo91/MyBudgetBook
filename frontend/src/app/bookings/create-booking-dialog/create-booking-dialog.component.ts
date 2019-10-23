import {Component, OnInit} from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';
import {AccountService} from '../../services/account.service';
import {CategoryService} from '../../services/category.service';
import {BookingService} from '../../services/booking.service';
import {Observable} from 'rxjs';
import {Category} from '../../models/category.model';
import {map} from 'rxjs/operators';
import {Account} from '../../models/account.model';

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-booking-dialog.component.html',
  styleUrls: ['./create-booking-dialog.component.scss']
})
export class CreateBookingDialogComponent implements OnInit {

  title: string;
  comment: string;
  date: Date;
  amount: number;
  categoryId: string;
  accountId: string;
  isLoading = false;

  categories: Observable<Category[]>;
  accounts: Observable<Account[]>;

  constructor(public dialogRef: MatDialogRef<CreateBookingDialogComponent>,
              private accountService: AccountService,
              private categoryService: CategoryService,
              private bookingService: BookingService) {
  }

  ngOnInit() {
    this.categories = this.categoryService.getCategories().pipe(map(_ => _.content));
    this.accounts = this.accountService.getAccounts().pipe(map(_ => _.content));
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  createBooking() {
    this.isLoading = true;
    this.bookingService.createBooking(this.title, this.comment, this.date.toISOString(), this.amount, this.categoryId, this.accountId)
      .subscribe(booking => {
          console.log(booking);
          this.dialogRef.close({success: true});
        }, err => console.log(err),
        () => this.isLoading = false);
  }
}
