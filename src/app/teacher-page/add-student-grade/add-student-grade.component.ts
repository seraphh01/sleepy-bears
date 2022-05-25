import { Component, Input, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { StudentService } from 'src/app/Services/student.service';
import { TeacherService } from 'src/app/Services/teacher.service';
import { Course } from 'src/models/course.model';
import { Student } from 'src/models/Student';

@Component({
  selector: 'app-add-student-grade',
  templateUrl: './add-student-grade.component.html',
  styleUrls: ['./add-student-grade.component.css']
})
export class AddStudentGradeComponent implements OnInit {
  @Input() proposedCourses!: Course[];
  students:Student[] = [];
  grade!: number;

  constructor(private teacherService: TeacherService) { 
  }

  async ngOnInit() {
    let courseId = this.proposedCourses[0].ID;
    this.students = await this.teacherService.getStudentsByCourse(courseId.toString());
  }

  gradeStudent( courseName: string){
    let studentName = "Andrei";
    this.teacherService.gradeStudent(studentName,courseName,this.grade);
  }

}
