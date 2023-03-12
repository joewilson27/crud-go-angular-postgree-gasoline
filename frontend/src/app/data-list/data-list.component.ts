import { Observable } from "rxjs";
import { DataService } from "../data.service";
import { Data } from "../data";
import { Component, OnInit } from "@angular/core";
import { Router } from '@angular/router';

@Component({
  selector: 'app-data-list',
  templateUrl: './data-list.component.html',
  styleUrls: ['./data-list.component.css']
})
export class DataListComponent implements OnInit {
  datas: Observable<Data[]>;

  constructor(private dataService: DataService,
              private router     : Router) {
  }

  ngOnInit(): void {
    this.reloadData();
  }

  // TODO: Show datas list
  reloadData() {
    this.datas = this.dataService.getDatasList();
  }

  // deleteUser
  deleteData(id: string) {
  
    this.dataService.deleteData(id)
      .subscribe(
        data => {
          console.log(data);
          this.reloadData();
        },
        error => console.log(error));
  }

  // updateUser
  updateData(id: string){
    this.router.navigate(['update', id]);
  }

  // userDetails
  dataDetails(id: string){
    this.router.navigate(['details', id]);
  }

  numberWithDot(x) {
    return x.toString().replace(/\B(?<!\.\d*)(?=(\d{3})+(?!\d))/g, ".");
  }

}
