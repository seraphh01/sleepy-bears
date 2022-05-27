import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SortStudentByGroupComponent } from './sort-student-by-group.component';

describe('SortStudentByGroupComponent', () => {
  let component: SortStudentByGroupComponent;
  let fixture: ComponentFixture<SortStudentByGroupComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SortStudentByGroupComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SortStudentByGroupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
