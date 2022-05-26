import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment.prod';
import { AcademicYear } from 'src/models/academic-year.model';
import { Course } from 'src/models/course.model';
import { Pipe } from './pipe.model';

@Injectable({
  providedIn: 'root'
})
export class StudentService {

  constructor(private client: HttpClient) { }

  public getMandatoryCourses(year: number): Promise<Course[]>{
      let promise = new Promise<Course[]>(resolve => {
        this.client.get<Course[]>(`${environment.url}/courses/year_of_study/${year}`).subscribe((res: Course[]) => {
          resolve(res);
        });
      });

      return promise;
  }

  public getOptionalCourses(user_name: String): Promise<Course[]>{
    let promise = new Promise<Course[]>(resolve => {
      this.client.get<Course[]>(`${environment.url}/enrollments/getbyusername/${user_name}`).subscribe((res: Course[]) => {
        resolve(res);
      });
    });

    return promise;
  }

  public enroll(courseId:string){
    return Pipe.makePipe(this.client.post(`${environment.url}/enrollments/add/${courseId}`, {}))
  }

  public getGrades(){
    return Pipe.makePipe(this.client.get(`${environment.url}/enrollments/grades`))
  }

  public signContract(){
    return Pipe.makePipe(this.client.post(`${environment.url}/enrollments/signcontract/628f570503edd9821d795c29`, {}))
  }
}
