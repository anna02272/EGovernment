import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { HomeComponent } from './components/home/home.component';
import { ConfigService } from './services/config.service';
import { ApiService } from './services/api.service';
import { RefreshService } from './services/refresh.service';
import { LoginComponent } from './components/login/login.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { TokenInterceptor } from './interceptor/TokenInterceptor';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { AuthService } from './services/auth.service';
import { UserService } from './services/user.service';
import { ReactiveFormsModule } from '@angular/forms';
import { HomeStatisticsComponent } from './components/home-statistics/home-statistics.component';
import { HomeCourtComponent } from './components/home-court/home-court.component';
import { HomeVehiclesComponent } from './components/home-vehicles/home-vehicles.component';
import { HomePoliceComponent } from './components/home-police/home-police.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    HomeComponent,
    LoginComponent,
    RegistrationComponent,
    HomeStatisticsComponent,
    HomeCourtComponent,
    HomeVehiclesComponent,
    HomePoliceComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [
  {
    provide: HTTP_INTERCEPTORS,
    useClass: TokenInterceptor,
    multi: true,
  },
  ConfigService,
  ApiService,
  RefreshService,
  AuthService,
  UserService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }