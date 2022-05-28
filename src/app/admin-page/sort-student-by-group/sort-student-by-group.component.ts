import { Component, OnInit } from '@angular/core';
import { timeStamp } from 'console';
import { AdminService } from 'src/app/Services/admin.service';
import { UserService } from 'src/app/Services/user.service';
import { Group } from 'src/models/group.model';
import { StudentGrade } from 'src/models/studentGrade.model';
import { UserModel } from 'src/models/user.model';
import { threadId } from 'worker_threads';

@Component({
  selector: 'app-sort-student-by-group',
  templateUrl: './sort-student-by-group.component.html',
  styleUrls: ['./sort-student-by-group.component.css']
})
export class SortStudentByGroupComponent implements OnInit {
  students!: UserModel[];
  groups!: Group[];
  grades!: number[];

  studentGrades!: StudentGrade[];

  constructor(private adminService: AdminService, private userService: UserService) { }

  ngOnInit(): void {
    this.getPerformance();
    this.userService.getGroups().subscribe(res => {
      if(typeof res === 'string')
        return
      this.groups = res;
    })
  }

  public getPerformance(){
    this.adminService.getGroupsPerformane().subscribe(res => {
      this.grades = res['averageGrade']
      this.students = res['students']

      this.studentGrades = [];

      var i: number;
      for(i =0;i< this.grades.length;i++){
        this.studentGrades.push({
          student : this.students[i],
          grades : [this.grades[i]]
        } as StudentGrade)
      }

    }, err => alert(err))
  }
  
  public getPerformanceByGroup(group_number: number){
    this.adminService.getGroupsPerformane().subscribe(res => {
      this.grades = res['averageGrade']
      this.students = res['students']

      this.studentGrades = [];

      var i: number;
      for(i =0;i< this.grades.length;i++){

        if(this.students[i].group.number != group_number)
          continue
        this.studentGrades.push({
          student : this.students[i],
          grades : [this.grades[i]]
        } as StudentGrade)
      }

    }, err => alert(err))
  }

  public viewGroup(e: any){
    let group_number = e.target.value;
    this.getPerformanceByGroup(group_number)
  }
}
