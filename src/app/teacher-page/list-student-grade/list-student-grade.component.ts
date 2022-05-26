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
  students: Student[] = [];
  courseGrades: Grade[] = [];

  constructor(private teacherService: TeacherService) { }

  ngOnInit(): void {
      this.teacherService.getStudentsAtCourse(this.proposedCourses[0].ID).subscribe((res: any) => {
      let students = res['students'];
      let studentGrades = res['grades'];
      var i : number;

      this.students = new Array<Student>();
      this.courseGrades = new Array<Grade>();

      if(students.length == 0){
        alert("No students enrolled");
         return;
      }

      for(i=0;i<students.length;i++){
        let student = students[i];
        let grade = studentGrades[i];
        this.students.push(student);
        this.courseGrades.push(grade);
      }
   });
  }

  getStudentGrade(studentName: string){
    let index = 0;
    for(let i=0; i<this.students.length; ++i){
      if(this.students[i].username == studentName){
        index = i;
        break;
      }
    }
    if(this.courseGrades[index] == null){
      return "no grade";
    }
    return this.courseGrades[index];
  }

}
