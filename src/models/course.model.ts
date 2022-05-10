import { Amount } from "./amount.model";
import { ObjectId } from "mongodb";

export interface Course {
    ID: ObjectId;
    name: string;
    courseType: string;
    year: number;
    maxAmount: Amount;
}