import { Component, OnInit } from '@angular/core';
import { DataService } from './data.service';
import { TimestampData } from './data.model';

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"],
})

export class AppComponent implements OnInit {
  timestampData?: TimestampData;

  constructor(private dataService: DataService) { }

  ngOnInit() {
    this.dataService.getUpdatedTimestampData().subscribe(
      (response) => {
        this.timestampData = response;
        console.log(this.timestampData);
      },
      (error) => {
        console.error('Error fetching data:', error);
      }
    );
  }
}
