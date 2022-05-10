import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
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

   public getStudents() : Promise<Array<Student>>{
    let headers = new HttpHeaders().set(
      "token", sessionStorage.getItem("token")!
    )

    var promise = new Promise<Array<Student>>(resolve => {
      this.client.get<Array<Student>>(`${environment.url}/users/type/STUDENT`, {headers: headers}).subscribe(res => {
        resolve(res);
      })
    })

    return promise;
   }

}
