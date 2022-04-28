import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
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
  private token!: string;

  constructor(private route: ActivatedRoute, private authService: AuthService) { 

  }

  ngOnInit(): void {
    let username = this.route.snapshot.paramMap.get("username");
    let token = this.route.snapshot.paramMap.get("token");
    
    this.username = username!;
    this.token = token!;
    this.authService.getUser(this.username!, this.token!).subscribe((res) => {
      this.user = res;
      console.log(res);
    });
  }

}
