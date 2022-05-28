import { Component, Input, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ObjectId } from 'mongodb';
import { StudentService } from 'src/app/Services/student.service';
import { TeacherService } from 'src/app/Services/teacher.service';
import { UserService } from 'src/app/Services/user.service';
import { Course } from 'src/models/course.model';
import { Group } from 'src/models/group.model';
import { Student } from 'src/models/Student';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-add-student-grade',
  templateUrl: './add-student-grade.component.html',
  styleUrls: ['./add-student-grade.component.css']
})
export class AddStudentGradeComponent implements OnInit {
  @Input() proposedCourses!: Course[];
  @Input() courses!: Course[];
  @Input() groups!: Group[];

  group_number!: number;

  myCourses: Course[] = [];
  studentsOfCourse: Map<ObjectId, UserModel[]> = new Map<ObjectId, UserModel[]>();
  students:Student[] = [];
  grade!: number;

  constructor(private teacherService: TeacherService, public userService: UserService) { 
  }

  ngOnInit() {
    this.courses.forEach(c => this.myCourses.push(c))
    this.proposedCourses.forEach(c => this.myCourses.push(c));
    for(let course of this.myCourses){
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
    this.teacherService.getStudentsByCourse(courseId).subscribe((res: any) => {
      
      if(typeof res === 'string'){
        return;
      }
      for(let student of res)
        this.studentsOfCourse.get(courseId)?.push(student);

      console.log(this.studentsOfCourse)
    }, error => {
      alert(error)
    });
  }

  public changeGroup(e: any){
    this.group_number = e.target.value;
  }

  public getCurrentGroupNumber(){
    if(this.group_number != null)
      return this.group_number;
    return this.groups[0].number;
  }

}
