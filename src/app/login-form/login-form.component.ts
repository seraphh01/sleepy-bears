import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, EmailValidator } from '@angular/forms';
import { AuthService } from '../Services/auth.service';
import { Router, ActivatedRoute, Route } from '@angular/router';
import { asLiteral } from '@angular/compiler/src/render3/view/util';
import { RegisterModel } from '../models/register.model';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})
export class LoginFormComponent implements OnInit {
  private username?: String;
  private password?: String;
  public form: FormGroup;

  constructor(public authService: AuthService, private router: Router, 
    private activatedRoute: ActivatedRoute) {
    this.form = new FormGroup({
      username : new FormControl(),
      password: new FormControl()
    });
   }

  ngOnInit(): void {

  }

  ngAfterViewInit(): void {
    this.form.get('username')?.valueChanges.subscribe(data => this.username = data);
    this.form.get('passwords')?.statusChanges.subscribe(data => this.password = data);
} 

  public onLogin(){
    this.username = this.form.get('username')?.value;
    this.password = this.form.get('password')?.value;

    if(this.username != undefined && this.password != undefined){
      this.authService.login(this.username, this.password).subscribe(result => {
        sessionStorage.setItem("token", result.token);
        sessionStorage.setItem("username", result.username);
        this.redirect();
      }, error => {
          let containerBody = <HTMLElement>document.querySelector('.box');
          containerBody.classList.add('grow');
          let containerText = <HTMLElement>document.getElementById('boxText');
          containerText.innerHTML = "Invalid Username or Password";
          containerText.classList.add('growText');
          let containerButton = <HTMLElement>document.getElementById('boxButton');
          containerButton.style.display = 'unset';
          containerButton.style.width = '40px';
          containerButton.style.height = '20px';
      });
    }else{
      alert("Username or password is empty");
    }
  }

  public redirect(){
    this.router.navigate([`users`]);
  }
}
