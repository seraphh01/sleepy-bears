import { Component, Input, OnInit } from '@angular/core';
import { ObjectId } from 'mongodb';
import { StudentService } from 'src/app/Services/student.service';
import { UserService } from 'src/app/Services/user.service';
import { Course } from 'src/models/course.model';

@Component({
  selector: 'app-view-enrollments',
  templateUrl: './view-enrollments.component.html',
  styleUrls: ['./view-enrollments.component.css']
})
export class ViewEnrollmentsComponent implements OnInit {
  @Input() courses!: Course[];
  @Input() courseType: string = "Mandatory";

  constructor(private service: StudentService, public userService: UserService) { }

  ngOnInit(): void {
  }

  unenroll(courseId: ObjectId){
    this.service.unenroll(courseId).subscribe(
      res => {
      confirm(res['success']);
      window.location.reload();
    }, err => {
        alert(err['error'])
      });
  }

}
