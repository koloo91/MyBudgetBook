import {BrowserModule} from '@angular/platform-browser';
import {LOCALE_ID, NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatButtonModule} from '@angular/material/button';
import {MatIconModule} from '@angular/material/icon';
import {MatListModule} from '@angular/material/list';
import {AccountsComponent} from './accounts/accounts.component';
import {HTTP_INTERCEPTORS, HttpClientModule} from '@angular/common/http';
import {FlexLayoutModule} from '@angular/flex-layout';
import {CreateAccountDialogComponent} from './accounts/create-account-dialog/create-account-dialog.component';
import {MatDialogModule} from '@angular/material/dialog';
import {MatFormFieldModule} from '@angular/material/form-field';
import {FormsModule} from '@angular/forms';
import {MatInputModule} from '@angular/material/input';
import {CategoriesComponent} from './categories/categories.component';
import {CreateCategoryDialogComponent} from './categories/create-category-dialog/create-category-dialog.component';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {BasicAuthInterceptor} from './helper/basic-auth.interceptor';
import {ErrorInterceptor} from './helper/error.interceptor';
import {LoginComponent} from './login/login.component';
import {HomeComponent} from './home/home.component';
import {MatCardModule} from '@angular/material/card';
import {BookingsComponent} from './bookings/bookings.component';
import {CreateBookingDialogComponent} from './dialogs/create-booking-dialog/create-booking-dialog.component';
import {MatSelectModule} from '@angular/material/select';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatNativeDateModule} from '@angular/material/core';
import {registerLocaleData} from '@angular/common';
import localDe from '@angular/common/locales/de';
import localeDeExtra from '@angular/common/locales/extra/de';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatExpansionModule} from '@angular/material/expansion';
import {UpdateBookingDialogComponent} from './dialogs/update-booking-dialog/update-booking-dialog.component';

registerLocaleData(localDe, 'de-DE', localeDeExtra)

@NgModule({
  declarations: [
    AppComponent,
    AccountsComponent,
    CreateAccountDialogComponent,
    CreateCategoryDialogComponent,
    CreateBookingDialogComponent,
    CategoriesComponent,
    LoginComponent,
    HomeComponent,
    BookingsComponent,
    UpdateBookingDialogComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FlexLayoutModule,
    HttpClientModule,
    MatToolbarModule,
    MatSidenavModule,
    MatButtonModule,
    MatIconModule,
    MatListModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatProgressSpinnerModule,
    MatCardModule,
    MatSelectModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatCheckboxModule,
    MatExpansionModule
  ],
  entryComponents: [
    CreateAccountDialogComponent,
    CreateCategoryDialogComponent,
    CreateBookingDialogComponent,
    UpdateBookingDialogComponent
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS, useClass: BasicAuthInterceptor, multi: true
    },
    {
      provide: HTTP_INTERCEPTORS, useClass: ErrorInterceptor, multi: true
    },
    {
      provide: LOCALE_ID, useValue: 'de-DE'
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {

}
