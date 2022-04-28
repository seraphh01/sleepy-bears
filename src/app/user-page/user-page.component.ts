import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { UserModel } from '../models/user.model';
import { AuthService } from '../Services/auth.service';

@Component({
  selector: 'app-user-page',
  templateUrl: './user-page.component.html',
  styleUrls: ['./user-page.component.css']
})
export class UserPageComponent implements OnInit {
  public user!: UserModel;
  public username?: string;

  constructor(private route: ActivatedRoute, private authService: AuthService, private router: Router) { 

  }

  ngOnInit(): void {
    let username = this.route.snapshot.paramMap.get("username");
    
    this.username = username!;
    this.authService.getUser(this.username!).subscribe((res) => {
      this.user = res;
      console.log(res);
    });
  }

  public logOut(){
    this.authService.clearToken();
    this.router.navigate([""]);
  }
}
