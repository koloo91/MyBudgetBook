import {async, ComponentFixture, TestBed} from '@angular/core/testing';

import {CreateBookingDialogComponent} from './create-booking-dialog.component';

describe('CreateAccountDialogComponent', () => {
  let component: CreateBookingDialogComponent;
  let fixture: ComponentFixture<CreateBookingDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [CreateBookingDialogComponent]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateBookingDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
