import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Vehicle } from 'src/app/models/police/vehicle';
import { RefreshService } from 'src/app/services/refresh.service';
import { Category } from 'src/app/models/statisics/category';
import { VehicleModel } from 'src/app/models/police/vehicleModel';
import { VehicleService } from 'src/app/services/vehicles/vehicleService';
import { ResponseCount } from 'src/app/models/police/responseCount';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-registered-vehicles',
  templateUrl: './registered-vehicles.component.html',
  styleUrls: ['./registered-vehicles.component.css']
})
export class RegisteredVehiclesComponent implements OnInit {
  vehicleForm: FormGroup;
  vehicleModels = Object.values(VehicleModel);
  categories = Object.values(Category);
  searchCategory: string = 'B';
  responseCount?: ResponseCount;
  backendError: string | null = null;
  searchCategoryPdf: Category = Category.B;
  registeredVehicles: Vehicle[] = []; 
  vehicleCountsByCategory: { [key: string]: number } = {};
  searchPlate: string = ''; 

  constructor(
    private fb: FormBuilder,
    private vehicleService: VehicleService,
    private refreshService: RefreshService,
    private userService: UserService,

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
    this.fetchVehicleCountsByCategory();

  }

  onSubmit(): void {
    if (this.vehicleForm.invalid) {

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
    this.vehicleService.getAllRegisteredVehicles().subscribe({
      next: (registeredVehicles: Vehicle[]) => {
        this.registeredVehicles = registeredVehicles;
      },
      error: (errorResponse) => {
        console.error('Error fetching all registered vehicles:', errorResponse);
      }
    });
  }

  getRole() {
    return this.userService.currentUser?.user.userRole;
  }


  downloadPdf(): void {
    this.vehicleService.getRegisteredVehiclesPdf().subscribe({
      next: (pdfBlob: Blob) => {
        const url = window.URL.createObjectURL(pdfBlob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'Registered_Vehicles_Report.pdf';
        a.click();
        window.URL.revokeObjectURL(url);
      },
      error: (error) => {
        console.error('Error downloading PDF:', error);
      }
    });
  }


  downloadPdfCategory(searchCategory: string): void {
    this.vehicleService.getRegisteredVehiclesCategoryPdf(searchCategory).subscribe({
      next: (pdfBlob: Blob) => {
        const url = window.URL.createObjectURL(pdfBlob);
        const a = document.createElement('a');
        a.href = url;
        a.download = `Registered_Vehicles_Report_${searchCategory}.pdf`;
        a.click();
        window.URL.revokeObjectURL(url);
      },
      error: (error) => {
        console.error('Error downloading PDF:', error);
      }
    });
  
  }


  fetchVehicleCountsByCategory(): void {
    this.categories.forEach(category => {
      this.vehicleService.getNumberOfRegVehiclesCategory(category).subscribe({
        next: (response: ResponseCount) => {
          this.vehicleCountsByCategory[category] = response.count;
        },
        error: (errorResponse) => {
          console.error(`Error fetching count for category ${category}:`, errorResponse);
          this.vehicleCountsByCategory[category] = 0;
        }
      });
    });
}

}
