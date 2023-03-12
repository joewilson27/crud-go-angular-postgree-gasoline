import { DataService } from '../data.service';
import { Data } from '../data';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-data',
  templateUrl: './create-data.component.html',
  styleUrls: ['./create-data.component.css']
})
export class CreateDataComponent implements OnInit {

  data: Data = new Data();
  submitted = false;

  constructor(private dataService: DataService,
              private router: Router) { }

  ngOnInit(): void {
  }

  // newUser
  newData(): void {
    this.submitted = false;
    this.data = new Data();
  }

  save() {
    this.dataService.createData(this.data).subscribe(data => {
      console.log(data)
      this.data = new Data();
      this.gotoList();
    }, 
    error => console.log(error));
  }

  onSubmit() {
    this.save();  
    this.submitted = true;
  }

  gotoList() {
    this.router.navigate(['/datas']);
  }

  resetForm() {
    this.submitted = false;
    this.data = new Data();
  }

}
