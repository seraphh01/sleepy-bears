import { Component, OnInit } from '@angular/core';
import { AdminService } from 'src/app/Services/admin.service';
import { UserService } from 'src/app/Services/user.service';
import { Group } from 'src/models/group.model';

@Component({
  selector: 'app-view-groups',
  templateUrl: './view-groups.component.html',
  styleUrls: ['./view-groups.component.css']
})
export class ViewGroupsComponent implements OnInit {
  groups!: Array<Group>;

  constructor(private adminService: AdminService, private userService: UserService) { }

  ngOnInit(): void {
    this.userService.getGroups().subscribe(res => {
      if(typeof res === 'string')
        return
      this.groups = res;
    }, err => {alert(err)})
  }

}
