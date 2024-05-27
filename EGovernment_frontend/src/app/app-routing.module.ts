import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './components/auth-service/login/login.component';
import { RegistrationComponent } from './components/auth-service/registration/registration.component';
import { HomeComponent } from './components/home/home.component';
import { HomeStatisticsComponent } from './components/statistics-service/home-statistics/home-statistics.component';
import { HomeVehiclesComponent } from './components/vehicles-service/home-vehicles/home-vehicles.component';
import { HomePoliceComponent } from './components/police-service/home-police/home-police.component';
import { HomeCourtComponent } from './components/court-service/home-court/home-court.component';
import { AuthGuard } from './services/permission.service';
import { RequestsComponent } from './components/statistics-service/requests/requests.component';
import { ReportRegisteredVehiclesComponent } from './components/statistics-service/report-registered-vehicles/report-registered-vehicles.component';

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
    path: 'izvestaj_o_tipu_saobracajne_nesrece',
    component: ReportRegisteredVehiclesComponent,
    // canActivate: [AuthGuard] 
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
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }