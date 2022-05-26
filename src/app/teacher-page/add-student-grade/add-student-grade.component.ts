import { Component, Input, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ObjectId } from 'mongodb';
import { StudentService } from 'src/app/Services/student.service';
import { TeacherService } from 'src/app/Services/teacher.service';
import { Course } from 'src/models/course.model';
import { Student } from 'src/models/Student';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-add-student-grade',
  templateUrl: './add-student-grade.component.html',
  styleUrls: ['./add-student-grade.component.css']
})
export class AddStudentGradeComponent implements OnInit {
  @Input() proposedCourses!: Course[];
  studentsOfCourse: Map<ObjectId, UserModel[]> = new Map<ObjectId, UserModel[]>();
  students:Student[] = [];
  grade!: number;

  constructor(private teacherService: TeacherService) { 
  }

  async ngOnInit() {
    let courseId = this.proposedCourses[0].ID;

    for(let course of this.proposedCourses){
      this.studentsOfCourse.set(course.ID, new Array<UserModel>());
      this.getStudents(course.ID);
    }


    //this.students = await this.teacherService.getStudentsByCourse(courseId.toString());
  }

  gradeStudent( courseId: ObjectId, studentUserName: String){
    let grade = +(<HTMLInputElement>document.getElementById(courseId.toString() + studentUserName.toString()))?.value;
    if(grade < 1 || grade > 10)
      {
        alert("Grade is invalid, should be between 1 and 10!")
        return;
      }
    this.teacherService.gradeStudent(courseId, studentUserName, grade!).subscribe(res =>{console.log(res)}, err => {alert(err)});
  }

  getStudents(courseId: ObjectId){
    this.teacherService.getStudentsByCourse(courseId).subscribe((res: UserModel[]) => {
      for(let student of res)
        this.studentsOfCourse.get(courseId)?.push(student);
    }, error => {
      alert(error)
    });
  }

}
