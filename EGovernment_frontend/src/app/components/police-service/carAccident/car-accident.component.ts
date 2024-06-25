import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';
import { CarAccident } from 'src/app/models/traffic-police/carAccident';

@Component({
  selector: 'app-car-accident',
  templateUrl: './car-accident.component.html',
  styleUrls: ['./car-accident.component.css']
})
export class CarAccidentComponent {
  @Input() carAccident: CarAccident | undefined;

  constructor(private router: Router) {}

  openCarAccidentDetails(): void {
    if (this.carAccident) {
      this.router.navigate(['/car-accident-details', this.carAccident.id]);
    }
  }
}