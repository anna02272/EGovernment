import { Component } from '@angular/core';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-home-court',
  templateUrl: './home-court.component.html',
  styleUrls: ['./home-court.component.css']
})
export class HomeCourtComponent {
  constructor( private userService: UserService) 
  { }
  
  getRole() {
    return this.userService.currentUser?.user.userRole;
  }
}
