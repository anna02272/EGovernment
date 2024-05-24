import { Component } from '@angular/core';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-home-vehicles',
  templateUrl: './home-vehicles.component.html',
  styleUrls: ['./home-vehicles.component.css']
})
export class HomeVehiclesComponent {
  constructor( private userService: UserService) 
  { }
  
  getRole() {
    return this.userService.currentUser?.user.userRole;
  }
}
