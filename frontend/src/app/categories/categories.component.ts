import {Component, OnInit} from '@angular/core';
import {CategoryService} from '../services/category.service';
import {Category} from '../models/category.model';
import {MatDialog} from '@angular/material/dialog';
import {CreateCategoryDialogComponent} from '../dialogs/create-category-dialog/create-category-dialog.component';
import {ErrorService} from '../services/error.service';
import {ErrorVo} from '../models/error.model';

interface CategoryNode {
  id: string;
  name: string;
  children?: CategoryNode[];
}

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.scss']
})
export class CategoriesComponent implements OnInit {
  isLoading: boolean = true;

  categories: Category[] = [];

  constructor(private categoryService: CategoryService,
              private errorService: ErrorService,
              public dialog: MatDialog) {
  }

  ngOnInit() {
    this.loadCategories();
  }

  loadCategories() {
    this.isLoading = true;
    this.categoryService.getCategories()
      .subscribe((categories) => {
        this.isLoading = false;
        this.categories = categories;
      }, (err: ErrorVo) => {
        this.isLoading = false;
        this.errorService.showErrorMessage(err.message);
      });
  }

  showCreateDialog() {
    const dialogRef = this.dialog.open(CreateCategoryDialogComponent, {
      width: '600px'
    });

    dialogRef.afterClosed().subscribe(result => {
      this.loadCategories();
    });
  }

  updateCategory(selectedCategory: Category) {
    const dialogRef = this.dialog.open(CreateCategoryDialogComponent, {
      width: '600px',
      data: selectedCategory
    });

    dialogRef.afterClosed().subscribe(result => {
      this.loadCategories();
    });
  }
}
