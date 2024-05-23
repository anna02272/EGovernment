import { Injectable } from '@angular/core';
import { HttpHeaders } from '@angular/common/http';
import { ApiService } from './api.service';
import { ConfigService } from './config.service';
import { Router } from '@angular/router';
import { map } from 'rxjs/operators';
import { UserService } from './user.service';
import { UserRole } from '../models/userRole';

@Injectable()
export class AuthService {

  constructor(
    private apiService: ApiService,
    private userService: UserService,
    private config: ConfigService,
    private router: Router
  ) {
  }

  private access_token = null;

  login(user: any) {
    const loginHeaders = new HttpHeaders({
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    });

    const body = {
      'email': user.email,
      'password': user.password
    };
    return this.apiService.post(this.config.login_url, JSON.stringify(body), loginHeaders)
      .pipe(map((res) => {
        console.log('Login success');
        this.access_token = res.accessToken;
        localStorage.setItem("jwt", res.accessToken)
        return this.userService.getMyInfo();
      }));
  }

  register(user: any) {
    const signupHeaders = new HttpHeaders({
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    });
    const body = {
      'username': user.username,
      'password': user.password,
      'email' : user.email,
      'jmbg': user.jmbg,
      'name' : user.name,
      'lastname' : user.lastname,
      'userRole' :  UserRole.Citizen
    };
    return this.apiService.post(this.config.register_url, JSON.stringify(body), signupHeaders)
      .pipe(map(() => {
        console.log('Register success');
      }));
  }

  logout() {
    this.userService.currentUser = null;
    this.access_token = null;
    this.router.navigate(['/prijava']);
  }

  tokenIsPresent() {
    return this.access_token != undefined && this.access_token != null;
  }

  getToken() {
    return this.access_token;
  }

  getRole() {
    return this.userService.currentUser.user.userRole;
  }
 
}
