import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Gender } from 'src/app/models/police/gender';
import { VehicleDriver } from 'src/app/models/police/vehicleDriver';
import { VehicleService } from 'src/app/services/vehicles/vehicleService';

@Component({
  selector: 'app-driver',
  templateUrl: './driver.component.html',
  styleUrls: ['./driver.component.css']
})
export class DriverComponent implements OnInit {
  vehicleDriverForm: FormGroup;
  genders = Object.values(Gender);
  backendError: string | null = null;
  vehicleDrivers: VehicleDriver[] = []; 
  searchID: string = ''; 



  constructor(
    private fb: FormBuilder,
    private vehicleService: VehicleService,
    private snackBar: MatSnackBar
  ) {
    this.vehicleDriverForm = this.fb.group({
      identification_number: ['', Validators.required],
      name: ['', Validators.required],
      last_name: ['', Validators.required],
      date_of_birth: ['', Validators.required],
      gender: ['', Validators.required]

    });
  }
  ngOnInit(): void {
    this.loadAllVehicleDrivers(); 
  }
  
  onSubmit(): void {
    if (this.vehicleDriverForm.invalid) {

      this.markAllAsTouched();
      return;
    }

    const dateOfBirth = this.vehicleDriverForm.get('date_of_birth')?.value;

    const dateOfBirthWithTime = dateOfBirth + 'T00:00:00Z';

    const newVehicleDriver: VehicleDriver = {
      ...this.vehicleDriverForm.value,
      date_of_birth: dateOfBirthWithTime
    };

    this.vehicleService.createVehicleDriver(newVehicleDriver).subscribe({
      next: () => {
        this.snackBar.open('Vozac kreiran.', 'Zatvori', {
          duration: 3000,
          panelClass: ['success-snackbar']
        });
        this.onCancel();
        this.backendError = null;
        this.loadAllVehicleDrivers();
      },
      error: (errorResponse) => {
        console.error('Greska prilikom kreiranja vozaca:', errorResponse);
        this.backendError = errorResponse.error?.error || 'Greska prilikom kreiranja vozaca. Pokusajte kasnije.';
        this.snackBar.open(this.backendError ?? 'Greska prilikom kreiranja vozaca. Pokusajte kasnije.', 'Zatvori', {
          duration: 3000,
          panelClass: ['error-snackbar']
        });
      }
    });
  }

  private markAllAsTouched(): void {
    this.vehicleDriverForm.markAllAsTouched();
  }

  onCancel(): void {
    this.vehicleDriverForm.reset();
    this.vehicleDriverForm.patchValue({
      gender: ''
    });
    this.backendError = null;
  }


  loadAllVehicleDrivers(): void {
    this.vehicleService.getAllVehicleDrivers().subscribe({
      next: (vehicleDrivers: VehicleDriver[]) => {
        this.vehicleDrivers = vehicleDrivers;
      },
      error: (errorResponse) => {
        console.error('Error fetching all vehicle drivers:', errorResponse);
      }
    });
  }


  searchVehicleDriversByID(): void {
    if (this.searchID.trim() === '') {
      return;
    }

    this.vehicleService.getDriverById(this.searchID.trim()).subscribe({
      next: (vehicleDriver: VehicleDriver) => {
        this.vehicleDrivers = [vehicleDriver]; 
      },
      error: (errorResponse) => {
        console.error('Error fetching vehicle by plate:', errorResponse);
      }
    });
  }


}
