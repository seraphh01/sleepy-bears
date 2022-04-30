import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RegisterDto } from '../models/register-dto.model';
import { environment } from 'src/environments/environment.prod';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  constructor(private client: HttpClient) { }

  public RegisterStudents(data: RegisterDto[], type: string){
    let headers = new HttpHeaders().set(
      "token", sessionStorage.getItem("token")!
    )
    return this.client.post(`${environment.url}/users/generate/${type}`, {userdtos: data}, {headers: headers});
  }
}
