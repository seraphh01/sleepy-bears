import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Course } from 'src/models/course.model';
import { Student } from 'src/models/Student';
import { UserModel } from '../models/user.model';
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

  constructor(private service: StudentService ) {

  }

  async ngOnInit(){
    let year = 1;
    if(this.student!.group!)
    year = this.student.group.year;
    this.courses = await this.service.getMandatoryCourses(year);
    this.optionalCourses = await this.service.getOptionalCourses(this.student!.username);
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

  // addCourse(){
  //   let ok = 0;
  //   for(let i=0; i<this.courseList.length; ++i){
  //     if(this.courseList[i] == this.courses[this.slideIndex]){
  //       ok = 1;
  //     }
  //   }
  //   if(ok == 0){
  //     this.allMyCourses+=  this.courses[this.slideIndex] + " | ";
  //     this.courseList.push( this.courses[this.slideIndex]);
  //   }
  // }
}
