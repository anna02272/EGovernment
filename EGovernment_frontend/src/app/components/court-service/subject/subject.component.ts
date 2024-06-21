import { Component, Input } from '@angular/core';

import { Router } from '@angular/router';
import { Subject } from 'src/app/models/court/subject';

@Component({
  selector: 'app-subject',
  templateUrl: './subject.component.html',
  styleUrls: ['./subject.component.css']
})
export class SubjectComponent {
  @Input() subject: Subject | undefined;

  constructor(private router: Router) {}

  openSubjectDetails(): void {
    if (this.subject) {
      this.router.navigate(['/subject-details', this.subject.id]);
    }
  }
}
