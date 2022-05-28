import { ObjectId } from "mongodb";
import { AcademicYear } from "./academic-year.model";

export interface Group {
    ID: ObjectId;
    academicyear: AcademicYear;
    number: number;
    year: number;
}