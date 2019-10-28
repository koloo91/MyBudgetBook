import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {CategoryService} from '../../services/category.service';
import {Category} from '../../models/category.model';

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-category-dialog.component.html',
  styleUrls: ['./create-category-dialog.component.scss']
})
export class CreateCategoryDialogComponent implements OnInit {

  categoryName: string;
  isLoading = false;

  constructor(public dialogRef: MatDialogRef<CreateCategoryDialogComponent>,
              private categoryService: CategoryService,
              @Inject(MAT_DIALOG_DATA) public data?: Category) {
    if (data) {
      this.categoryName = data.name;
    }
  }

  ngOnInit() {
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  onOkClick() {
    this.isLoading = true;

    if (this.data) {
      this.updateAccount();
    } else {
      this.createAccount();
    }
  }

  updateAccount() {
    this.categoryService.updateCategory(this.data.id, this.categoryName).subscribe(category => {
        console.log(category);
        this.dialogRef.close({success: true});
      }, err => console.log(err),
      () => this.isLoading = false);
  }

  createAccount() {
    this.categoryService.createCategory(this.categoryName).subscribe(category => {
        console.log(category);
        this.dialogRef.close({success: true});
      }, err => console.log(err),
      () => this.isLoading = false);
  }
}
