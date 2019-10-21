import {Component, OnInit} from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';
import {AccountService} from '../../services/account.service';

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-account-dialog.component.html',
  styleUrls: ['./create-account-dialog.component.scss']
})
export class CreateAccountDialogComponent implements OnInit {

  accountName: string;

  constructor(public dialogRef: MatDialogRef<CreateAccountDialogComponent>,
              private accountService: AccountService) {
  }

  ngOnInit() {
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  createAccount() {
    console.log(this.accountName);
    this.accountService.createAccount(this.accountName).subscribe(account => {
      console.log(account);
      this.dialogRef.close({success: true});
    }, err => {
      console.log(err);
    });
  }
}
