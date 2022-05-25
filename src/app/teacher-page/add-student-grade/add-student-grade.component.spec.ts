import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddStudentGradeComponent } from './add-student-grade.component';

describe('AddStudentGradeComponent', () => {
  let component: AddStudentGradeComponent;
  let fixture: ComponentFixture<AddStudentGradeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddStudentGradeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddStudentGradeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
