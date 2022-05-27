import { Component, Input, OnInit } from '@angular/core';
import { ObjectId } from 'mongodb';
import { skip } from 'rxjs';
import { StudentService } from 'src/app/Services/student.service';
import { TeacherService } from 'src/app/Services/teacher.service';
import { UserService } from 'src/app/Services/user.service';
import { Course } from 'src/models/course.model';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-enroll-optional',
  templateUrl: './enroll-optional.component.html',
  styleUrls: ['./enroll-optional.component.css']
})
export class EnrollOptionalComponent implements OnInit {
  courseList!:Course[];

  @Input() studentEnrollments!: Course[];
  @Input() student!: UserModel

  constructor(private studentService: StudentService, private teacherService: TeacherService, public userService: UserService) { 

  }
  
  enroll(course:Course){
    this.studentService.enroll(course.ID).subscribe(res => {
      alert(`Succesfully enrolled to ${course?.name}!`);
      window.location.reload();
    });
  }

  async ngOnInit() {
   this.teacherService.getProposedCoursesByAcademicYear(this.student.group!.year! | 1).subscribe(res => {
    this.courseList = res;

      if(typeof this.studentEnrollments === 'string')
      return;

    let newList: Course[] = [];
    this.courseList.forEach(course => {
      if(!this.studentEnrollments.find(c => c.name == course.name))
        newList.push(course);
    });
    this.courseList = newList;

    console.log(this.studentEnrollments);
    console.log(this.courseList);

    if(this.studentEnrollments.length == 0)
    if(this.userService.canSign())
      alert('Make sure to enroll in optional courses!')
    else if(this.userService.inAcademicYear() && this.userService.canSign() == false){
        // siging period is due time but the student hasn't yet enrolled in any course, 
        // auto enroll him in first two courses
        alert("You haven't enrolled in any course and the signing period is due. We will auto enroll you in two courses")
        var i: number;
        let messages : string[] = [];
        let mi: number = 0;
        for(i =0; i < 2 && i < this.courseList.length; i++){
          messages.push(`You were enrolled in ${this.courseList[i].name}`)
          this.studentService.enroll(this.courseList[i].ID).subscribe(res => {
            confirm(messages[mi++]);

            if( mi >= i)
              window.location.reload();

          }, err => alert(err));
          
        }
      }
   }, err => {

   });
  }

  private findCourse(courseId: ObjectId){
    for(let course of this.courseList){
      if(course.ID == courseId)
        return course;
    }

    return null;
  }
}
