import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ObjectId } from 'mongodb';
import { environment } from 'src/environments/environment.prod';
import { Pipe } from './pipe.model';

@Injectable({
  providedIn: 'root'
})
export class ChiefService {


  constructor(private client: HttpClient) { }

  public approveCourse(courseId: ObjectId){
    return Pipe.makePipe(this.client.post(`${environment.url}/courses/approve/${courseId}`, {}))
  }
}

