import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {AccountService} from '../../services/account.service';
import {Account} from '../../models/account.model';
import {ErrorService} from '../../services/error.service';
import {ErrorVo} from '../../models/error.model';

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
              private errorService: ErrorService,
              @Inject(MAT_DIALOG_DATA) public data?: Account) {
  }

  ngOnInit() {
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  createAccount() {
    this.isLoading = true;
    this.accountService.createAccount(this.accountName, this.startingBalance).subscribe(account => {
      this.dialogRef.close({success: true});
    }, (err: ErrorVo) => {
      this.isLoading = false;
      this.errorService.showErrorMessage(err.message);
    });
  }
}
