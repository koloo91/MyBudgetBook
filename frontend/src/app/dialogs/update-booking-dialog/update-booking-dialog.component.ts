import {Component, OnInit} from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';

@Component({
  selector: 'app-update-booking-dialog',
  templateUrl: './update-booking-dialog.component.html',
  styleUrls: ['./update-booking-dialog.component.scss']
})
export class UpdateBookingDialogComponent implements OnInit {

  constructor(public dialogRef: MatDialogRef<UpdateBookingDialogComponent>) {
  }

  ngOnInit() {
  }

  onJustThisClick() {
    this.dialogRef.close({updateAll: false});
  }

  onAllClick() {
    this.dialogRef.close({updateAll: true});
  }
}
