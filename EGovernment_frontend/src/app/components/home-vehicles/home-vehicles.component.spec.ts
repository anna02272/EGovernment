import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeVehiclesComponent } from './home-vehicles.component';

describe('HomeVehiclesComponent', () => {
  let component: HomeVehiclesComponent;
  let fixture: ComponentFixture<HomeVehiclesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomeVehiclesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomeVehiclesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
