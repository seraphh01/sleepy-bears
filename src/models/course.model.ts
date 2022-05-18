import { Amount } from "./amount.model";
import { ObjectId } from "mongodb";

export interface Course {
    ID: ObjectId;
    name: string;
    coursetype: string;
    year: number;
    maxamount: Amount;
}