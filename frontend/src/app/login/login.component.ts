import {Component, OnInit} from '@angular/core';
import {AuthenticationService} from '../services/authentication.service';
import {ActivatedRoute, Router} from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  username: string;
  password: string;
  isLoading: boolean;

  returnUrl: string;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private authenticationService: AuthenticationService) {
  }

  ngOnInit() {
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
  }

  login() {
    this.isLoading = true;
    this.authenticationService.login(this.username, this.password)
      .subscribe(() => {
        this.isLoading = false
        this.router.navigate([this.returnUrl]);
      }, (err) => {
        this.isLoading = false
        console.log(err);
      });
  }
}
