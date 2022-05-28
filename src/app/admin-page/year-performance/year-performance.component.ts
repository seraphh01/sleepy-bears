import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AdminService } from 'src/app/Services/admin.service';
import { StudentGrade } from 'src/models/studentGrade.model';

@Component({
  selector: 'app-year-performance',
  templateUrl: './year-performance.component.html',
  styleUrls: ['./year-performance.component.css']
})
export class YearPerformanceComponent implements OnInit {
  public formGroup!: FormGroup;
  public studentGrades!: StudentGrade[];
  constructor(private adminService: AdminService) {
    this.formGroup = new FormGroup({
      year: new FormControl(1, [Validators.required])
    })
   }

  ngOnInit(): void {
  }

  public viewPerformance(){
    let year = this.formGroup.get("year")?.value || 1;
    this.adminService.getYearPerformance(year).subscribe(res => {
      console.log(res);
      let grades: number[] = res['averageGrade'];
      let students = res['students'];

      this.studentGrades = [];

      var i: number;
      for(i =0;i< grades.length;i++){
        this.studentGrades.push({
          student : students[i],
          grades : [grades[i]]
        } as StudentGrade)
      }

      console.log(this.studentGrades);
    }, err => alert(err))
  }

}
