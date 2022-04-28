import { Injectable } from '@angular/core';
import { Observable, shareReplay } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { RegisterModel } from '../models/register.model';
import { UserModel } from '../models/user.model';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private port: String= "5000";
  private url: String = `http://localhost:${this.port}`;


  constructor(private http: HttpClient) { }


  public login(username:String, password: String){
    var data = {
      username: username,
      password: password
    };
    return this.http.post<any>(this.url + "/users/login", data);
  }

  public register(data: RegisterModel){

    // data = {
    //   email: "sserafim.socaciu@gmail.com",
    //   username:"seraphh01",
    //   usertype:"ADMIN",
    //   password:"hellohello",
    //   name:"serafim",
    //   profileDescription:"hello"
    // } as RegisterModel;

    return this.http.post<any>(`${this.url}/users/signup`, data);
  }

  public getUser(username: string, token: string): Observable<UserModel>{
    let headers = new HttpHeaders().set(
      "token", token
    );
    return this.http.get<UserModel>(`${this.url}/users/${username}`, {headers: headers});
  }
}
