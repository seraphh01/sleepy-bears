import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment.prod';
import { Course } from 'src/models/course.model';

@Injectable({
  providedIn: 'root'
})
export class TeacherService {

  constructor(private client: HttpClient) { }

  public proposeCourse(course: Course){
    this.client.post(`${environment.url}/proposed_courses/add`, course).subscribe(res => {
      console.log(res);
    });
  }

  public getProposedCourses(): Promise<Course[]>{
    let promise = new Promise<Course[]>(resolve => {
      this.client.get<Course[]>(`${environment.url}/proposed_courses`).subscribe((res: Course[]) => {
        resolve(res);
      })
    });

    return promise;
  }
}
