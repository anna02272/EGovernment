import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';
import { Delict } from 'src/app/models/traffic-police/delict';

@Component({
  selector: 'app-delict',
  templateUrl: './delict.component.html',
  styleUrls: ['./delict.component.css']
})
export class DelictComponent {
  @Input() delict: Delict | undefined;

  constructor(private router: Router) {}

  openDelictDetails(): void {
    if (this.delict) {
      this.router.navigate(['/delict-details', this.delict.id]);
    }
  }
}