import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ObjectId } from 'mongodb';
import { environment } from 'src/environments/environment.prod';
import { Amount } from 'src/models/amount.model';
import { Course } from 'src/models/course.model';
import { Pipe } from './pipe.model';

@Injectable({
  providedIn: 'root'
})
export class ChiefService {


  constructor(private client: HttpClient) { }

  public approveCourse(courseId: ObjectId){
    return Pipe.makePipe(this.client.post(`${environment.url}/courses/approve/${courseId}`, {max: 150}))
  }

  public addMandatoryCourse(userName: string, course: Course){
    return Pipe.makePipe(this.client.post(`${environment.url}/courses/addmandatory/${userName}`, course))
  }

  public getBestResults(){
    return Pipe.makePipe(this.client.get(`${environment.url}/enrollments/bestresults`))
  }

  public getWorstResults(){
    return Pipe.makePipe(this.client.get(`${environment.url}/enrollments/worstresults`))
  }
}

