import { Amount } from "./amount.model";
import { ObjectId } from "mongodb";
import { AcademicYear } from "./academic-year.model";
import { UserModel } from "./user.model";

export interface Course {
    ID: ObjectId;
    name: string;
    coursetype: string;
    year: number;
    maxamount: Amount;
    yearofstudy: number;
    academic_year: AcademicYear;
    credits: number;
    semester: number;
    proposer: UserModel;
}