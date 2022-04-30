import { Injectable } from '@angular/core';
import { Observable, shareReplay } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { RegisterModel } from '../models/register.model';
import { UserModel } from '../models/user.model';
import { environment } from 'src/environments/environment.prod';

@Injectable({
  providedIn: 'root'
})
export class AuthService {


  constructor(private http: HttpClient) { }


  public login(username:String, password: String){
    var data = {
      username: username,
      password: password
    };
    return this.http.post<any>(environment.url + "/users/login", data);
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

    return this.http.post<any>(`${environment.url}/users/signup`, data);
  }

  public getUser(username: string): Observable<UserModel>{
    let headers = new HttpHeaders().set(
      "token", sessionStorage.getItem("token")!
    );
    return this.http.get<UserModel>(`${environment.url}/users/${username}`, {headers: headers});
  }

  public clearToken(){
    sessionStorage.removeItem("token");
  }

  public deleteUser(username: string){
    let headers = new HttpHeaders().set(
      "token", sessionStorage.getItem("token")!
    );
    return this.http.delete(`${environment.url}/users/remove/${username}`, {headers: headers});
  }
}
