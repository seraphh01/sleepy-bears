import { Component, Input, OnInit } from '@angular/core';
import { Course } from 'src/models/course.model';
import { UserModel } from 'src/models/user.model';
import { UserService } from '../Services/user.service';

@Component({
  selector: 'app-chief-page',
  templateUrl: './chief-page.component.html',
  styleUrls: ['./chief-page.component.css']
})
export class ChiefPageComponent implements OnInit {
  @Input() user!: UserModel;

  optionalCourses! : Course[];

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.userService.getOptionalCourses().subscribe((res: any) => {
      this.optionalCourses = res;
    }, err => {alert(err)})
  }

}
