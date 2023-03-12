import { Data } from '../data';
import { Component, OnInit, Input } from '@angular/core';
import { DataService } from '../data.service';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-data-details',
  templateUrl: './data-details.component.html',
  styleUrls: ['./data-details.component.css']
})
export class DataDetailsComponent implements OnInit {

  id: string;
  data: Data;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private dataService: DataService) { }

  ngOnInit() {
    this.data = new Data();

    this.id = this.route.snapshot.params['id'];
    console.log("gh"+this.id);
    this.dataService.getData(this.id)
      .subscribe(data => {
        console.log(data)
        this.data = data;
      }, error => console.log(error));
  }

  list(){
    this.router.navigate(['datas']);
  }

}
