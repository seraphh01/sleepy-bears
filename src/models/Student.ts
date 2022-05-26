import { ObjectId } from "mongodb";
import { Group } from "./group.model";

export interface Student {
    _id: ObjectId;
    username: string;
    email: string;
    group: Group;
}
