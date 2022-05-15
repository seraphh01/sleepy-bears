import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, map, throwError } from 'rxjs';
import { environment } from 'src/environments/environment.prod';
import { Course } from 'src/models/course.model';

@Injectable({
  providedIn: 'root'
})
export class TeacherService {

  constructor(private client: HttpClient) { }

  public proposeCourse(course: Course){
    return this.client.post(`${environment.url}/proposed_courses/add`, course)
    .pipe(
      map(_ => {return `Course ${course.name} was proposed successfully!`}),
      catchError(err => {
        return throwError(() => new Error(err.error.error));
      })
    )
  }

  public getProposedCourses(): Promise<Course[]>{
    let promise = new Promise<Course[]>(resolve => {
      this.client.get<Course[]>(`${environment.url}/proposed_courses`).subscribe((res: Course[]) => {
        resolve(res);
      }, (err) => {alert(err)});
    });

    return promise;
  }

  public getProposedCoursesByTeacher(username: String): Promise<Course[]>{
    let promise = new Promise<Course[]>(resolve => {
      this.client.get<Course[]>(`${environment.url}/proposed_courses/getby/${username}`).subscribe((res: Course[]) => {
        resolve(res);
      })
    });

    return promise;
  }
}
