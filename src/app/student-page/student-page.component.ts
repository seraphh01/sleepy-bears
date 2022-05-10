import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-student-page',
  templateUrl: './student-page.component.html',
  styleUrls: ['./student-page.component.css']
})
export class StudentPageComponent implements OnInit {

  courses = ["Math","Algebra"];
  currentCourse = "";
  slideIndex = 0;
  allMyCourses = "";
  courseList = Array<string>();

  
  constructor(private http:HttpClient) {
  }

  showCourses(){
      this.currentCourse = this.courses[this.slideIndex];
      let nextButton = <HTMLElement>document.getElementById('next');
      nextButton.style.display = 'block';
      let prevButton = <HTMLElement>document.getElementById('prev');
      prevButton.style.display = 'block';
      let addButton = <HTMLElement>document.getElementById('add');
      addButton.style.display = 'block';
  }

  nextCourse(){
    this.slideIndex ++;
    this.showCourses();
  }

  prevCourse(){
    this.slideIndex --;
    this.showCourses();
  }

  addCourse(){
    let ok = 0;
    for(let i=0; i<this.courseList.length; ++i){
      if(this.courseList[i] == this.courses[this.slideIndex]){
        ok = 1;
      }
    }
    if(ok == 0){
      this.allMyCourses+=  this.courses[this.slideIndex] + " | ";
      this.courseList.push( this.courses[this.slideIndex]);
    }
  }

  ngOnInit(): void {
  }

}
