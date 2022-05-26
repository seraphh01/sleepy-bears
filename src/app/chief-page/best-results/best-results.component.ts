import { Component, OnInit } from '@angular/core';
import { ChiefService } from 'src/app/Services/chief.service';
import { UserService } from 'src/app/Services/user.service';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-best-results',
  templateUrl: './best-results.component.html',
  styleUrls: ['./best-results.component.css']
})
export class BestResultsComponent implements OnInit {
  bestTeacher!: UserModel;
  grade!: number;
  constructor(private service: ChiefService) { }

  ngOnInit(): void {
    this.getBestResults();
  }

  getBestResults(){
    this.service.getBestResults().subscribe((res: any) => {
      this.bestTeacher = res['bestTeacher'];
      this.grade = res['averageGrade'];
    })
  }

}
