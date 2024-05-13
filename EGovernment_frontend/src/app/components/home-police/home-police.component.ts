import { Component } from '@angular/core';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-home-police',
  templateUrl: './home-police.component.html',
  styleUrls: ['./home-police.component.css']
})
export class HomePoliceComponent {
  constructor( private userService: UserService) 
  { }
  
  getRole() {
    return this.userService.currentUser?.user.userRole;
  }
}
