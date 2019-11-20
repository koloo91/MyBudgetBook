import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {AccountService} from '../../services/account.service';
import {Account} from '../../models/account.model';
import {ErrorService} from '../../services/error.service';
import {ErrorVo} from '../../models/error.model';
import {FormControl, FormGroup, Validators} from '@angular/forms';

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-account-dialog.component.html',
  styleUrls: ['./create-account-dialog.component.scss']
})
export class CreateAccountDialogComponent implements OnInit {

  accountName: string;
  startingBalance: number;
  isLoading = false;

  accountFormGroup: FormGroup;

  constructor(public dialogRef: MatDialogRef<CreateAccountDialogComponent>,
              private accountService: AccountService,
              private errorService: ErrorService,
              @Inject(MAT_DIALOG_DATA) public data?: Account) {
  }

  ngOnInit() {
    this.accountFormGroup = new FormGroup({
        'name': new FormControl(this.accountName, Validators.required),
        'startingBalance': new FormControl(this.startingBalance, Validators.required)
      }
    );
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
