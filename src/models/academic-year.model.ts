import { ObjectId } from "mongodb";

export interface AcademicYear {
    ID: ObjectId;
    startdate: string;
    enddate: string;
}