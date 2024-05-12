import {Injectable} from '@angular/core';
import {ApiService} from './api.service';
import {ConfigService} from './config.service';
import {map} from 'rxjs/operators';
import { User } from '../models/user';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  currentUser: any;

  private currentUserSubject = new BehaviorSubject<User | null>(null);
  currentUser$ = this.currentUserSubject.asObservable();
  
  constructor(
    private apiService: ApiService,
    private config: ConfigService,
  ) {
  }

  getMyInfo() {
    return this.apiService.get(this.config.currentUser_url)
      .pipe(map(user => {
        this.currentUser = user;
        return user;
      }));
  } 
   
  setCurrentUser(user: User | null) {
    this.currentUserSubject.next(user);
  }
 
}
