import { Component, OnInit } from '@angular/core';
import { StudentService } from 'src/app/Services/student.service';

@Component({
  selector: 'app-sign-contract',
  templateUrl: './sign-contract.component.html',
  styleUrls: ['./sign-contract.component.css']
})
export class SignContractComponent implements OnInit {

  constructor(private service: StudentService) { }

  ngOnInit(): void {
  }

  public signContract(){
    this.service.signContract().subscribe(res => {console.log(res)}, err => {alert(err)})
  }
}
