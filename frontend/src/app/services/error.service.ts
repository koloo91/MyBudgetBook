import {Injectable} from '@angular/core';
import {MatSnackBar} from '@angular/material/snack-bar';

@Injectable({
  providedIn: 'root'
})
export class ErrorService {

  constructor(private snackBar: MatSnackBar) {
  }

  showErrorMessage(errorMessage: string) {
    console.log(errorMessage);
    this.snackBar.open(errorMessage, 'OK', {duration: 2000});
  }
}
