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

  navigateToCreateCarAccident() {
    this.router.navigate(['/create-car-accident']);
  }
}
