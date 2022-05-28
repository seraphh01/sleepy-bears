import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ChiefService } from 'src/app/Services/chief.service';
import { UserService } from 'src/app/Services/user.service';
import { AcademicYear } from 'src/models/academic-year.model';
import { Course } from 'src/models/course.model';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-add-course',
  templateUrl: './add-course.component.html',
  styleUrls: ['./add-course.component.css']
})
export class AddCourseComponent implements OnInit {
  @Input() user!: UserModel;

  public teachers!: UserModel[];

  courseForm!: FormGroup;

  constructor(private service: ChiefService, private userService: UserService) { }

  ngOnInit(): void {
    this.courseForm = new FormGroup({
      name: new FormControl('Name', [Validators.required, Validators.minLength(2)]),
      yearofstudy: new FormControl(1, [Validators.required]),
      credits: new FormControl(5, [Validators.required, Validators.min(1)]),
      semester: new FormControl(1, [Validators.required, Validators.min(1)]),
      teacher: new FormControl("dadi1234", [Validators.required])
    });

    this.teachers = [];

    this.userService.getUsers('TEACHER').subscribe((res: UserModel[]) => {
      if(typeof res === 'string')
        return;
        res.forEach(t => this.teachers.push(t))

    }, err => {alert(err)})

    this.userService.getUsers('CHIEF').subscribe((res: UserModel[]) => {
      if(typeof res === 'string')
        return;
        res.forEach(t => this.teachers.push(t))

    }, err => {alert(err)})

  }

  public addCourse(){
    let course: Course = this.courseForm.getRawValue();
    let teacher = this.courseForm.get('teacher')?.value || this.user.username;
    course.academic_year = {
      startdate : new Date(2021, 10, 5).toDateString(),
      enddate : new Date(2001, 11, 6).toDateString()
    } as AcademicYear;
    this.service.addMandatoryCourse(teacher, course).subscribe(
      _ => {
        confirm("Course was added successfully!")
      },
       err => {alert(err)});
  }
}
