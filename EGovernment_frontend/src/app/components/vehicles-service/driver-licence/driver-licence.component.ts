import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DriverLicence } from 'src/app/models/police/driverLicence';
import { VehicleService } from 'src/app/services/vehicles/vehicleService';
import { Category } from 'src/app/models/police/category';
import { Location } from 'src/app/models/police/location';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-driver-licence',
  templateUrl: './driver-licence.component.html',
  styleUrls: ['./driver-licence.component.css']
})
export class DriverLicenceComponent implements OnInit {
  driverLicenceForm: FormGroup;
  locations = Object.values(Location);
  categories = Object.values(Category);
  backendError: string | null = null;
  driverLicences: DriverLicence[] = [];
  searchIDByDriver: string = '';
  searchIDByLicence: string = '';

  constructor(
    private fb: FormBuilder,
    private vehicleService: VehicleService,
    private snackBar: MatSnackBar,
    private userService: UserService,
  ) {
    this.driverLicenceForm = this.fb.group({
      vehicle_driver: ['', Validators.required],
      // licence_number: ['', Validators.required],
      location_licenced: ['', Validators.required],
      categories: [[], Validators.required]
    });
  }

  ngOnInit(): void {
    this.loadAllDriverLicences();
  }
  onSubmit(): void {
    if (this.driverLicenceForm.invalid) {
      console.log("Form is invalid", this.driverLicenceForm);
      this.markAllAsTouched();
      console.log("Vehicle Driver Control:", this.driverLicenceForm.controls['vehicle_driver'].errors);
      console.log("Licence Number Control:", this.driverLicenceForm.controls['licence_number'].errors);
      console.log("Location Licenced Control:", this.driverLicenceForm.controls['location_licenced'].errors);
      console.log("Categories Control:", this.driverLicenceForm.controls['categories'].errors);
      return;
    }
    console.log("here in function");


    const newDriverLicence: DriverLicence = {
      ...this.driverLicenceForm.value,
      categories: this.driverLicenceForm.value.categories
    };


    this.vehicleService.createDriverLicence(newDriverLicence).subscribe({
      next: () => {
        this.snackBar.open('Vozacka dozvola kreirana.', 'Zatvori', {
          duration: 3000,
          panelClass: ['success-snackbar']
        });
        this.onCancel();
        this.backendError = null;
        this.loadAllDriverLicences();
      },
      error: (errorResponse) => {
        console.error('Greska prilikom kreiranja vozacke dozvole:', errorResponse);
        this.backendError = errorResponse.error?.error || 'Greska prilikom kreiranja vozacke dozvole. Pokusajte kasnije.';
        this.snackBar.open(this.backendError ?? 'Greska prilikom kreiranja vozacke dozvole. Pokusajte kasnije.', 'Zatvori', {
          duration: 3000,
          panelClass: ['error-snackbar']
        });
      }
    });
  }

  private markAllAsTouched(): void {
    this.driverLicenceForm.markAllAsTouched();
  }

  onCancel(): void {
    this.driverLicenceForm.reset();
    this.driverLicenceForm.patchValue({
      location_licenced: '',
      categories: []
    });
    this.backendError = null;
  }



  loadAllDriverLicences(): void {
    this.vehicleService.getAllLicences().subscribe({
      next: (driverLicences: DriverLicence[]) => {

        driverLicences.forEach(function (DriverLicence) {
          console.log(DriverLicence);
      });
        this.driverLicences = driverLicences;
      },
      error: (errorResponse) => {
        console.error('Error fetching all driver licences:', errorResponse);
      }
    });
  }


  searchDriverLicencesByID(): void {
    if (this.searchIDByLicence.trim() === '') {
      return;
    }

    this.vehicleService.getDriverLicenceById(this.searchIDByLicence.trim()).subscribe({
      next: (driverLicence: DriverLicence) => {
        this.driverLicences = [driverLicence];
      },
      error: (errorResponse) => {
        console.error('Error fetching driver licences by ID:', errorResponse);
      }
    });
  }

  searchDriverLicencesByDriver(): void {
    if (this.searchIDByDriver.trim() === '') {
      return;
    }

    this.vehicleService.getDriverLicenceByDriver(this.searchIDByDriver.trim()).subscribe({
      next: (driverLicence: DriverLicence) => {
        this.driverLicences = [driverLicence];
      },
      error: (errorResponse) => {
        console.error('Error fetching driver licences by ID:', errorResponse);
      }
    });
  }

  getRole() {
    return this.userService.currentUser?.user.userRole;
  }

}
