import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { CarAccidentCreate } from 'src/app/models/traffic-police/carAccidentCreate';
import { CarAccidentService } from 'src/app/services/traffic-police/carAccidentService';
import { CarAccidentType } from 'src/app/models/traffic-police/carAccidentType';
import { DegreeOfAccident } from 'src/app/models/traffic-police/degreeOfAccident';
import { UserService } from 'src/app/services/auth/user.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DelictService } from 'src/app/services/traffic-police/delictService';

@Component({
  selector: 'app-create-car-accident',
  templateUrl: './create-car-accident.component.html',
  styleUrls: ['./create-car-accident.component.css']
})
export class CreateCarAccidentComponent {
  carAccident: CarAccidentCreate = new CarAccidentCreate(
    '', '', '', '', '', '', '', CarAccidentType.KnockingDownPedestrians, DegreeOfAccident.NoHarm, 0
  );

  carAccidentTypes = Object.values(CarAccidentType);
  degreesOfAccident = Object.values(DegreeOfAccident);
  imageFiles: File[] = [];
  responseId: any;

  constructor(
    private carAccidentService: CarAccidentService,
    private router: Router,
    private userService: UserService,
    private snackBar: MatSnackBar,
    private delictService: DelictService
  ) {}

  getCurrentUserId() {
    return this.userService.currentUser?.user.id;
  }

  onFileChange(event: any): void {
    if (event.target.files && event.target.files.length) {
      this.imageFiles = Array.from(event.target.files);
    }
  }

  createCarAccident(): void {
    if (!this.isFormValid()) {
      this.openSnackBar("Molimo popunite sva polja", "");
      return;
    }

    this.carAccidentService.insertCarAccident(this.carAccident).subscribe(
      (response: any) => {
        this.responseId = response.id;
        this.uploadImages();
        this.openSnackBar("Saobracajna nesreca je uspesno kreirana!", "");
        this.router.navigate(['/all-car-accidents']);
      },
      error => {
        console.error('Greska pri kreiranju saobracajne nesrece:', error);
      }
    );
  }

  uploadImages(): void {
    if (this.imageFiles.length === 0) {
      this.openSnackBar("No images selected for upload!", "");
      console.warn('No images selected for upload.');
      return;
    }

    const formData = new FormData();
    for (const file of this.imageFiles) {
      formData.append('images', file);
    }

    this.delictService.uploadImages(this.responseId, formData).subscribe(
      () => {
        //console.log('Images uploaded successfully!');
        this.router.navigate(['/all-car-accidents']);
      },
      (error) => {
        console.error('Error uploading images:', error);
        this.openSnackBar("Error uploading images!", "");
      }
    );
  }

  isFormValid(): boolean {
    return (
      !!this.carAccident.driver_identification_number_first &&
      !!this.carAccident.driver_identification_number_second &&
      !!this.carAccident.vehicle_licence_number_first &&
      !!this.carAccident.vehicle_licence_number_second &&
      !!this.carAccident.driver_email &&
      !!this.carAccident.location &&
      !!this.carAccident.description &&
      !!this.carAccident.car_accident_type &&
      !!this.carAccident.degree_of_accident &&
      this.carAccident.number_of_penalty_points >= 0
    );
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {
      duration: 3000,
    });
  }
}
