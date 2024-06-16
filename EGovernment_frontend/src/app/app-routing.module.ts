import { RouterModule, Routes } from '@angular/router';

import { AuthGuard } from './services/permission.service';
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
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }