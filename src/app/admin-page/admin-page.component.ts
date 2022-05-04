import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, AbstractControl, ValidatorFn, ValidationErrors } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { Validators } from '@angular/forms';
import { AuthService } from '../Services/auth.service';
import { FileService } from '../Services/file-service.service';
import { RegisterService } from '../Services/register.service';
import { RegisterDto } from '../models/register-dto.model';
import { RegisterModel } from '../models/register.model';
import {UserModel} from '../models/user.model'
import { AdminService } from '../Services/admin.service';
import { Student } from 'src/models/Student';

@Component({
  selector: 'app-admin-page',
  templateUrl: './admin-page.component.html',
  styleUrls: ['./admin-page.component.css']
})
export class AdminPageComponent implements OnInit {
  students!: Student[];
  constructor(private service: AdminService){}

  async ngOnInit(): Promise<void> {
   
    this.students = await this.service.getStudents();
  }
}
