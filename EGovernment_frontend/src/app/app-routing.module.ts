import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { HomeComponent } from './components/home/home.component';
import { HomeStatisticsComponent } from './components/home-statistics/home-statistics.component';
import { HomeVehiclesComponent } from './components/home-vehicles/home-vehicles.component';
import { HomePoliceComponent } from './components/home-police/home-police.component';
import { HomeCourtComponent } from './components/home-court/home-court.component';
import { AuthGuard } from './services/permission.service';

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