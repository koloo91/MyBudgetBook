import {Component, OnInit} from '@angular/core';
import {AccountService} from '../services/account.service';
import {forkJoin, Observable} from 'rxjs';
import {MatDialog} from '@angular/material/dialog';
import {CreateAccountDialogComponent} from '../dialogs/create-account-dialog/create-account-dialog.component';
import {Account} from '../models/account.model';
import {BalanceService} from '../services/balance.service';
import {Balance} from '../models/balance.model';

@Component({
  selector: 'app-accounts',
  templateUrl: './accounts.component.html',
  styleUrls: ['./accounts.component.scss'],
  host: {
    'class': 'router-flex'
  }
})
export class AccountsComponent implements OnInit {

  isLoading: boolean = true;
  accounts$: Observable<Account[]>;
  balances$: Observable<Balance[]>;
  balances: Balance[] = [];

  constructor(private accountService: AccountService,
              private balanceService: BalanceService,
              public dialog: MatDialog) {
  }

  ngOnInit() {
    this.loadData();
  }

  showCreateDialog() {
    const dialogRef = this.dialog.open(CreateAccountDialogComponent, {
      width: '600px'
    });

    dialogRef.afterClosed().subscribe(result => {
      this.loadData();
    });
  }

  private loadData() {
    this.isLoading = true;

    this.accounts$ = this.accountService.getAccounts();
    this.balances$ = this.balanceService.getBalances();

    forkJoin(this.accounts$, this.balances$)
      .subscribe(([_, balances]) => {
          this.balances = balances;
          this.isLoading = false
        },
        () => this.isLoading = false
      );
  }

  getBalanceForAccount(accountId: string): number {
    return this.balances.find(_ => _.accountId === accountId).balance;
  }
}
