import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeCourtComponent } from './home-court.component';

describe('HomeCourtComponent', () => {
  let component: HomeCourtComponent;
  let fixture: ComponentFixture<HomeCourtComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomeCourtComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomeCourtComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
