import { UserModel } from "./user.model";

export interface StudentGrade{
    student: UserModel;
    grades: Array<number>;
}