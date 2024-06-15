import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatisticsHeaderComponent } from './statistics-header.component';

describe('StatisticsHeaderComponent', () => {
  let component: StatisticsHeaderComponent;
  let fixture: ComponentFixture<StatisticsHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ StatisticsHeaderComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StatisticsHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
