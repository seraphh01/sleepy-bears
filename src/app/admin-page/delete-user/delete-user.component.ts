import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AdminService } from 'src/app/Services/admin.service';

@Component({
  selector: 'app-delete-user',
  templateUrl: './delete-user.component.html',
  styleUrls: ['./delete-user.component.css']
})
export class DeleteUserComponent implements OnInit {

  public usernameGroup!: FormGroup;

  constructor(private adminService: AdminService) { 
    this.usernameGroup = new FormGroup({
      "username": new FormControl('', Validators.required)
    })
  }

  ngOnInit(): void {
  }

  public deleteUser(){
    this.adminService.deleteUser(this.usernameGroup.get("username")!.value).subscribe(res => {alert("Deleted")}, err => {alert("User not found")});
  }
}
