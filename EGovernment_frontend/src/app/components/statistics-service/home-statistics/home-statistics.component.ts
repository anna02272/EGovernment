import { Component } from '@angular/core';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-home-statistics',
  templateUrl: './home-statistics.component.html',
  styleUrls: ['./home-statistics.component.css']
})
export class HomeStatisticsComponent {
  selectedReport: string = 'registered-vehicles';

  constructor( private userService: UserService) 
  { }
  
  selectReport(report: string): void {
    this.selectedReport = report;
  }
  
  getRole() {
    return this.userService.currentUser?.user.userRole;
  }
}
