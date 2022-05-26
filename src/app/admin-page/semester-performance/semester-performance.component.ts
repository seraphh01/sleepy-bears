import { Component, OnInit } from '@angular/core';
import { AdminService } from 'src/app/Services/admin.service';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-semester-performance',
  templateUrl: './semester-performance.component.html',
  styleUrls: ['./semester-performance.component.css']
})
export class SemesterPerformanceComponent implements OnInit {
  performance!: Map<UserModel, number>;;
  semester!: number;
  constructor(private service: AdminService) { }

  ngOnInit(): void {
  }

  public showPerformance(){
    this.service.getSemesterStudentsPerformance(this.semester!).subscribe(res => {
      let grades: number[] = res['averageGrade'];
      let students: UserModel[] = res['students'];

      this.performance = new Map<UserModel, number>();

      if(grades == null || students == null){
        alert("No grades were added for this semester to any student!");
        return;
      }
      
      var i: number;
      for(i = 0; i< grades.length; i++){
        this.performance.set(students[i], grades[i]);
      }
    })
  }
}
