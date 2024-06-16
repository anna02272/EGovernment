import { Component, OnInit } from '@angular/core';

import { Subject } from 'src/app/models/court/subject';
import { SubjectService } from 'src/app/services/court/subject.service';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-home-court',
  templateUrl: './home-court.component.html',
  styleUrls: ['./home-court.component.css']
})
export class HomeCourtComponent implements OnInit {
  subjects: Subject[] = [];

  constructor(private userService: UserService, private subjectService: SubjectService) { }

  ngOnInit() {
    this.loadSubjects();
  }

  getRole() {
    return this.userService.currentUser?.user.userRole;
  }

  loadSubjects() {
    this.subjectService.getAllSubjects().subscribe(
      subjects => {
        this.subjects = subjects.filter(subject => subject.status === 'WAITING');
      },
      error => {
        console.error('Error loading subjects:', error);
      }
    );
  }
}
