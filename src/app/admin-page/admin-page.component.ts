import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, AbstractControl, ValidatorFn, ValidationErrors } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { RegisterModel } from '../models/register.model';
import { Validators } from '@angular/forms';
import { AuthService } from '../Services/auth.service';

@Component({
  selector: 'app-admin-page',
  templateUrl: './admin-page.component.html',
  styleUrls: ['./admin-page.component.css']
})
export class AdminPageComponent implements OnInit {
  public register?: RegisterModel;
  public username?: String;
  public token?:String;
  public form: FormGroup;

  constructor(private route: ActivatedRoute, private authService: AuthService) { 
    this.form = new FormGroup (
      {
        email : new FormControl('', [Validators.required, Validators.email]),
        username: new FormControl(
        '', [Validators.required, Validators.minLength(6), Validators.maxLength(32)]
        ),
        password: new FormControl(
          '', [Validators.required, Validators.minLength(8), Validators.maxLength(24)]
        ),
        usertype: new FormControl(
          'STUDENT', [Validators.required]
        ),
        name: new FormControl(''),
        profileDescription: new FormControl('A new user')
      }
    );

    this.form.valueChanges.subscribe((val) => {
      this.register = val;

    });
  }

  private userTypeValidator(): ValidatorFn {

    return (control: AbstractControl) : ValidationErrors | null => {
      const value = control.value;

      if(!value){
        return null;
      }
  
      if(value == "ADMIN" || value=="STUDENT" || value == "TEACHER" || value == "HEAD"){
        return {validType: true};
      }
  
      return null;
    }   
  }

  ngOnInit(): void {
    let username = this.route.snapshot.paramMap.get("username");
    let token = this.route.snapshot.paramMap.get("token");
    
    this.username = username ? username : "undefined";
    this.token = token ? token : "undefined";
  }

  public addUser(){
    console.log(this.register);
    if(!this.form.valid){
      alert("Invalid values in the form");
      return;
    }

    this.authService.register(this.register!).subscribe((res) =>{
      console.log(res);
    });
  }
}
