import { Component } from '@angular/core';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  constructor( private userService: UserService,
    private authService: AuthService) 
  { }
  
  hasSignedIn() {
    return !!this.userService.currentUser;
  }

  getUsername() {
    return this.userService.currentUser.user.username;
  }

  logout() {
    this.authService.logout();
  }
  getRole() {
    return this.userService.currentUser.user.userRole;
  }
}
