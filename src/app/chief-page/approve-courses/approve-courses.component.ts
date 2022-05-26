import { Component, Input, OnInit } from '@angular/core';
import { ObjectId } from 'mongodb';
import { ChiefService } from 'src/app/Services/chief.service';
import { Course } from 'src/models/course.model';

@Component({
  selector: 'app-approve-courses',
  templateUrl: './approve-courses.component.html',
  styleUrls: ['./approve-courses.component.css']
})
export class ApproveCoursesComponent implements OnInit {
  @Input() courses!: Course[];
  constructor(private service: ChiefService) { }

  ngOnInit(): void {
  }

  public approveCourse(courseId: ObjectId){
    this.service.approveCourse(courseId).subscribe(res => {alert(res)}, err => {alert(err)})
  }
}
