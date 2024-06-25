import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Vehicle } from 'src/app/models/police/vehicle';
import { RefreshService } from 'src/app/services/refresh.service';
import { Category } from 'src/app/models/statisics/category';
import { VehicleModel } from 'src/app/models/police/vehicleModel';
import { VehicleService } from 'src/app/services/vehicles/vehicleService';

@Component({
  selector: 'app-vehicles',
  templateUrl: './vehicles.component.html',
  styleUrls: ['./vehicles.component.css']
})
export class VehiclesComponent implements OnInit {
  vehicleForm: FormGroup;
  vehicleModels = Object.values(VehicleModel);
  categories = Object.values(Category);
  backendError: string | null = null;
  vehicles: Vehicle[] = []; 
  searchCategory: string = 'B';
  searchYear: number | null = new Date().getFullYear();
  years: number[] = [];
  searchPlate: string = ''; 
  

  constructor(
    private fb: FormBuilder,
    private vehicleService: VehicleService,
    private snackBar: MatSnackBar
  ) {
    this.vehicleForm = this.fb.group({
      registration_plate: ['', Validators.required],
      vehicle_model: ['', Validators.required],
      vehicle_owner: ['', Validators.required],
      category: ['', Validators.required]
    });
  }

  ngOnInit(): void {
    this.loadAllVehicles(); 
    this.generateYears();
  }

  onSubmit(): void {
    if (this.vehicleForm.invalid) {
      console.log("API krece")

      this.markAllAsTouched();
      return;
    }

    const currentDate = new Date().toISOString(); 
    const newVehicle: Vehicle = {
      ...this.vehicleForm.value,
      registration_date: currentDate 
    };

    this.vehicleService.create(newVehicle).subscribe({
      next: () => {
        this.snackBar.open('Vozilo kreirano.', 'Zatvori', {
          duration: 3000,
          panelClass: ['success-snackbar']
        });
        this.onCancel();
        this.backendError = null;
        this.loadAllVehicles();
      },
      error: (errorResponse) => {
        console.error('Error creating vehicle:', errorResponse);
        this.backendError = errorResponse.error?.error || 'Greska prilikom kreiranja vozila. Pokusajte kasnije.';
        this.snackBar.open(this.backendError ?? 'Greska prilikom kreiranja vozila. Pokusajte kasnije.', 'Zatvori', {
          duration: 3000,
          panelClass: ['error-snackbar']
        });
      }
    });
  }

  onCancel(): void {
    this.vehicleForm.reset();
    this.vehicleForm.patchValue({
      vehicle_model: '',
      category: ''
    });
    this.backendError = null;
  }

  private markAllAsTouched(): void {
    this.vehicleForm.markAllAsTouched();
  }

   loadAllVehicles(): void {
    this.vehicleService.getAll().subscribe({
      next: (vehicles: Vehicle[]) => {
        this.vehicles = vehicles;
      },
      error: (errorResponse) => {
        console.error('Error fetching all vehicles:', errorResponse);
      }
    });
  }

  generateYears(): void {
    const currentYear = new Date().getFullYear();
    for (let year = 2000; year <= currentYear; year++) {
      console.log(year);
      this.years.push(year);
    }
  }

  searchVehicleByPlate(): void {
    if (this.searchPlate.trim() === '') {
      return;
    }

    this.vehicleService.getById(this.searchPlate.trim()).subscribe({
      next: (vehicle: Vehicle) => {
        this.vehicles = [vehicle]; 
      },
      error: (errorResponse) => {
        console.error('Error fetching vehicle by plate:', errorResponse);
      }
    });
  }


  searchVehiclesByCategoryAndYear(): void {
    if (!this.searchCategory || !this.searchYear) {
      return;
    }

    this.vehicleService.getByCategoryAndYear(this.searchCategory, this.searchYear).subscribe({
      next: (vehicles: Vehicle[]) => {
        this.vehicles = vehicles;
      },
      error: (errorResponse) => {
        console.error('Error fetching vehicles by category and year:', errorResponse);
      }
    });
  }
}
