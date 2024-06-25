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
import { DelictService } from './services/traffic-police/delictService';
import { CarAccidentService } from './services/traffic-police/carAccidentService';
import { ReportRegisteredVehiclesService } from './services/statistics/reportRegisteredVehicles.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatDialogModule } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { CreateReportCarAccidentDegreeComponent } from './components/statistics-service/create-report-car-accident-degree/create-report-car-accident-degree.component';
import { CreateReportCarAccidentTypeComponent } from './components/statistics-service/create-report-car-accident-type/create-report-car-accident-type.component';
import { CreateReportDelictComponent } from './components/statistics-service/create-report-delict/create-report-delict.component';
import { CreateReportRegisteredVehiclesComponent } from './components/statistics-service/create-report-registered-vehicles/create-report-registered-vehicles.component';
import { HearingComponent } from './components/court-service/hearing/hearing.component';
import { SubjectComponent } from './components/court-service/subject/subject.component';
import { SubjectDetailsComponent } from './components/court-service/subject-details/subject-details.component';
import { ScheduleComponent } from './components/court-service/schedule/schedule.component';
import { SubjectTabComponent } from './components/court-service/subject-tab/subject-tab.component';
import { EditSubjectComponent } from './components/court-service/edit-subject/edit-subject.component';
import { DelictComponent } from './components/police-service/delict/delict.component';
import { CarAccidentComponent } from './components/police-service/carAccident/car-accident.component';
import { DelictDetailsComponent } from './components/police-service/delict-details/delict-details.component';
import { CarAccidentDetailsComponent } from './components/police-service/car-accident-details/car-accident-details.component';
import { CreateDelictComponent } from './components/police-service/create-delict/create-delict.component';
import { AllDelictsComponent } from './components/police-service/all-delicts/all-delicts.component';
import { AllCarAccidentsComponent } from './components/police-service/all-car-accidents/all-car-accidents.component';
import { PoliceHeaderComponent } from './components/police-service/police-header/police-header.component';

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
    StatisticsHeaderComponent,
    CreateReportCarAccidentDegreeComponent,
    CreateReportCarAccidentTypeComponent,
    CreateReportDelictComponent,
    CreateReportRegisteredVehiclesComponent,
    HearingComponent,
    SubjectComponent,
    SubjectDetailsComponent,
    ScheduleComponent,
    SubjectTabComponent,
    EditSubjectComponent,
    DelictComponent,
    CarAccidentComponent,
    DelictDetailsComponent,
    CarAccidentDetailsComponent,
    CreateDelictComponent,
    AllDelictsComponent,
    AllCarAccidentsComponent,
    PoliceHeaderComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    FormsModule,
    BrowserAnimationsModule,
    MatSnackBarModule,
    MatDialogModule,
    MatButtonModule,
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
  ReportCarAccidentDegreeService,
  DelictService,
  CarAccidentService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
