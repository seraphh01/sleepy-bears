import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListStudentGradeComponent } from './list-student-grade.component';

describe('ListStudentGradeComponent', () => {
  let component: ListStudentGradeComponent;
  let fixture: ComponentFixture<ListStudentGradeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListStudentGradeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ListStudentGradeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
