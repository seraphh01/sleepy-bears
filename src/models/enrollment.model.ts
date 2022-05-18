import { ObjectID } from "bson";
import { Course } from "./course.model";
import { Grade } from "./Grade.model";
import {UserModel} from "./user.model"

export interface Enrollment{
    _id: ObjectID;
    user: UserModel;
    course: Course;
    grades: Grade[];
}