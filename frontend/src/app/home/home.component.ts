import {Component, OnInit, ViewChild} from '@angular/core';
import {Router} from '@angular/router';
import {MatDrawer} from '@angular/material/sidenav';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  title = 'My Budget Book';

  @ViewChild('drawer', {static: false})
  private matDrawer: MatDrawer;

  constructor(private router: Router) {
  }

  ngOnInit() {
    this.router.events.subscribe((event) => {
      this.matDrawer.close();
    })
  }
}
