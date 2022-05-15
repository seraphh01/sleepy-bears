import { Component, OnInit } from '@angular/core';
import {MatInputModule} from '@angular/material/input';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Course } from 'src/models/course.model';
import { TeacherService } from 'src/app/Services/teacher.service';
import { Amount } from 'src/models/amount.model';

@Component({
  selector: 'app-propose-course',
  templateUrl: './propose-course.component.html',
  styleUrls: ['./propose-course.component.css']
})
export class ProposeCourseComponent implements OnInit {
  formGroup: FormGroup;
  constructor(private teacherService: TeacherService) {
    this.formGroup = new FormGroup({
      name: new FormControl("Name", Validators.required),
      year: new FormControl(1, [Validators.required, Validators.min(1)]),
      maxStudents: new FormControl(100, [Validators.required, Validators.max(200)])
    })
   }

  ngOnInit(): void {
  }

  public proposeCourse(){
    let course: Course = {} as Course;
    course.courseType = "OPTIONAL";
    course.name = "name"
    course.year = 1
    course.maxAmount = {max: 40} as Amount;
    this.teacherService.proposeCourse(course);
    console.log(course);
  }
}
