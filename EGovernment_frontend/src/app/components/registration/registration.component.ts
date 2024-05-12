import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators ,AbstractControl} from '@angular/forms';
import { ActivatedRoute, Router, Params } from '@angular/router';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { UserRole } from 'src/app/models/userRole';
import { AuthService } from 'src/app/services/auth.service'; 
import { UserService } from 'src/app/services/user.service';

interface DisplayMessage {
  msgType: string;
  msgBody: string;
}

@Component({
  selector: 'app-register',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent {
  password: string = '';
  personalInfoForm: FormGroup = new FormGroup({});
  submitted = false;
  name : any;
  new : any;
  notification: DisplayMessage = {} as DisplayMessage;
  returnUrl = '';
  private ngUnsubscribe: Subject<void> = new Subject<void>();

  constructor(
    private authService: AuthService,
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute,
    private formBuilder: FormBuilder
  ) {
    this.personalInfoForm = this.formBuilder.group({
      username: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(32)]],
      password: ['', [Validators.required, Validators.minLength(8), Validators.maxLength(32), Validators.pattern(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&]{8,}$/)]],
      email: ['', [Validators.required, Validators.email, Validators.minLength(6), Validators.maxLength(64)]],
      name: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(32)]],
      lastname: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(32)]],
    });

  }
  ngOnInit() {
    this.route.params
      .pipe(takeUntil(this.ngUnsubscribe))
      .subscribe((params: Params) => {
        this.notification = params as DisplayMessage || { msgType: '', msgBody: '' };
      });

    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
  }

  onSubmit() {
    this.notification = { msgType: '', msgBody: '' };
    this.submitted = true;
    this.name= this.personalInfoForm.get('name')
    this.new = {}
    this.new.username= this.personalInfoForm.get('username')?.value
    this.new.password= this.personalInfoForm.get('password')?.value
    this.new.email= this.personalInfoForm.get('email')?.value
    this.new.name= this.personalInfoForm.get('name')?.value
    this.new.lastname= this.personalInfoForm.get('lastname')?.value

    this.authService.register(this.new).subscribe({
      next: () => {
      this.authService.login(this.new).subscribe({
        next: () => {
          this.submitted = true;
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
        if (error.status === 409) {
          if (error.error.message === 'Username already exists') {
            this.notification = { msgType: 'error', msgBody: ' Korisničko ime već postoji' };
          } else if (error.error.message === 'Invalid password format') {
            this.notification = { msgType: 'error', msgBody: 'Nevažeći format lozinke' };
        } else if (error.error.message === 'Email already exists') {
          this.notification = { msgType: 'error', msgBody: 'Email već postoji' };
        }
          else {
            this.notification = { msgType: 'error', msgBody: 'Registracija nije uspela. Molimo Vas, pokušajte ponovo.' };
          }
        } else {
          this.notification = { msgType: 'error', msgBody: 'Registracija nije uspela. Molimo Vas, pokušajte ponovo.' };
        }
        this.submitted = false;
      }
      });
    }
   });
  }  
  

  ngOnDestroy() {
    this.ngUnsubscribe.next();
    this.ngUnsubscribe.complete();
  }
}