import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateCarAccidentComponent } from './create-car-accident.component';

describe('CreateCarAccidentComponent', () => {
  let component: CreateCarAccidentComponent;
  let fixture: ComponentFixture<CreateCarAccidentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateCarAccidentComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateCarAccidentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});