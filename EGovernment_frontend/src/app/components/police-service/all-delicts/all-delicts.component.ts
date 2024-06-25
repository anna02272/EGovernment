import { Component, OnInit } from '@angular/core';
import { Delict } from 'src/app/models/traffic-police/delict';
import { DelictService } from 'src/app/services/traffic-police/delictService';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-all-delicts',
  templateUrl: './all-delicts.component.html',
  styleUrls: ['./all-delicts.component.css']
})
export class AllDelictsComponent implements OnInit {
  delicts: Delict[] = [];

  constructor(
    private delictService: DelictService,
    private userService: UserService
  ) { }

  ngOnInit() {
    this.loadDelicts();
  }

  loadDelicts() {
    const role = this.userService.currentUser?.user.userRole;

    if (role === 'TrafficPoliceman') {
      this.delictService.getAllDelictsByPolicemanID().subscribe(
        (delicts) => {
          this.delicts = delicts;
        },
        (error) => {
          console.error('Error loading delicts:', error);
        }
      );
    } else if (role === 'Citizen') {
      this.delictService.getAllDelictsByDriver().subscribe(
        (delicts) => {
          this.delicts = delicts;
        },
        (error) => {
          console.error('Error loading delicts:', error);
        }
      );
    }
  }
}