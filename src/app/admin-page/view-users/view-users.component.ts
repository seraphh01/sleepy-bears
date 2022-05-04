import { Component, Input, OnInit } from '@angular/core';
import { Student } from 'src/models/Student';

@Component({
  selector: 'app-view-users',
  templateUrl: './view-users.component.html',
  styleUrls: ['./view-users.component.css']
})
export class ViewUsersComponent implements OnInit {
  @Input() students: Student[] = [];
  constructor() { }

  ngOnInit(): void {

  }

}
