import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisteredVehiclesComponent } from './registered-vehicles.component';

describe('RegisteredVehiclesComponent', () => {
  let component: RegisteredVehiclesComponent;
  let fixture: ComponentFixture<RegisteredVehiclesComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RegisteredVehiclesComponent]
    });
    fixture = TestBed.createComponent(RegisteredVehiclesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
