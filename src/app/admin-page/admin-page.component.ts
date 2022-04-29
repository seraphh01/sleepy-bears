import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, AbstractControl, ValidatorFn, ValidationErrors } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { RegisterModel } from '../models/register.model';
import { Validators } from '@angular/forms';
import { AuthService } from '../Services/auth.service';
import { RegisterService } from '../Services/register.service';
import { RegisterDto } from '../models/register-dto.model';

@Component({
  selector: 'app-admin-page',
  templateUrl: './admin-page.component.html',
  styleUrls: ['./admin-page.component.css']
})
export class AdminPageComponent implements OnInit {
  public register?: RegisterDto;
  public username?: String;
  public form: FormGroup;

  constructor(private route: ActivatedRoute, private authService: AuthService, private registerService: RegisterService) { 
    this.form = new FormGroup (
      {
        name: new FormControl('', Validators.required),
        cnp: new FormControl('', [Validators.required, Validators.minLength(13), Validators.maxLength(13)] )
      }
    );

    this.form.valueChanges.subscribe((val) => {
      this.register = val;
    });
  }

  ngOnInit(): void {
    let username = this.route.snapshot.paramMap.get("username");
    let token = this.route.snapshot.paramMap.get("token");
    
    this.username = username ? username : "undefined";
  }

  public addUser(){
    console.log(this.register);
    if(!this.form.valid){
      alert("Invalid values in the form");
      return;
    }

    this.registerService.RegisterStudent([this.register!]);
  }
}
