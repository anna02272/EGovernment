import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router, Params } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service'; 
import { Subject,  lastValueFrom } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { UserService } from 'src/app/services/user.service';
import { UserRole } from 'src/app/models/userRole';
import { User } from 'src/app/models/user';

interface DisplayMessage {
  msgType: string;
  msgBody: string;
}

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit, OnDestroy {
  form: FormGroup = new FormGroup({});
  submitted = false;
  notification: DisplayMessage = {} as DisplayMessage;
  returnUrl = '';
  private ngUnsubscribe: Subject<void> = new Subject<void>();

  constructor(
    private userService: UserService,
    private authService: AuthService,
    private router: Router,
    private route: ActivatedRoute,
    private formBuilder: FormBuilder,
  ) {}

  ngOnInit() {
    this.route.params
      .pipe(takeUntil(this.ngUnsubscribe))
      .subscribe((params: Params) => {
        this.notification = params as DisplayMessage || { msgType: '', msgBody: '' };
      });
  
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';

    this.form = this.formBuilder.group({
      password: ['', [Validators.required, Validators.minLength(8), Validators.maxLength(32), Validators.pattern(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&]{8,}$/)]],
      email: ['', [Validators.required, Validators.email, Validators.minLength(6), Validators.maxLength(64)]],
    });
  }
  ngOnDestroy() {
    this.ngUnsubscribe.next();
    this.ngUnsubscribe.complete();
  }

  onSubmit() {
    this.notification = { msgType: '', msgBody: '' };
    this.submitted = true;
  
    this.authService.login(this.form.value).subscribe({
      next: () => {
        this.userService.getMyInfo().subscribe(() => { 
          const role = this.userService.currentUser.user.userRole;
          console.log(role);
          if (role === UserRole.Citizen){
            this.router.navigate(['/pocetna']);
          } else if (role === UserRole.Employee){
            this.router.navigate(['/zavodZaStatistiku']);
          } else if (role === UserRole.Policeman){
            this.router.navigate(['/mupVozila']);
          } else if (role === UserRole.TrafficPoliceman){
            this.router.navigate(['/saobracajnaPolicija']);
          } else if (role === UserRole.Judge){
            this.router.navigate(['/prekrsajniSud']);   
          }
        });
      },
      error: (error) => {
        this.submitted = false;
  
        if (error.statusText === 'Unknown Error') {
          this.notification = {
            msgType: 'error',
            msgBody: 'Usluga autorizacije nije dostupna.'
          };
        } else {
          this.notification = {
            msgType: 'error',
            msgBody: 'Nepravilno korisničko ime ili lozinka.'
          };
        }
      }
    });
  }  
}
