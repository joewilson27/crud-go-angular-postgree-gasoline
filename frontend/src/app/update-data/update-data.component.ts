import { Component, OnInit } from '@angular/core';
import { Data } from '../data';
import { ActivatedRoute, Router } from '@angular/router';
import { DataService } from '../data.service';

@Component({
  selector: 'app-update-data',
  templateUrl: './update-data.component.html',
  styleUrls: ['./update-data.component.css']
})
export class UpdateDataComponent implements OnInit {

  id: string;
  data: Data;
  updated = false;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private dataService: DataService) { }

  ngOnInit() {
    this.data = new Data();

    this.id = this.route.snapshot.params['id'];
    
    this.dataService.getData(this.id)
                    .subscribe(data => {
                      this.data = data;
                    }, error => console.log(error));
  }

  // updateUser
  updateData() {
    this.dataService.updateData(this.id, this.data)
      .subscribe(data => {
        console.log(data);
        this.data = new Data();
        this.gotoList();
      }, error => console.log(error));
  }

  onSubmit() {
    this.updateData(); 
    this.updated = true;   
  }

  gotoList() {
    this.router.navigate(['/datas']);
  }

}
