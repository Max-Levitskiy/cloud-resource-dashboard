import { Component, OnInit } from '@angular/core';
import {ApiService} from '../../../services/api.service';

@Component({
  selector: 'app-control-buttons',
  templateUrl: './control-buttons.component.html',
  styleUrls: ['./control-buttons.component.scss']
})
export class ControlButtonsComponent implements OnInit {

  constructor(private api: ApiService) { }

  ngOnInit(): void {
  }

  runFullScan() {
    this.api.runFullScan();
  }
}
