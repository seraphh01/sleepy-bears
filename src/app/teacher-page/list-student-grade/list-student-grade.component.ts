import { Component, Input, OnInit } from '@angular/core';
import { TeacherService } from 'src/app/Services/teacher.service';
import { Course } from 'src/models/course.model';
import { Grade } from 'src/models/Grade.model';

@Component({
  selector: 'app-list-student-grade',
  templateUrl: './list-student-grade.component.html',
  styleUrls: ['./list-student-grade.component.css']
})
export class ListStudentGradeComponent implements OnInit {
  @Input() proposedCourses!: Course[];
  constructor(private teacherService: TeacherService) { }

  ngOnInit(): void {

  }

  printGrades(){
  }

}
