import {async, ComponentFixture, TestBed} from '@angular/core/testing';

import {UpdateBookingDialogComponent} from './update-booking-dialog.component';

describe('UpdateBookingDialogComponent', () => {
  let component: UpdateBookingDialogComponent;
  let fixture: ComponentFixture<UpdateBookingDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [UpdateBookingDialogComponent]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UpdateBookingDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
