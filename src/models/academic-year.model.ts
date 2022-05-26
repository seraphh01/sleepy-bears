import { ObjectId } from "mongodb";

export interface AcademicYear {
    ID: ObjectId;
    start_date: Date;
    end_date: Date;
}