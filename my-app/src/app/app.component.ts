import { Component, OnInit } from '@angular/core';
import { DataService } from './data.service';
import { DataModel } from './data.model';

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"],
})

export class AppComponent implements OnInit {
  data?: DataModel;

  constructor(private dataService: DataService) { }

  ngOnInit() {
    this.dataService.getData().subscribe(
      (response) => {
        this.data = response;
        console.log('Data from server:', this.data);
      },
      (error) => {
        console.error('Error fetching data:', error);
      }
    );
  }
}
