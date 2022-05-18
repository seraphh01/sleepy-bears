import { Group } from "./group.model";

export interface UserModel {
    username: String;
    email: String;
    usertype:String;
    password:String;
    name:String;
    profileDescription:String;
    token:string;
    group: Group;
    refreshToken: string;
}