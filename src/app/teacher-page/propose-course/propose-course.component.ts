import { Component, Input, OnInit } from '@angular/core';
import {MatInputModule} from '@angular/material/input';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Course } from 'src/models/course.model';
import { TeacherService } from 'src/app/Services/teacher.service';
import { Amount } from 'src/models/amount.model';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-propose-course',
  templateUrl: './propose-course.component.html',
  styleUrls: ['./propose-course.component.css']
})
export class ProposeCourseComponent implements OnInit {
  @Input() user!: UserModel;
  formGroup: FormGroup;
  constructor(private teacherService: TeacherService) {
    this.formGroup = new FormGroup({
      name: new FormControl("Course Name", Validators.required),
      yearofstudy: new FormControl(1, [Validators.required, Validators.min(1)]),
      maxStudents: new FormControl(100, [Validators.required, Validators.max(200), Validators.min(20)]),
      credits: new FormControl(5, [Validators.required, Validators.min(1)]),
      semester: new FormControl(1, [Validators.required, Validators.min(1)])
    })
   }

  ngOnInit(): void {
  }

  public proposeCourse(){
    let course: Course = this.formGroup.getRawValue();
    console.log(course);
    course.maxamount = {max: this.formGroup.get("maxStudents")!.value} as Amount;
    course.coursetype = 'OPTIONAL';

    try {
      this.teacherService.proposeCourse(course, this.user!.username.toString()).subscribe(res => {
        alert(res);
        window.location.reload();
      }, err => alert(err.message));
    }catch(error: any){
      alert(error);
    }
  }
}
