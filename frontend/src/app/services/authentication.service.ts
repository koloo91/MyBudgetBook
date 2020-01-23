import {Injectable} from '@angular/core';
import {environment} from '../../environments/environment';
import {BehaviorSubject, interval, Observable, Subscription} from 'rxjs';
import {HttpClient} from '@angular/common/http';
import {User} from '../models/user.model';
import {map, tap} from 'rxjs/operators';
import {LoginResponse} from '../models/login_response.model';

@Injectable({providedIn: 'root'})
export class AuthenticationService {
  private BASE_URL = `${environment.host}/lgn`;
  private currentUserSubject: BehaviorSubject<User>;
  public currentUser: Observable<User>;

  private refreshTokenInterval: Subscription;

  constructor(private http: HttpClient) {
    this.currentUserSubject = new BehaviorSubject<User>(JSON.parse(localStorage.getItem('currentUser')));
    this.currentUser = this.currentUserSubject.asObservable();
  }

  public get currentUserValue(): User {
    return this.currentUserSubject.value;
  }

  login(username: string, password: string): Observable<User> {
    return this.http.post<LoginResponse>(`${this.BASE_URL}/api/login`, {name: username, password: password})
      .pipe(
        map((response) => {
          let user = new User(username, response.accessToken, response.refreshToken);
          localStorage.setItem('currentUser', JSON.stringify(user));
          this.currentUserSubject.next(user);
          return user;
        }),
        tap(() => this.startRefreshInterval())
      );
  }

  startRefreshInterval() {
    this.refreshTokenInterval = interval(60 * 1000)
      .subscribe(() => {
        console.log('Will refresh token..');
      });
  }

  logout() {
    // remove user from local storage to log user out
    localStorage.removeItem('currentUser');
    this.currentUserSubject.next(null);
    this.refreshTokenInterval.unsubscribe();
  }
}
