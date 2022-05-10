import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProposeCourseComponent } from './propose-course.component';

describe('ProposeCourseComponent', () => {
  let component: ProposeCourseComponent;
  let fixture: ComponentFixture<ProposeCourseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProposeCourseComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProposeCourseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
