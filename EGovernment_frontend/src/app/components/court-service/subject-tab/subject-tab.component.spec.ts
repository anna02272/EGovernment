import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubjectTabComponent } from './subject-tab.component';

describe('SubjectTabComponent', () => {
  let component: SubjectTabComponent;
  let fixture: ComponentFixture<SubjectTabComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SubjectTabComponent]
    });
    fixture = TestBed.createComponent(SubjectTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
