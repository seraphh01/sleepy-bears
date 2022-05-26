import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Course } from 'src/models/course.model';
import { Grade } from 'src/models/Grade.model';
import { Student } from 'src/models/Student';
import { UserModel } from 'src/models/user.model';
import { StudentService } from '../Services/student.service';

@Component({
  selector: 'app-student-page',
  templateUrl: './student-page.component.html',
  styleUrls: ['./student-page.component.css']
})
export class StudentPageComponent implements OnInit {
  @Input() student!: UserModel;

  show = false;
  courses!:Course[];
  optionalCourses!: Course[];
  currentCourse!: Course;
  slideIndex = 0;
  allMyCourses = "";
  courseList = Array<string>();
  presentGrades = false;
  grades:Grade[] = [];

  constructor(private service: StudentService ) {

  }

  async ngOnInit(){
    let year = 1;
    if(this.student!.group!)
    year = this.student.group.year;
    this.courses = await this.service.getMandatoryCourses(year);
    this.service.getOptionalCourses(this.student!.username).subscribe(res => {
      if(typeof res === 'string')
        this.optionalCourses = new Array<Course>();
      else
      this.optionalCourses = res;
    }, err => {
      console.log(err);
      this.optionalCourses = new Array<Course>();
    });
  }
  
  showCourses(){
      this.currentCourse = this.courses[this.slideIndex];
      this.show = true;
  }

  nextCourse(){
    this.slideIndex ++;
    this.showCourses();
  }

  prevCourse(){
    this.slideIndex --;
    this.showCourses();
  }

  showGrades(){
    this.presentGrades = true;
  }

}
