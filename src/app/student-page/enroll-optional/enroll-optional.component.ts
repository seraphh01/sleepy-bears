import { Component, Input, OnInit } from '@angular/core';
import { ObjectId } from 'mongodb';
import { skip } from 'rxjs';
import { StudentService } from 'src/app/Services/student.service';
import { TeacherService } from 'src/app/Services/teacher.service';
import { Course } from 'src/models/course.model';

@Component({
  selector: 'app-enroll-optional',
  templateUrl: './enroll-optional.component.html',
  styleUrls: ['./enroll-optional.component.css']
})
export class EnrollOptionalComponent implements OnInit {
  courseList!:Course[];

  @Input() studentEnrollments!: Course[];

  constructor(private studentService: StudentService, private teacherService: TeacherService) { 

  }
  
  enroll(courseID:ObjectId){
    this.studentService.enroll(courseID.toString());
  }

  async ngOnInit() {
    this.courseList = await this.teacherService.getProposedCourses();
    let newList: Course[] = [];
    this.courseList.forEach(course => {
      if(!this.studentEnrollments.find(c => c.ID == course.ID))
        newList.push(course);
    });
    this.courseList = newList;
  }

}
