import { Component, OnInit } from '@angular/core';

import { Hearing } from 'src/app/models/court/hearing';
import { HearingService } from 'src/app/services/court/hearing.service';
import { Router } from '@angular/router';
import { Subject } from 'src/app/models/court/subject';
import { SubjectService } from 'src/app/services/court/subject.service';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-home-court',
  templateUrl: './home-court.component.html',
  styleUrls: ['./home-court.component.css'],
})
export class HomeCourtComponent implements OnInit {
  subjects: Subject[] = [];
  hearings: Hearing[] = [];

  constructor(
    private userService: UserService,
    private subjectService: SubjectService,
    private hearingService: HearingService,
    private router: Router,
  ) {}

  ngOnInit() {
    this.loadSubjects();
    this.loadHearings();
  }

  getRole() {
    return this.userService.currentUser?.user.userRole;
  }

  loadSubjects() {
    this.subjectService.getAllSubjects().subscribe(
      (subjects) => {
        this.subjects = subjects.filter(
          (subject) => subject.status === 'WAITING',
          console.log(subjects),
        );
      },
      (error) => {
        console.error('Error loading subjects:', error);
      },
    );
  }

  loadHearings() {
    this.hearingService.getAllHearingsByJudge().subscribe(
      (hearings) => {
        console.log(hearings);
        this.hearings = hearings;
        this.loadSubjectsForHearings();
      },
      (error: any) => {
        console.error('Error loading hearings:', error);
      },
    );
  }
  loadSubjectsForHearings() {
    this.hearings.forEach((hearing) => {
      this.hearingService.getSubjectById(hearing.subject_id).subscribe(
        (subject) => {
          const hearingIndex = this.hearings.findIndex(
            (h) => h.id === hearing.id,
          );
          if (hearingIndex !== -1) {
            this.hearings[hearingIndex].subject = subject;
          }
        },
        (error) => {
          console.error(
            `Error loading subject for hearing ${hearing.id}:`,
            error,
          );
        },
      );
    });
  }
  navigateToHearingDetails(hearing: any) {
    console.log(hearing);
  }
  navigateToSubjectDetails(subjectId: string, id: any) {
    this.router.navigate(['/subjectTab', subjectId, { hearingId: id }]);
  }
}
