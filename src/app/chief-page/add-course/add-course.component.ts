import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ObjectId } from 'mongodb';
import { ChiefService } from 'src/app/Services/chief.service';
import { AcademicYear } from 'src/models/academic-year.model';
import { Course } from 'src/models/course.model';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-add-course',
  templateUrl: './add-course.component.html',
  styleUrls: ['./add-course.component.css']
})
export class AddCourseComponent implements OnInit {
  @Input() user!: UserModel;
  courseForm!: FormGroup;

  constructor(private service: ChiefService) { }

  ngOnInit(): void {
    this.courseForm = new FormGroup({
      name: new FormControl('Name', [Validators.required, Validators.minLength(2)]),
      year_of_study: new FormControl(1, [Validators.required]),
      credits: new FormControl(5, [Validators.required, Validators.min(1)]),
      semester: new FormControl(1, [Validators.required, Validators.min(1)]),
    });
  }

  public addCourse(){
    let course: Course = this.courseForm.getRawValue();
    
    course.academic_year = {
      startdate : new Date(2021, 10, 5).toDateString(),
      enddate : new Date(2001, 11, 6).toDateString()
    } as AcademicYear;
    this.service.addMandatoryCourse(this.user!.username.toString(), course).subscribe(
      _ => {
        confirm("Course was added successfully!")
      },
       err => {alert(err)});
  }
}
