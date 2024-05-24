import { Component } from '@angular/core';
import { CategoryPerson } from 'src/app/models/statisics/categoryPerson';

@Component({
  selector: 'app-request',
  templateUrl: './request.component.html',
  styleUrls: ['./request.component.css']
})
export class RequestComponent {
  selectedCategoryPerson: CategoryPerson | '' = '';
  categoriesPerson: CategoryPerson[] = Object.values(CategoryPerson);
}
