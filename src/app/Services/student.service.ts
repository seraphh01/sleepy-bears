import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ObjectId } from 'mongodb';
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
        this.client.get<Course[]>(`${environment.url}/courses`).subscribe((res: Course[]) => {
          resolve(res);
        });
      });

      return promise;
  }

  public getOptionalCourses(user_name: String){
    return Pipe.makePipe( this.client.get<Course[]>(`${environment.url}/enrollments/getbyusername/${user_name}`));
  }

  public enroll(courseId: ObjectId){
    return Pipe.makePipe(this.client.post(`${environment.url}/enrollments/add/${courseId}`, {}))
  }

  public unenroll(courseId:ObjectId){
    return Pipe.makePipe(this.client.delete(`${environment.url}/enrollments/remove/${courseId}`, {}))
  }

  public getGrades(){
    return Pipe.makePipe(this.client.get(`${environment.url}/enrollments/grades`))
  }

  public getGradesByYear(username: string){
    return Pipe.makePipe(this.client.get(`${environment.url}/enrollments/grades/by-year/${username}`))
  }

  public signContract(year: number){
    return Pipe.makePipe(this.client.post(`${environment.url}/enrollments/add/by-year/${year}`, {}))
  }
}
