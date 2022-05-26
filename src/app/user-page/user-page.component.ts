import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { UserModel } from 'src/models/user.model';
import { AuthService } from '../Services/auth.service';
import { UserService } from '../Services/user.service';

@Component({
  selector: 'app-user-page',
  templateUrl: './user-page.component.html',
  styleUrls: ['./user-page.component.css']
})
export class UserPageComponent implements OnInit {
  public user!: UserModel;
  public username?: string;
  
  constructor(private userService: UserService, private route: ActivatedRoute, private authService: AuthService, private router: Router) { 

  }

  ngOnInit(): void {
    this.getAcademicYear();
    let username = sessionStorage.getItem("username")
    
    this.username = username!;

    this.authService.getUser(this.username!).subscribe((res) => {
      this.user = res;
      console.log(res);
    });
  }

  public logOut(){
    this.authService.endSession();
    this.router.navigate([""]);
  }

  public getAcademicYear(){
      this.userService.getAcademicYear().subscribe(res => {
        sessionStorage.setItem("academic_year_id", res['ID'])
      })
  }
}
