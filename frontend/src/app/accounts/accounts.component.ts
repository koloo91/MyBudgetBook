import {Component, OnInit} from '@angular/core';
import {AccountService} from '../services/account.service';
import {Observable} from 'rxjs';
import {map} from 'rxjs/operators';

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

  constructor(private accountService: AccountService) {
  }

  ngOnInit() {
    this.accounts = this.accountService.getAccounts()
      .pipe(
        map(pagedEntity => pagedEntity.content)
      );
  }

}
