import { Component, OnInit, Output,Inject,EventEmitter} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';


@Component({
  selector: 'step-analysis',
  templateUrl: './step-analysis.component.html',
  styleUrls: ['./step-analysis.component.scss']
})

export class AnalysisStepComponent implements OnInit {
  constructor() {
  }
  ngOnInit() {
  }

  openNeo4j(){
    window.open("www.baidu.com","_blank")
  }
}

