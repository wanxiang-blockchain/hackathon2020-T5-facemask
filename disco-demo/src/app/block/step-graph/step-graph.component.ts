import { Component, OnInit, Output,Inject,EventEmitter} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';


@Component({
  selector: 'step-graph',
  templateUrl: './step-graph.component.html',
  styleUrls: ['./step-graph.component.scss']
})

export class GraphStepComponent implements OnInit {
  constructor() {
  }
  ngOnInit() {
  }

  openGraph(){
    window.open("www.baidu.com","_blank")
  }
}

