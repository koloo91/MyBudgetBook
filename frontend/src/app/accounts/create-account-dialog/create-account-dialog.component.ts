import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {AccountService} from '../../services/account.service';
import {Account} from '../../models/account.model';

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-account-dialog.component.html',
  styleUrls: ['./create-account-dialog.component.scss']
})
export class CreateAccountDialogComponent implements OnInit {

  accountName: string;
  startingBalance: number;
  isLoading = false;

  constructor(public dialogRef: MatDialogRef<CreateAccountDialogComponent>,
              private accountService: AccountService,
              @Inject(MAT_DIALOG_DATA) public data?: Account) {
    console.log(data);
  }

  ngOnInit() {
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  createAccount() {
    console.log(this.accountName);
    this.isLoading = true;
    this.accountService.createAccount(this.accountName, this.startingBalance).subscribe(account => {
        console.log(account);
        this.dialogRef.close({success: true});
      }, err => console.log(err),
      () => this.isLoading = false);
  }
}
