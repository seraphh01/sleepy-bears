import { Component, OnInit } from '@angular/core';
import { ChiefService } from 'src/app/Services/chief.service';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-worst-results',
  templateUrl: './worst-results.component.html',
  styleUrls: ['./worst-results.component.css']
})
export class WorstResultsComponent implements OnInit {
  worstTeacher!: UserModel;
  grade!: number;
  constructor(private service: ChiefService) { }

  ngOnInit(): void {
    this.getBestResults();
  }

  getBestResults(){
    this.service.getWorstResults().subscribe((res: any) => {
      this.worstTeacher = res['bestTeacher'];
      this.grade = res['averageGrade'];
    })
  }
}
