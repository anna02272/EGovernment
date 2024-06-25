import { Component, OnInit } from '@angular/core';
import { CarAccident } from 'src/app/models/traffic-police/carAccident';
import { CarAccidentService } from 'src/app/services/traffic-police/carAccidentService';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-all-car-accidents',
  templateUrl: './all-car-accidents.component.html',
  styleUrls: ['./all-car-accidents.component.css']
})
export class AllCarAccidentsComponent implements OnInit {
  carAccidents: CarAccident[] = [];

  constructor(
    private carAccidentService: CarAccidentService,
    private userService: UserService
  ) { }

  ngOnInit() {
    this.loadCarAccidents();
  }

  loadCarAccidents() {
    const role = this.userService.currentUser?.user.userRole;

    if (role === 'TrafficPoliceman') {
      this.carAccidentService.getAllCarAccidentsByPolicemanID().subscribe(
        (carAccidents) => {
          this.carAccidents = carAccidents;
        },
        (error) => {
          console.error('Error loading car accidents:', error);
        }
      );
    } else if (role === 'Citizen') {
      this.carAccidentService.getAllCarAccidentsByDriver().subscribe(
        (carAccidents) => {
          this.carAccidents = carAccidents;
        },
        (error) => {
          console.error('Error loading car accidents:', error);
        }
      );
    }
  }
}