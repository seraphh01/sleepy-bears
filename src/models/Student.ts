import { ObjectId } from "mongodb";
import { Group } from "./group.model";

export interface Student {
    username: string;
    email: string;
    group: Group;
}
