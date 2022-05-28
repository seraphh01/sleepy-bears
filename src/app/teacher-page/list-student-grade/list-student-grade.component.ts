import { Component, Input, OnInit } from '@angular/core';
import { ObjectID } from 'bson';
import { TeacherService } from 'src/app/Services/teacher.service';
import { Course } from 'src/models/course.model';
import { Grade } from 'src/models/Grade.model';
import { Student } from 'src/models/Student';
import { StudentGrade } from 'src/models/studentGrade.model';

@Component({
  selector: 'app-list-student-grade',
  templateUrl: './list-student-grade.component.html',
  styleUrls: ['./list-student-grade.component.css']
})
export class ListStudentGradeComponent implements OnInit {
  @Input() mandatoryCourses!: Course[];
  @Input() proposedCourses!: Course[];
  students: Student[] = [];
  courseGrades: Grade[][] = [];
  courses: Course[] = [];
  courseStudentGrades!: Map<Course,StudentGrade[]>;

  constructor(private teacherService: TeacherService) { 
    this.courseStudentGrades = new Map<Course,StudentGrade[]>();
  }

  ngOnInit(): void {
      this.courses = this.mandatoryCourses;
      this.courses.forEach(element => {
        this.getGrades(element)
      });
  }

  getGrades(course: Course){
    this.teacherService.getStudentsAtCourse(course.ID).subscribe((res: any) => {
      this.courseStudentGrades.set(course,new Array<StudentGrade>());

      let students = res['students'];
      let studentGrades = res['grades'];
      var i : number;

      if(students == null || studentGrades == null)
        return;

      for(i=0;i<students.length;i++){
        let grades = Array<number>();
        studentGrades[i].forEach(
          (grade:Grade) => grades.push(grade.grade)
        )
        this.courseStudentGrades.get(course)?.push({student: students[i],grades: grades});
      }
   });
      console.log(this.courseStudentGrades);
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
      return [];
    }
    return this.courseGrades[index];
  }

}
