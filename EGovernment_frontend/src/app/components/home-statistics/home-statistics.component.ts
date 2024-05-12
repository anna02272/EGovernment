import { Component } from '@angular/core';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-home-statistics',
  templateUrl: './home-statistics.component.html',
  styleUrls: ['./home-statistics.component.css']
})
export class HomeStatisticsComponent {
  constructor( private userService: UserService) 
  { }
  
  getRole() {
    return this.userService.currentUser?.user.userRole;
  }
}
