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
import { ReportRegisteredVehiclesComponent } from './components/statistics-service/report-registered-vehicles/report-registered-vehicles.component';
import { ReportCarAccidentDegreeComponent } from './components/statistics-service/report-car-accident-degree/report-car-accident-degree.component';
import { ReportCarAccidentTypeComponent } from './components/statistics-service/report-car-accident-type/report-car-accident-type.component';
import { ReportDelictsComponent } from './components/statistics-service/report-delicts/report-delicts.component';
import { RequestsComponent } from './components/statistics-service/requests/requests.component';
import { ResponseComponent } from './components/statistics-service/response/response.component';
import { RequestService } from './services/statistics/request.service';
import { StatisticsHeaderComponent } from './components/statistics-service/statistics-header/statistics-header.component';
import { ResponseService } from './services/statistics/response.service';
import { ReportDelicTypeService } from './services/statistics/reportDelicType.service';
import { ReportCarAccidentTypeService } from './services/statistics/reportCarAccidentTypeService.service';
import { ReportCarAccidentDegreeService } from './services/statistics/reportCarAccidentDegree.service';
import { ReportRegisteredVehiclesService } from './services/statistics/reportRegisteredVehicles.service';

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
    RequestComponent,
    ReportRegisteredVehiclesComponent,
    ReportCarAccidentDegreeComponent,
    ReportCarAccidentTypeComponent,
    ReportDelictsComponent,
    RequestsComponent,
    ResponseComponent,
    StatisticsHeaderComponent
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
  RequestService,
  ResponseService,
  ReportRegisteredVehiclesService,
  ReportDelicTypeService,
  ReportCarAccidentTypeService,
  ReportCarAccidentDegreeService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
