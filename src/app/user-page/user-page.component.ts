import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AcademicYear } from 'src/models/academic-year.model';
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
  public academicYear!: AcademicYear;
  
  constructor(public userService: UserService, private route: ActivatedRoute, private authService: AuthService, private router: Router) { 

  }

  ngOnInit(): void {
    this.getAcademicYear();
    let username = sessionStorage.getItem("username")
    
    this.username = username!;

    this.authService.getUser(this.username!).subscribe((res) => {
      this.user = res;
    });
  }

  public logOut(){
    this.authService.endSession();
    this.router.navigate([""]);
  }

  public getAcademicYear(){
      this.userService.getAcademicYear().subscribe(res => {
        sessionStorage.setItem("academic_year_id", res['ID'])

        let year = res as AcademicYear;
        this.academicYear = year;
        
        let currentDate = Date.now();
        let startDate = Date.parse(year.startdate);
        let enddate = Date.parse(year.enddate);

        let diff = Math.floor((currentDate - startDate) / (1000 * 3600 * 24))

        let can_sign = diff <= 14 && diff >= 0;
        let in_year = startDate <= currentDate && currentDate <= enddate;

        sessionStorage.setItem("can_sign", can_sign ? "true": "false");
        sessionStorage.setItem("in_year", in_year ? "true": "false");
      }, _ => {
        sessionStorage.setItem("can_sign", "false");
        sessionStorage.setItem("in_year", "false");
      })
  }

  public getCurrentYear(){
    let format: string = "";
    format += this.academicYear.startdate.slice(0, 10) + " to ";
    format += this.academicYear.enddate.slice(0, 10);

    return format;

  }
}
