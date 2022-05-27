import { AfterViewInit, COMPILER_OPTIONS, Component, OnInit } from '@angular/core';
import { StudentService } from 'src/app/Services/student.service';
import { Course } from 'src/models/course.model';
import { Grade } from 'src/models/Grade.model';
import { ChangeDetectorRef } from '@angular/core';
import { ArrayType } from '@angular/compiler';

@Component({
  selector: 'app-view-grades',
  templateUrl: './view-grades.component.html',
  styleUrls: ['./view-grades.component.css']
})
export class ViewGradesComponent implements OnInit, AfterViewInit {
  courses!: Course[];
  grades!: Grade[][];

  constructor(private service: StudentService, private changeDetector: ChangeDetectorRef) { }

  ngOnInit(): void {

  }

  ngAfterViewInit(){
    this.service.getGrades().subscribe((res: any) => {
      let courses = res['courses'];
      let coursesGrades = res['grades'];
      var i : number;

      this.courses = new Array<Course>();
      this.grades = new Array<Grade[]>();

      if(courses == null)
        return;

      if(courses.length == 0){
         return;
      }

      for(i =0 ;i< courses.length;i++){
        let course = courses[i];
        let grades = coursesGrades[i];

        if(grades == null || grades.length == 0){
          continue;
        }

        this.courses.push(course);
        this.grades.push(grades);
      }
    });
  }

  public average(grades: Grade[]): number{
    let total: number = 0;
    grades.forEach(grade => total += grade.grade);

    return +(total / grades.length ).toFixed(3);
  }

  public getGradesValues(gradesObj: Grade[]){
    let gradeValues: number[] = new Array<number>();
    gradesObj.forEach(grade => gradeValues.push(grade.grade));

    return gradeValues;
  }
}
