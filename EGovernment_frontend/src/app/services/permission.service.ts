import { Injectable, inject } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivateFn, Router, RouterStateSnapshot } from '@angular/router';
import { AuthService } from './auth/auth.service';

 @Injectable({
  providedIn: 'root'
})
export class PermissionsService  {

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  canActivate(route: ActivatedRouteSnapshot, _state: RouterStateSnapshot): boolean {
    if (this.authService.tokenIsPresent()) {
      const roles = route.data['roles'] as string[];
      if (roles && roles.length > 0) {
        const userRole = this.authService.getRole();

        if (roles.includes(userRole)) {
          return true;
        } else {
          this.router.navigate(['/prijava']);
          return false;
        }
      }
      return true;
    } else {
      this.router.navigate(['/prijava']);
      return false;
    }
  }
}

export const AuthGuard: CanActivateFn = (next: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean => {
  return inject(PermissionsService).canActivate(next, state);
}