import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ObjectId } from 'mongodb';
import { catchError, map, Observable, throwError } from 'rxjs';
import { environment } from 'src/environments/environment.prod';
import { Group } from 'src/models/group.model';
import { RegisterDto } from 'src/models/register-dto.model';
import { RegisterModel } from 'src/models/register.model';
import {Student} from 'src/models/Student';
import { UserModel } from 'src/models/user.model';
import { Pipe } from './pipe.model';

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  constructor(private client: HttpClient) {

   }

   public getUsers(type: string){
    return this.sendGetRequest("/users/type", [type])
   }

   public getStudents() : Promise<Array<UserModel>>{
    let headers = new HttpHeaders().set(
      "token", sessionStorage.getItem("token")!
    )

    var promise = new Promise<Array<UserModel>>(resolve => {
      this.client.get<Array<UserModel>>(`${environment.url}/users/type/STUDENT`, {headers: headers}).subscribe(res => {
        resolve(res);
      })
    })

    return promise;
   }

   public deleteUser(username: String){
     return this.client.delete(`${environment.url}/users/remove/${username}`).pipe(
      map( res => {return res;} ), 
      catchError(err => {return throwError(() => new Error(err.error))})
    )
   }

   public getSemesterStudentsPerformance(semester: number){
      return Pipe.makePipe(this.client.get(`${environment.url}/users/studentsbysemester/performancedesc/${semester}`))
   }

   public getStudentsByGroup(group_id: ObjectId){
     return this.sendGetRequest("/users/group", [group_id])
   }

   public generateUsers(type: string, users: RegisterDto[], group_id: ObjectId){
     if(type != 'STUDENT' && type != 'TEACHER' && type != 'ADMIN')
     {
      console.error("AdminService: invalid type for user type generation!")
      return;
     }
    return this.sendPostRequest("/users/generate/", [type, group_id], {userdtos: users})
   }

   public addGroup(group: Group){
     return this.sendPostRequest("/groups/add", [], group);
   }

   public addStudentToGroup(group_number: number, user_name: string){
     return this.sendPostRequest("/groups/add_student", [group_number, user_name], {})
   }

   public getYearPerformance(year: number){
     return this.sendGetRequest("/users/studentsbyyear/performancedesc", [year])
   }

   public geAllGroupsPerformance(){
     return this.sendGetRequest("/users/allgroups/allstudents/performancedesc", [])
   }

   public makeChief(user_name: string){
     return this.sendPostRequest("/makechief", [user_name], {})
   }

   public getGroupsPerformane(){
     return this.sendGetRequest("/users/allgroups/allstudents/performancedesc", []);
   }

   public sendGetRequest(route: string, params: any[]){
     let paramString: string = "";
     params.forEach(param => paramString += `/${param}`)
     return Pipe.makePipe(this.client.get(`${environment.url}${route}${paramString}`))
   }

   public sendPostRequest(route: string, params: any[], body: any){
    let paramString: string = "";
    params.forEach(param => paramString += `/${param}`)
    return Pipe.makePipe(this.client.post(`${environment.url}${route}${paramString}`, body))
  }

  public sendPutRequest(route: string, params: any[], body: any){
    let paramString: string = "";
    params.forEach(param => paramString += `/${param}`)
    return Pipe.makePipe(this.client.put(`${environment.url}${route}${paramString}`, body))
  }

  public sendDeleteRequest(route: string, params: any[], body: any){
    let paramString: string = "";
    params.forEach(param => paramString += `/${param}`)
    return Pipe.makePipe(this.client.delete(`${environment.url}${route}${paramString}`, body))
  }
}
