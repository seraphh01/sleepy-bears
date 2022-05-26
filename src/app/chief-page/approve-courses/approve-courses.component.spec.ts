import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApproveCoursesComponent } from './approve-courses.component';

describe('ApproveCoursesComponent', () => {
  let component: ApproveCoursesComponent;
  let fixture: ComponentFixture<ApproveCoursesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApproveCoursesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApproveCoursesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
