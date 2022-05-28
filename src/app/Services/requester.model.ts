import { HttpClient } from "@angular/common/http";
import { environment } from "src/environments/environment.prod";
import { Pipe } from "./pipe.model";

export class Requests {
   public static sendGetRequest(client: HttpClient, route: string, params: any[]){
    let paramString: string = "";
    params.forEach(param => paramString += `/${param}`)
    return Pipe.makePipe(client.get(`${environment.url}${route}${paramString}`))
  }

  public static sendPostRequest(client: HttpClient, route: string, params: any[], body: any){
   let paramString: string = "";
   params.forEach(param => paramString += `/${param.toString()}`)
   return Pipe.makePipe(client.post(`${environment.url}${route}${paramString}`, body))
 }

 public static sendPutRequest(client: HttpClient, route: string, params: any[], body: any){
   let paramString: string = "";
   params.forEach(param => paramString += `/${param.toString()}`)
   return Pipe.makePipe(client.put(`${environment.url}${route}${paramString}`, body))
 }

 public static  sendDeleteRequest(client: HttpClient, route: string, params: any[], body: any){
   let paramString: string = "";
   params.forEach(param => paramString += `/${param.toString()}`)
   return Pipe.makePipe(client.delete(`${environment.url}${route}${paramString}`, body))
 }
}