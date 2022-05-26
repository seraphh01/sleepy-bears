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
    this.proposedCourses = await this.service.getProposedCoursesByTeacher(this.user!.username);
    this.service.getCoursesByTeacher(this.user!.username).subscribe(res => {
      if(typeof res === 'string'){
        this.courses = new Array<Course>();
        return;
      }
      console.log(res);
      this.courses = res;
    });
  }

}
