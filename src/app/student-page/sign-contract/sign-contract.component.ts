import { Component, Input, OnInit } from '@angular/core';
import { StudentService } from 'src/app/Services/student.service';
import { UserModel } from 'src/models/user.model';

@Component({
  selector: 'app-sign-contract',
  templateUrl: './sign-contract.component.html',
  styleUrls: ['./sign-contract.component.css']
})
export class SignContractComponent implements OnInit {
  @Input() user!: UserModel;
  year: number = 1;
  constructor(private service: StudentService) { }

  ngOnInit(): void {
  }

  public signContract(){
    this.service.signContract(this.year).subscribe(res => {console.log(res)}, err => {alert(err)})
  }
}
