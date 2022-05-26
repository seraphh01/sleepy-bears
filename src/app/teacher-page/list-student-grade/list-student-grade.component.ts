import { Component, Input, OnInit } from '@angular/core';
import { TeacherService } from 'src/app/Services/teacher.service';
import { Course } from 'src/models/course.model';
import { Grade } from 'src/models/Grade.model';
import { Student } from 'src/models/Student';

@Component({
  selector: 'app-list-student-grade',
  templateUrl: './list-student-grade.component.html',
  styleUrls: ['./list-student-grade.component.css']
})
export class ListStudentGradeComponent implements OnInit {
  @Input() proposedCourses!: Course[];
  courseName!: string;
  showGrades = false;
  grades: Grade[] = [];
  students: Student[] = [];

  constructor(private teacherService: TeacherService) { }

  ngOnInit(): void {
    this.teacherService.getGradesAtCourse(this.proposedCourses[0].ID)
      .subscribe((res:any) => {
        let students = res['students'];
        let courseGrades = res['grades'];
        this.students = new Array<Student>();
        this.grades = new Array<Grade>();
        for(let i =0;i< students.length;i++){
          let student = students[i];
          let grades = courseGrades[i];
  
          if(grades == null || grades.length == 0){
            continue;
          }
          this.students.push(student);
          this.grades.push(grades);
        }
      });
  }

  getStudentGrade(studentName: string){
    var index = 0;
    for(let i=0; i<this.students.length; ++i){
        if(this.students[i].username == studentName){
            index = i;
            break;
        }
    }
    return this.grades[0];
  }
}
