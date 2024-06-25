import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CarAccident } from 'src/app/models/traffic-police/carAccident';
import { CarAccidentService } from 'src/app/services/traffic-police/carAccidentService';

@Component({
  selector: 'app-car-accident-details',
  templateUrl: './car-accident-details.component.html',
  styleUrls: ['./car-accident-details.component.css']
})
export class CarAccidentDetailsComponent implements OnInit {
  carAccident: CarAccident | undefined;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private carAccidentService: CarAccidentService
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const carAccidentId = params.get('id');
      if (carAccidentId) {
        this.carAccidentService.getCarAccidentById(carAccidentId).subscribe(
          carAccident => {
            this.carAccident = carAccident;
          },
          error => {
            console.error('Error loading car accident:', error);
          }
        );
      }
    });
  }

}