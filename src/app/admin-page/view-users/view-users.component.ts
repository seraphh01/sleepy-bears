import { identifierName } from '@angular/compiler';
import { Component, Input, OnInit } from '@angular/core';
import { ObjectId } from 'mongodb';
import { AdminService } from 'src/app/Services/admin.service';
import { UserService } from 'src/app/Services/user.service';
import { Group } from 'src/models/group.model';
import { Student } from 'src/models/Student';
import { UserModel } from 'src/models/user.model';
import { threadId } from 'worker_threads';

@Component({
  selector: 'app-view-users',
  templateUrl: './view-users.component.html',
  styleUrls: ['./view-users.component.css']
})
export class ViewUsersComponent implements OnInit {
  @Input() students!: UserModel[];

  public groups!: Group[];

  constructor(private adminService: AdminService, private userService: UserService) { }

  ngOnInit(): void {
    this.adminService.addStudentToGroup(915, "seraphh0").subscribe()
    this.userService.getGroups().subscribe(res => {
      if(typeof res === 'string')
        return;
      this.groups = res;

    }, err => console.error(err));
  }

  public changeGroup(user_name: String, e: any){
    let group_number = e.target.value;

    this.adminService.addStudentToGroup(group_number, user_name.toString()).subscribe(res => {
      this.students[this.students.indexOf(this.students.find(s => s.username == user_name)!)].group = res.group;
      // this.students = this.students.filter(s => s.username != user_name);
      // this.students.unshift(res);
    }, err => alert(err));
  }

  public async viewGroup(e: any){
    let group: any = e.target.value;

    if(group == 'All'){
      this.students = await this.adminService.getStudents();
      return;
    }

    this.adminService.getStudentsByGroup(group).subscribe(res => {
      if(typeof res === 'string'){
        this.students = [];
        return;
      }
      
      this.students = res;
    },err => alert(err))
  }
}
