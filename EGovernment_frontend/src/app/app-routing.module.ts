import { RouterModule, Routes } from '@angular/router';

import { AuthGuard } from './services/permission.service';
import { EditSubjectComponent } from './components/court-service/edit-subject/edit-subject.component';
import { HearingComponent } from './components/court-service/hearing/hearing.component';
import { HomeComponent } from './components/home/home.component';
import { HomeCourtComponent } from './components/court-service/home-court/home-court.component';
import { HomePoliceComponent } from './components/police-service/home-police/home-police.component';
import { HomeStatisticsComponent } from './components/statistics-service/home-statistics/home-statistics.component';
import { HomeVehiclesComponent } from './components/vehicles-service/home-vehicles/home-vehicles.component';
import { LoginComponent } from './components/auth-service/login/login.component';
import { NgModule } from '@angular/core';
import { RegistrationComponent } from './components/auth-service/registration/registration.component';
import { RequestsComponent } from './components/statistics-service/requests/requests.component';
import { SubjectDetailsComponent } from './components/court-service/subject-details/subject-details.component';
import { SubjectTabComponent } from './components/court-service/subject-tab/subject-tab.component';
import { DelictComponent } from './components/police-service/delict/delict.component';
import { DelictDetailsComponent } from './components/police-service/delict-details/delict-details.component';
import { CarAccidentComponent } from './components/police-service/carAccident/car-accident.component';
import { CarAccidentDetailsComponent } from './components/police-service/car-accident-details/car-accident-details.component';
import { CreateDelictComponent } from './components/police-service/create-delict/create-delict.component';
import { AllDelictsComponent } from './components/police-service/all-delicts/all-delicts.component';
import { AllCarAccidentsComponent } from './components/police-service/all-car-accidents/all-car-accidents.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: '/prijava',
    pathMatch: 'full'
  },
  {
    path: 'prijava',
    component: LoginComponent
  },
  {
    path: 'registracija',
    component: RegistrationComponent
  },
  {
    path: 'pocetna',
    component: HomeComponent,
    canActivate: [AuthGuard] 
  },

  {
    path: 'zavodZaStatistiku',
    component: HomeStatisticsComponent,
    canActivate: [AuthGuard] 
  },
  {
    path: 'zahtevi',
    component: RequestsComponent,
    canActivate: [AuthGuard] 
  },
  {
    path: 'mupVozila',
    component: HomeVehiclesComponent,
    canActivate: [AuthGuard] 

  },
  {
    path: 'saobracajnaPolicija',
    component: HomePoliceComponent,
    canActivate: [AuthGuard] 
  },
  {
    path: 'prekrsajniSud',
    component: HomeCourtComponent,
    canActivate: [AuthGuard] 
  },
  {
    path: 'subject-details/:id',
    component: SubjectDetailsComponent,
    canActivate: [AuthGuard]
  },
  {
    path: 'rociste/:subjectId',
    component: HearingComponent
  },
  { path: 'subjectTab/:id', component: SubjectTabComponent },
  { path: 'editSubject/:id', component: EditSubjectComponent },

  {
    path: 'prekrsaj',
    component: DelictComponent,
    canActivate: [AuthGuard] 
  },

  {
    path: 'delict-details/:id',
    component: DelictDetailsComponent,
    canActivate: [AuthGuard]
  },

  {
    path: 'nesreca',
    component: CarAccidentComponent,
    canActivate: [AuthGuard] 
  },

  {
    path: 'car-accident-details/:id',
    component: CarAccidentDetailsComponent,
    canActivate: [AuthGuard]
  },

  {
    path: 'create-delict',
    component: CreateDelictComponent,
    canActivate: [AuthGuard] 
  },

  { path: 'all-delicts',
     component: AllDelictsComponent ,
     canActivate: [AuthGuard] 
  },

  { path: 'all-car-accidents',
      component: AllCarAccidentsComponent ,
      canActivate: [AuthGuard] 
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }