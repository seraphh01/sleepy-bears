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

  public getUsers(type: string){
    return Pipe.makePipe(this.client.get(`${environment.url}/users/type/${type}`))
  }

  public getAcademicYear(){
    return Pipe.makePipe(this.client.get(`${environment.url}/academic_year`))
  }

  public canSign() : boolean{
    let can_sign = sessionStorage.getItem("can_sign");
    if(can_sign == null)
      return true;
      
    return can_sign == "true";
  }

  public inAcademicYear(): boolean {
    let in_year = sessionStorage.getItem("in_year");
    if(in_year == null)
      return true;
      
    return in_year == "true";
  }
}
