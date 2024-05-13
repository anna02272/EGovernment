import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomePoliceComponent } from './home-police.component';

describe('HomePoliceComponent', () => {
  let component: HomePoliceComponent;
  let fixture: ComponentFixture<HomePoliceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomePoliceComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomePoliceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
