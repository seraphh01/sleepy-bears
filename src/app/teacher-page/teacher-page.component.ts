import { Component, Input, OnInit } from '@angular/core';
import { Course } from 'src/models/course.model';
import { Group } from 'src/models/group.model';
import { UserModel } from 'src/models/user.model';
import { TeacherService } from '../Services/teacher.service';
import { UserService } from '../Services/user.service';

@Component({
  selector: 'app-teacher-page',
  templateUrl: './teacher-page.component.html',
  styleUrls: ['./teacher-page.component.css']
})
export class TeacherPageComponent implements OnInit {
  @Input() user!: UserModel;
  groups!: Group[];
  proposedCourses!: Course[];
  courses!: Course[];
  constructor(private service: TeacherService, private userService: UserService) { }

  ngOnInit() {
    this.service.getProposedCoursesByTeacher(this.user!.username).subscribe(res => {
      if(typeof res === 'string'){
        this.proposedCourses = new Array<Course>();
        return;
      }
      
      this.proposedCourses = res;
    }, err => {console.log(err)});

    this.userService.getGroups().subscribe( res => {
      if (typeof res === 'string')
        return;
      this.groups = res;
    }, err => console.error(err))

    this.service.getCoursesByTeacher(this.user!.username).subscribe((res: Course[]) => {
      if(typeof res === 'string'){
        this.courses = new Array<Course>();
        return;
      }
      
      this.courses = res;
    });
  }

}
