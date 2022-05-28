import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, AbstractControl, ValidatorFn, ValidationErrors } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { Validators } from '@angular/forms';
import { AuthService } from '../Services/auth.service';
import { FileService } from '../Services/file-service.service';
import { RegisterService } from '../Services/register.service';
import { RegisterDto } from '../../models/register-dto.model';
import { RegisterModel } from '../../models/register.model';
import {UserModel} from '../../models/user.model'
import { AdminService } from '../Services/admin.service';
import { Student } from 'src/models/Student';
import { threadId } from 'worker_threads';
import { Group } from 'src/models/group.model';
import { UserService } from '../Services/user.service';

@Component({
  selector: 'app-admin-page',
  templateUrl: './admin-page.component.html',
  styleUrls: ['./admin-page.component.css']
})
export class AdminPageComponent implements OnInit {
  public groups!: Group[];
  RegisterPage: string = "register";
  StudentPage: string = "students";

  activePage = "none";
  register = false;
  students!: UserModel[];
  constructor(private service: AdminService, private userService: UserService){}

  async ngOnInit(): Promise<void> {
    this.userService.getGroups().subscribe(res => {
      if(typeof res === 'string')
        return
      this.groups = res;
    },err => console.error(err));
  }

  public async getStudents(){
    this.students =  await this.service.getStudents();
    this.switchActivePage(this.StudentPage);
  }

  public switchActivePage(newPage: string){
    if(this.activePage == newPage){
      this.activePage = "none";
      return;
    }

    this.activePage = newPage;
  }

  public isPageActive(pageName: string): boolean{
    return pageName == this.activePage;
  }
}
