import {HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from '@angular/common/http';
import {Observable} from 'rxjs';
import {AuthenticationService} from '../services/authentication.service';
import {Injectable} from '@angular/core';

@Injectable()
export class BasicAuthInterceptor implements HttpInterceptor {
  constructor(private authenticationService: AuthenticationService) {
  }

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    console.log('BasicAuthInterceptor: intercept');
    // add authorization header with basic auth credentials if available
    const currentUser = this.authenticationService.currentUserValue;
    console.log(currentUser);
    if (currentUser && currentUser.token) {
      console.log('Set auth header');
      request = request.clone({
        setHeaders: {
          Authorization: `Basic ${currentUser.token}`
        }
      });
    }

    return next.handle(request);
  }
}