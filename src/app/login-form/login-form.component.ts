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
        
        this.redirect(result);
      }, error => {
        confirm("Invalid username or password!");
      });
    }else{
      alert("Username or password is empty");
    }
  }

  public redirect(userData: any){
    
    this.router.navigate([`users/${this.username}/${userData.token}`]);
  }
}
