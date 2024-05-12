import { Injectable, inject } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivateFn, Router, RouterStateSnapshot } from '@angular/router';
import { AuthService } from './auth.service';
import { UserRole } from '../models/userRole';

 @Injectable({
  providedIn: 'root'
})
export class PermissionsService  {

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    if (this.authService.tokenIsPresent()) {
      const roles = route.data['roles'] as UserRole[];
      const userRole = this.authService.getRole();

      if (roles && roles.length > 0 && roles.includes(userRole)) {
        if (userRole === UserRole.Citizen){
          this.router.navigate(['/pocetna']);
        } else if (userRole === UserRole.Employee){
          this.router.navigate(['/zavodZaStatistiku']);
        } else if (userRole === UserRole.Policeman){
          this.router.navigate(['/mupVozila']);
        } else if (userRole === UserRole.TrafficPoliceman){
          this.router.navigate(['/saobracajnaPolicija']);
        } else if (userRole === UserRole.Judge){
          this.router.navigate(['/prekrsajniSud']);   
        }
        return true;
      } else {
        this.router.navigate(['/prijava']);
        return false;
      }
    } else {
      this.router.navigate(['/prijava']);
      return false;
    }
  }
}

export const AuthGuard: CanActivateFn = (next: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean => {
  return inject(PermissionsService).canActivate(next, state);
}