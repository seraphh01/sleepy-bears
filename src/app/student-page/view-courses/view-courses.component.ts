import { Component, Input, OnInit } from '@angular/core';
import { Course } from 'src/models/course.model';

@Component({
  selector: 'app-view-courses',
  templateUrl: './view-courses.component.html',
  styleUrls: ['./view-courses.component.css']
})
export class ViewCoursesComponent implements OnInit {
  @Input() courses!: Course[];
  @Input() courseType: string = "Mandatory";

  constructor() { }

  ngOnInit(): void {
  }

}
