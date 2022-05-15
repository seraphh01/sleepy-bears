import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment.prod';
import { Course } from 'src/models/course.model';

@Injectable({
  providedIn: 'root'
})
export class StudentService {

  constructor(private client: HttpClient) { }

  public getMandatoryCourses(year: number): Promise<Course[]>{
      let promise = new Promise<Course[]>(resolve => {
        this.client.get<Course[]>(`${environment.url}/courses/${year}`).subscribe((res: Course[]) => {
          console.log(res);
          resolve(res);
        });
      });

      return promise;
  }

  public getOptionalCourses(user_name: String): Promise<Course[]>{
    let promise = new Promise<Course[]>(resolve => {
      this.client.get<Course[]>(`${environment.url}/enrollments/getbyusername/${user_name}`).subscribe((res: Course[]) => {
        console.log(res);
        resolve(res);
      });
    });

    return promise;
  }

  public enroll(courseId:string){
    this.client.post(`${environment.url}/enrollments/add/${courseId}`, {}).subscribe((res:any) => {
        console.log(res);
    })
  }
}
