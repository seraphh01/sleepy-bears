import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-user-page',
  templateUrl: './user-page.component.html',
  styleUrls: ['./user-page.component.css']
})
export class UserPageComponent implements OnInit {
  public username?: String;

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    console.log("hello from userpage")
    let username = this.route.snapshot.paramMap.get("username");
    
    console.log(username);
    
    this.username = username ? username : "undefined";
  }

}
