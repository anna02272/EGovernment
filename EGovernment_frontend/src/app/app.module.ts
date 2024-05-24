import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './components/home/home.component';
import { ConfigService } from './services/config.service';
import { ApiService } from './services/api.service';
import { RefreshService } from './services/refresh.service';
import { LoginComponent } from './components/auth-service/login/login.component';
import { RegistrationComponent } from './components/auth-service/registration/registration.component';
import { TokenInterceptor } from './interceptor/TokenInterceptor';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { UserService } from './services/auth/user.service';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HomeStatisticsComponent } from './components/statistics-service/home-statistics/home-statistics.component';
import { HomeCourtComponent } from './components/court-service/home-court/home-court.component';
import { HomeVehiclesComponent } from './components/vehicles-service/home-vehicles/home-vehicles.component';
import { HomePoliceComponent } from './components/police-service/home-police/home-police.component';
import { HeaderComponent } from './components/header/header.component';
import { AuthService } from './services/auth/auth.service';
import { RequestComponent } from './components/statistics-service/request/request.component';

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
    HomePoliceComponent,
    RequestComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    FormsModule
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
