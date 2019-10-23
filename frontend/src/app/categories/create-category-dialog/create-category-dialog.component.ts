import {Component, OnInit} from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';
import {CategoryService} from '../../services/category.service';

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-category-dialog.component.html',
  styleUrls: ['./create-category-dialog.component.scss']
})
export class CreateCategoryDialogComponent implements OnInit {

  categoryName: string;
  isLoading = false;

  constructor(public dialogRef: MatDialogRef<CreateCategoryDialogComponent>,
              private categoryService: CategoryService) {
  }

  ngOnInit() {
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  createAccount() {
    this.isLoading = true;
    console.log(this.categoryName);
    this.categoryService.createCategory(this.categoryName).subscribe(category => {
        console.log(category);
        this.dialogRef.close({success: true});
      }, err => console.log(err),
      () => this.isLoading = false);
  }
}
