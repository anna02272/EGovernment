import { Component, OnInit } from '@angular/core';
import { Delict } from 'src/app/models/traffic-police/delict';
import { CarAccident } from 'src/app/models/traffic-police/carAccident';
import { DelictService } from 'src/app/services/traffic-police/delictService';
import { CarAccidentService } from 'src/app/services/traffic-police/carAccidentService';
import { UserService } from 'src/app/services/auth/user.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-police',
  templateUrl: './home-police.component.html',
  styleUrls: ['./home-police.component.css'],
})

export class HomePoliceComponent implements OnInit {

  constructor(
    private userService: UserService,
    private router: Router
  ) {}

  ngOnInit() {}

  getRole() {
    return this.userService.currentUser?.user.userRole;
  }

  navigateToAllDelicts() {
    this.router.navigate(['/all-delicts']);
  }

  navigateToAllCarAccidents() {
    this.router.navigate(['/all-car-accidents']);
  }

  navigateToCreateDelict() {
    this.router.navigate(['/create-delict']);
  }
}
/*export class HomePoliceComponent implements OnInit {
  delicts: Delict[] = [];
  carAccidents: CarAccident[] = [];

  constructor(
    private userService: UserService,
    private delictService: DelictService,
    private carAccidentService: CarAccidentService,
    private router: Router,
  ) {}

  ngOnInit() {
    this.loadDelicts();
    this.loadCarAccidents();
  }

  getRole() {
    return this.userService.currentUser?.user.userRole;
  }

  loadDelicts() {
    this.delictService.getAllDelicts().subscribe(
      (delicts) => {
        this.delicts = delicts;
      },
      (error) => {
        console.error('Error loading delicts:', error);
      },
    );
  }

  loadCarAccidents() {
    this.carAccidentService.getAllCarAccidents().subscribe(
      (carAccidents) => {
        this.carAccidents = carAccidents;
      },
      (error) => {
        console.error('Error loading car accidents:', error);
      },
    );
  }

  navigateToDelictDetails(delictId: string) {
    this.router.navigate(['/delict-details', delictId]);
  }

  navigateToCarAccidentDetails(carAccidentId: string) {
    this.router.navigate(['/car-accident-details', carAccidentId]);
  }
  navigateToCreateDelict() {
    this.router.navigate(['/create-delict']);
  }
}*/
