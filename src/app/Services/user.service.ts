import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment.prod';
import { Pipe } from './pipe.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private client: HttpClient) { }

  public getOptionalCourses(){
    return Pipe.makePipe(this.client.get(`${environment.url}/proposed_courses`))
  }
}
