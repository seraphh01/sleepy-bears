import { Component, Input, OnInit } from '@angular/core';
import { Course } from 'src/models/course.model';
import { UserModel } from 'src/models/user.model';
import { TeacherService } from '../Services/teacher.service';

@Component({
  selector: 'app-teacher-page',
  templateUrl: './teacher-page.component.html',
  styleUrls: ['./teacher-page.component.css']
})
export class TeacherPageComponent implements OnInit {
  @Input() user!: UserModel;
  proposedCourses!: Course[];
  courses!: Course[];
  constructor(private service: TeacherService) { }

  async ngOnInit() {
    this.service.getProposedCoursesByTeacher(this.user!.username).subscribe(res => {
      if(typeof res === 'string'){
        this.proposedCourses = new Array<Course>();
        return;
      }
      
      this.proposedCourses = res;
    }, err => {console.log(err)});



    this.service.getCoursesByTeacher(this.user!.username).subscribe(res => {
      if(typeof res === 'string'){
        this.courses = new Array<Course>();
        return;
      }
      
      this.courses = res;
    });
  }

}
