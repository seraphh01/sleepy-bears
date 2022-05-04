import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment.prod';
import {Student} from 'src/models/Student';

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  constructor(private client: HttpClient) {

   }

   public getUsers(){
    let headers = new HttpHeaders().set(
      "token", sessionStorage.getItem("token")!
    )
    return this.client.get(`${environment.url}/users`, {headers: headers});
   }
   public getStudents(): Promise<Array<Student>>{
     let promise = new Promise<Array<Student>>(resolve => {
      this.getUsers().subscribe((users: any) => {
        let students: Array<Student> = new Array<Student>();
        users["user_items"].forEach((user: any) => {
          if(user["usertype"] == "STUDENT"){
            students.push(user as Student);
          }
        });
        resolve(students);
      });
     });
    
    return promise;
   }

}
