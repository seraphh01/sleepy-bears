import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ObjectId } from 'mongodb';
import { catchError, map, throwError } from 'rxjs';
import { environment } from 'src/environments/environment.prod';
import { Course } from 'src/models/course.model';
import { Pipe } from './pipe.model';

@Injectable({
  providedIn: 'root'
})
export class TeacherService {

  constructor(private client: HttpClient) { }

  public proposeCourse(course: Course, username: string){
    return this.client.post(`${environment.url}/proposed_courses/add/${username}`, course)
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
  
  public getProposedCoursesByAcademicYear(year_of_study: number){
    let academic_year = sessionStorage.getItem("academic_year_id");
    return Pipe.makePipe(this.client.get(`${environment.url}/proposed_courses/academic_year/${academic_year}/${year_of_study}`))
  }

  public getProposedCoursesByTeacher(username: String){
    return Pipe.makePipe(this.client.get<Course[]>(`${environment.url}/proposed_courses/getby/${username}`));

  }

  public getCoursesByTeacher(username: String) {
    return Pipe.makePipe( this.client.get<Course[]>(`${environment.url}/mandatorycourses/getby/${username}`));
  }

  public getStudentsByCourse(courseID: ObjectId){
    return Pipe.makePipe(this.client.get(`${environment.url}/students/bycourse/${courseID}`));
  }

  public gradeStudent(courseId: ObjectId, studentUserName: String, grade: number){
    console.log(grade)
    return Pipe.makePipe(this.client.post(`${environment.url}/grades/add/${studentUserName}/${courseId}`, {grade: grade}))
  }

  public getGradesAtCourse(courseId: ObjectId){
    return Pipe.makePipe(this.client.get(`${environment.url}/students/average/${courseId}`))
  }

}
