import {Component, OnInit} from '@angular/core';
import {CategoryService} from '../services/category.service';
import {Observable} from 'rxjs';
import {Category} from '../models/category.model';
import {map} from 'rxjs/operators';
import {MatDialog} from '@angular/material/dialog';
import {CreateCategoryDialogComponent} from './create-category-dialog/create-category-dialog.component';

interface CategoryNode {
  id: string;
  parentId: string;
  name: string;
  children?: CategoryNode[];
}

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.scss']
})
export class CategoriesComponent implements OnInit {

  categories: Observable<Category[]>;

  constructor(private categoryService: CategoryService,
              public dialog: MatDialog) {
  }

  ngOnInit() {
    this.loadCategories();
  }

  loadCategories() {
    this.categories = this.categoryService.getCategories()
      .pipe(map(result => result.content));
  }

  showCreateDialog() {
    const dialogRef = this.dialog.open(CreateCategoryDialogComponent, {
      width: '250px'
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
      this.loadCategories();
    });
  }

}