import { Component, Input, OnInit } from '@angular/core';
import { FormGroup, FormControl, AbstractControl, ValidatorFn, ValidationErrors } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { Validators } from '@angular/forms';
import { AuthService } from '../../Services/auth.service';
import { FileService } from '../../Services/file-service.service';
import { RegisterService } from '../../Services/register.service';
import { RegisterDto } from '../../../models/register-dto.model';
import { RegisterModel } from '../../../models/register.model';
import {UserModel} from '../../../models/user.model'
import { Group } from 'src/models/group.model';

@Component({
  selector: 'app-register-users',
  templateUrl: './register-users.component.html',
  styleUrls: ['./register-users.component.css']
})
export class RegisterUsersComponent implements OnInit {
  @Input() public groups!: Group[];
  public group!: Group;
  public register?: RegisterDto;
  public usertype: string = 'STUDENT';
  public form: FormGroup;
  public file!:File;
  public fileContent!:string;
  public registeredStudents!: RegisterModel[];

  constructor(private route: ActivatedRoute, private authService: AuthService,
    private registerService: RegisterService, private fileService: FileService) { 
    this.form = new FormGroup (
      {
        name: new FormControl('', Validators.required),
        cnp: new FormControl('', [Validators.required, Validators.minLength(13), Validators.maxLength(13)] ),
        usertype: new FormControl('STUDENT', [Validators.required]),
        group: new FormControl({} as Group, [Validators.required])
      }
    );

    this.form.valueChanges.subscribe((val) => {
      this.register = val;
    });
  }

  ngOnInit(): void {
    this.group = this.groups[0];
    this.form.patchValue({group: this.group})
  }

  public addUser(){
    if(!this.form.valid){
      alert("Invalid values in the form");
      return;
    }

    this.registerService.RegisterStudents([this.register!], this.form.get("usertype")?.value, this.form.get("group")?.value).subscribe((res: any) => {
      
      this.registeredStudents = res;
      this.downloadStudentsList();
    });
  }

  async fileChanged(e: any) {
    this.file = e.target.files[0];
    this.fileContent = await this.fileService.readFileContent(this.file)!;
  }

  usertypeChanged(e: any){
    this.usertype = e.target.value;
  }

  registerUsers(){
    if(!this.fileContent){
      return;
    }

    let studentList = new Array<RegisterDto>();

    for (const line of this.fileContent.split(/[\r\n]+/)){
      if(line.length == 0)
        continue; 
      let tokens = line.split(/[,]+/);
      let name:string = tokens[0];
      let cnp:string = tokens[1].replace(' ', '');

      if(cnp.length != 13 || name.length == 0)
        continue;

      studentList.push({name: name, CNP: cnp} as RegisterDto);
      
    }

    this.registerService.RegisterStudents(studentList, this.usertype, this.group).subscribe((res: any) => {
      this.registeredStudents = res;

      // for(let user of this.registeredStudents){
      //   this.authService.deleteUser(user.username).subscribe((res) => {
      //     console.log("deleted: " + user.username);
      //   });
      // }

    }, error => {console.log(error)});
  }

  downloadStudentsList(){
    if(!this.registeredStudents || this.registeredStudents.length == 0){
      alert("There are no new registered students");
      return;
    }

    let studentsList: string = "username, password, email, name, usertype\n";
    
    for(let student of this.registeredStudents){
      studentsList += `${student.username}, ${student.password}, ${student.email}, ${student.name}, ${student.usertype}\n`
    }

    this.fileService.downloadTextFile(studentsList, '.csv');
  }

  public groupChanged(e: any){
    this.group = e.target.value;
  }
}
