import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AdminService } from 'src/app/Services/admin.service';
import { Group } from 'src/models/group.model';

@Component({
  selector: 'app-add-group',
  templateUrl: './add-group.component.html',
  styleUrls: ['./add-group.component.css']
})
export class AddGroupComponent implements OnInit {
  public formGroup!: FormGroup;

  constructor(private adminService: AdminService) {
    this.formGroup = new FormGroup ({
      number: new FormControl(911, [Validators.required, Validators.min(100), Validators.max(1000)]),
      year: new FormControl(1, [Validators.required, Validators.min(1)])
    })

    this.formGroup.get('number')?.valueChanges.subscribe(value => {
      this.formGroup.patchValue({year: Math.floor(value/10) % 10})
    })
  }

  ngOnInit(): void {
  }

  public addGroup(){
    let group = this.formGroup.getRawValue();
    group.academicyear = {};
    
    if(Math.floor(group.number / 10) % 10 != group.year){
      alert("Second digit of group number must be equal to year!");
      return;
    }

    this.adminService.addGroup(group).subscribe(res => {
      confirm(`Group with number ${group.number}, year ${group.year} was added successfully!`)
    }, err => alert(err));
  }
}
