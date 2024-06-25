import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CarAccidentComponent } from './car-accident.component';

describe('CarAccidentComponent', () => {
  let component: CarAccidentComponent;
  let fixture: ComponentFixture<CarAccidentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CarAccidentComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CarAccidentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});