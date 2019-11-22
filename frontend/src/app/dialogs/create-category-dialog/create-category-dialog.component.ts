import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {CategoryService} from '../../services/category.service';
import {Category} from '../../models/category.model';
import {ErrorService} from '../../services/error.service';
import {ErrorVo} from '../../models/error.model';
import {FormControl, FormGroup, Validators} from '@angular/forms';

@Component({
  selector: 'app-create-account-dialog',
  templateUrl: './create-category-dialog.component.html',
  styleUrls: ['./create-category-dialog.component.scss']
})
export class CreateCategoryDialogComponent implements OnInit {

  isLoading = false;

  categoryFormGroup: FormGroup;

  constructor(public dialogRef: MatDialogRef<CreateCategoryDialogComponent>,
              private categoryService: CategoryService,
              private errorService: ErrorService,
              @Inject(MAT_DIALOG_DATA) public category: Category) {

    this.categoryFormGroup = new FormGroup({
      'name': new FormControl(this.category.name, [Validators.required, Validators.minLength(1)])
    });
  }

  ngOnInit() {
  }

  onNoClick() {
    this.dialogRef.close({success: false});
  }

  onOkClick() {
    this.isLoading = true;

    if (this.category.id && this.category.id.length > 0) {
      this.updateCategory();
    } else {
      this.createCategory();
    }
  }

  updateCategory() {
    this.categoryService.updateCategory(this.category.id, this.categoryFormGroup.value).subscribe(category => {
      this.dialogRef.close({success: true});
    }, err => {
      this.errorService.showErrorMessage(err.error);
    });
  }

  createCategory() {
    this.categoryService.createCategory(this.categoryFormGroup.value).subscribe(category => {
      this.dialogRef.close({success: true});
    }, (err: ErrorVo) => {
      this.isLoading = false;
      this.errorService.showErrorMessage(err.message);
    });
  }
}
