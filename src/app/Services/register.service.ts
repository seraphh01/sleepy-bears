import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RegisterDto } from '../models/register-dto.model';
import { environment } from 'src/environments/environment.prod';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  constructor(private client: HttpClient) { }

  public RegisterStudent(data: RegisterDto[]){
    let headers = new HttpHeaders().set(
      "token", sessionStorage.getItem("token")!
    )
    this.client.post(`${environment.url}/users/generate`, {userdtos: data}, {headers: headers}).subscribe(res => {
        console.log(res);
    }, error => {
        console.error(error);
    });
  }
}
