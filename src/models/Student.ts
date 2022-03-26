import { ObjectId } from "mongodb";

export default class Student {
    constructor(public name: string, public id?: ObjectId) {
        
    }
}
