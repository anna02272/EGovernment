import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CategoryPerson } from 'src/app/models/statisics/categoryPerson';
import { Request } from 'src/app/models/statisics/request';
import { RequestService } from 'src/app/services/statistics/request.service';

@Component({
  selector: 'app-request',
  templateUrl: './request.component.html',
  styleUrls: ['./request.component.css']
})
export class RequestComponent  implements OnInit {
  requestForm: FormGroup;
  selectedCategoryPerson: CategoryPerson | '' = '';
  categoriesPerson: CategoryPerson[] = Object.values(CategoryPerson);

constructor(
  private fb: FormBuilder,
  private requestService: RequestService
) {
  this.requestForm = this.fb.group({
    name: ['', Validators.required],
    lastname: ['', Validators.required],
    email: ['', [Validators.required, Validators.email]],
    phone_number: ['', Validators.required],
    category: ['', Validators.required],
    question: ['', Validators.required]
  });
}

ngOnInit(): void {}

onSubmit(): void {
  if (this.requestForm.valid) {
    const newRequest: Request = this.requestForm.value;
    this.requestService.create(newRequest).subscribe(response => {
      console.log('Request created successfully', response);
    });
  }
}

onCancel(): void {
  this.requestForm.reset();
}
}