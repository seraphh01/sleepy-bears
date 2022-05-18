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

@Component({
  selector: 'app-admin-page',
  templateUrl: './admin-page.component.html',
  styleUrls: ['./admin-page.component.css']
})
export class AdminPageComponent implements OnInit {
  RegisterPage: string = "register";
  StudentPage: string = "students";

  activePage = "none";
  register = false;
  students!: Student[];
  constructor(private service: AdminService){}

  async ngOnInit(): Promise<void> {
   
  }

  public async getStudents(){
    this.students =  await this.service.getStudents();
    console.log(this.students);
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
