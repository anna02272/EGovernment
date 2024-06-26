import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CarAccident } from 'src/app/models/traffic-police/carAccident';
import { CarAccidentService } from 'src/app/services/traffic-police/carAccidentService';
import { DelictService } from 'src/app/services/traffic-police/delictService';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-car-accident-details',
  templateUrl: './car-accident-details.component.html',
  styleUrls: ['./car-accident-details.component.css']
})
export class CarAccidentDetailsComponent implements OnInit {
  carAccident: CarAccident | undefined;
  imagesUrl: any[] = [];
  images: any[] = [];
  imageFiles: File[] = [];
  currentSlideIndex = 0;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private carAccidentService: CarAccidentService,
    private delictService: DelictService,
    private snackBar: MatSnackBar,
    private userService: UserService,
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const carAccidentId = params.get('id');
      if (carAccidentId) {
        this.carAccidentService.getCarAccidentById(carAccidentId).subscribe(
          carAccident => {
            this.carAccident = carAccident;
            this.loadImages(carAccident.id);
          },
          error => {
            console.error('Error loading car accident:', error);
          }
        );
      }
    });
  }

  loadImages(carAccidentId: string): void {
    this.delictService.getImagesUrls(carAccidentId).subscribe(
      (images: any) => {
        this.imagesUrl = images;
        if (this.imagesUrl.length !== 0) {
          this.imagesUrl.forEach(imageName => this.loadImageContent(carAccidentId, imageName));
        }
      },
      error => {
        console.error('Error getting delict images:', error);
      }
    );
  }

  loadImageContent(folderName: string, imageName: string): void {
    this.delictService.getImages(folderName, imageName).subscribe(
      (imageContent: any) => {
        this.displayImage(imageContent);
      },
      error => {
        console.error('Error getting images:', error);
      }
    );
  }

  displayImage(blob: Blob) {
    const reader = new FileReader();
    reader.onload = (e: any) => {
      this.images.push(e.target.result);
    };
    reader.readAsDataURL(blob);
  }

  onFileChange(event: any): void {
    if (event.target.files && event.target.files.length) {
      this.imageFiles = Array.from(event.target.files);
    }
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

    if (this.carAccident) {
      this.delictService.uploadImages(this.carAccident.id, formData).subscribe(
        () => {
          console.log('Images uploaded successfully!');
          this.openSnackBar("Images uploaded successfully!", "");
          this.ngOnInit(); // Reload the page after successful upload
        },
        (error) => {
          console.error('Error uploading images:', error);
          this.openSnackBar("Error uploading images!", "");
        }
      );
    } else {
      console.error('Delict is undefined, cannot upload images.');
    }
  }

  prevSlide() {
    if (this.currentSlideIndex > 0) {
      this.currentSlideIndex--;
    }
  }

  nextSlide() {
    if (this.currentSlideIndex < this.images.length - 1) {
      this.currentSlideIndex++;
    }
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {
      duration: 3000,
    });
  }

  getRole() {
    return this.userService.currentUser?.user.userRole;
  }

}