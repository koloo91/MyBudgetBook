import {Component, OnInit} from '@angular/core';
import {AccountService} from '../services/account.service';
import {Observable} from 'rxjs';
import {map} from 'rxjs/operators';
import {MatDialog} from '@angular/material/dialog';
import {CreateAccountDialogComponent} from '../dialogs/create-account-dialog/create-account-dialog.component';
import {Account} from '../models/account.model';

@Component({
  selector: 'app-accounts',
  templateUrl: './accounts.component.html',
  styleUrls: ['./accounts.component.scss'],
  host: {
    'class': 'router-flex'
  }
})
export class AccountsComponent implements OnInit {

  accounts: Observable<Account[]>;

  constructor(private accountService: AccountService,
              public dialog: MatDialog) {
  }

  ngOnInit() {
    this.loadAccounts();
  }

  showCreateDialog() {
    const dialogRef = this.dialog.open(CreateAccountDialogComponent, {
      width: '600px'
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
      this.loadAccounts();
    });
  }

  private loadAccounts() {
    this.accounts = this.accountService.getAccounts()
      .pipe(
        map(pagedEntity => pagedEntity.content)
      );
  }
}
