import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})
export class LoginFormComponent implements OnInit {
  public email?: String;
  public register: Boolean = false
  constructor() { }

  ngOnInit(): void {

  }

  public onLogin(){
    
  }

  public onRegister(){

  }
}
