import { Component, Input, OnInit } from '@angular/core';
import { Course } from 'src/models/course.model';

@Component({
  selector: 'app-view-proposed',
  templateUrl: './view-proposed.component.html',
  styleUrls: ['./view-proposed.component.css']
})
export class ViewProposedComponent implements OnInit {
  @Input() proposedCourses!: Course[];
  constructor() { }

  ngOnInit(): void {
  }

}
