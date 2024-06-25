import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CarAccidentDetailsComponent } from './car-accident-details.component';

describe('CarAccidentDetailsComponent', () => {
  let component: CarAccidentDetailsComponent;
  let fixture: ComponentFixture<CarAccidentDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CarAccidentDetailsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CarAccidentDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});