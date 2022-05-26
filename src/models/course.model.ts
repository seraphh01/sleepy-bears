import { Amount } from "./amount.model";
import { ObjectId } from "mongodb";
import { AcademicYear } from "./academic-year.model";

export interface Course {
    ID: ObjectId;
    name: string;
    coursetype: string;
    year: number;
    maxamount: Amount;
    year_of_study: number;
    academic_year: AcademicYear;
    credits: number;
    semester: number;
}