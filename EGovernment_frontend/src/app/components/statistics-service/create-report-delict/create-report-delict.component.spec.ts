import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateReportDelictComponent } from './create-report-delict.component';

describe('CreateReportDelictComponent', () => {
  let component: CreateReportDelictComponent;
  let fixture: ComponentFixture<CreateReportDelictComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateReportDelictComponent]
    });
    fixture = TestBed.createComponent(CreateReportDelictComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
